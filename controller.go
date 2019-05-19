// +build linux darwin !appengine

package tiny

import (
	"encoding/xml"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const SW_VERSION = "1.6.0"

type Flags struct {
	Disabled bool `json:"disabled,omitempty"`  // if true, the structure must not be considered anymore
	Hidden   bool `json:"hidden,omitempty"`    // if true, the structure may not be considered by UI
	Lost     bool `json:"lost,omitempty"`      // if true, communication with the controller is lost
	Realtime bool `json:"realtime,omitempty"`  // if true, the structure must be considered by RT engine
	TestMode bool `json:"test_mode,omitempty"` // if true, do not store results
}

// Outs structure for controllers.
type Outs struct {
	Out0 int `xml:"out0"` // onboard relay
	Out1 int `xml:"out1"` // extension relay 1
	Out2 int `xml:"out2"` // extension relay 2
	Out3 int `xml:"out3"` // extension relay 3
	Out4 int `xml:"out4"` // extension relay 4
	Out5 int `xml:"out5"` // extension relay 5
	Out6 int `xml:"out6"` // extension relay 6
}

// Invert every out status. Useful because v2 are inversed.
func (o *Outs) invertOuts() {
	o.Out0 ^= 1
	o.Out1 ^= 1
	o.Out2 ^= 1
	o.Out3 ^= 1
	o.Out4 ^= 1
	o.Out5 ^= 1
	o.Out6 ^= 1
}

type Controller struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ApiKey      string `json:"api_key"`
	Ip          string `json:"ip"`
	Port        int    `json:"port"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Hw          string `json:"hw"`
	Hw_Version  int    `json:"hw_version"`
	Lk2         `json:"-"`
	Lk3         `json:"-"`
	Pi3         `json:"-"`
	Outs        `json:"-"`
	Flags
	Sensors []Sensor `json:"sensors"`
}

func (c *Controller) getField(field string) reflect.Value {
	r := reflect.ValueOf(c)
	f := reflect.Indirect(r).FieldByName(field)
	return f
}

func (c *Controller) getTypeField(field string) string {
	f := c.getField(field)
	if f.IsValid() {
		return f.Type().Name()
	}
	return "invalid"
}

func (c *Controller) getIntField(field string) int64 {
	f := c.getField(field)
	if f.IsValid() {
		return f.Int()
	}
	return 0
}

func (c *Controller) getFloatField(field string) float64 {
	f := c.getField(field)
	if f.IsValid() {
		return f.Float()
	}
	return 0
}

func (c *Controller) getStringField(field string) string {
	f := c.getField(field)
	if f.IsValid() {
		return f.String()
	}
	return "invalid"
}

func (c *Controller) getBoolField(field string) bool {
	f := c.getField(field)
	if f.IsValid() {
		return f.Bool()
	}
	return false
}

// Build an sensors array with different types of the value.
func (c *Controller) buildSensors() {
	for i := range c.Sensors {
		s := &c.Sensors[i]
		s.Lost = c.Lost
		switch c.getTypeField(s.Name) {
		case "int":
			s.Value = c.getIntField(s.Name)
		case "float64":
			v := c.getFloatField(s.Name)
			v *= float64(s.SetMultiplier())
			v /= float64(s.SetDivisor())
			s.Value = v
		case "string":
			s.Value = c.getStringField(s.Name)
		case "bool":
			s.Value = c.getBoolField(s.Name)
		default:
			s.Disabled = true
		}
	}
}

// Make a request to the controller with specified method.
func (c *Controller) request(m, r string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Duration(time.Second),
	}
	req, err := http.NewRequest(m, "http://"+c.Ip+":"+strconv.Itoa(c.Port)+"/"+r, nil)
	req.SetBasicAuth(c.User, c.Password)
	res, err := client.Do(req)
	c.Lost = false
	if err != nil || res.StatusCode != 200 {
		c.Lost = true
	}
	return res, err
}

// Get send commands to the controller using GET method.
func (c *Controller) get(x string) error {
	res, err := c.request("GET", x)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

// Get an xml file from controller and decode response in the proper structure
// return Lost flag if the controller not responding.
func (c *Controller) getFile(x string) error {
	res, err := c.request("GET", x)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err = xml.NewDecoder(res.Body).Decode(&c); err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetName() string {
	return c.Lk2.GetName() + c.Lk3.GetName() + c.Pi3.GetName()
}

func (c *Controller) GetVersion() string {
	return c.Lk2.GetVersion() + c.Lk3.GetVersion() + c.Pi3.GetVersion()
}

func (c *Controller) GetSensors() {
	switch c.Hw {
	case "lk":
		switch c.Hw_Version {
		case 2:
			c.getFile("st0.xml")
			c.invertOuts()
		case 3:
			c.getFile("xml/ix.xml")
		}
	case "pi":
	}
	c.buildSensors()
}

func (c *Controller) GetStatus() {
	switch c.Hw {
	case "lk":
		switch c.Hw_Version {
		case 2:
			c.getFile("st2.xml")
		case 3:
			c.getFile("xml/stat.xml")
		}
	case "pi":
	}
}

func (c *Controller) GetConfig() {
	switch c.Hw {
	case "lk":
		switch c.Hw_Version {
		case 2:
			c.getFile("board.xml")
		case 3:
			c.getFile("xml/st.xml")
		}
	case "pi":
		c.getFile("")
	}
}

// Set physical out state for v2 and v3 controllers.
// o out out0-6
// s status standard is 0-off,1-on
func (c *Controller) SetOut(o string, s int) {
	switch c.Hw {
	case "lk":
		if c.Hw_Version == 2 {
			s ^= 1 // relay status is inverted in v2
		}
		c.get("outs.cgi?" + strings.ToLower(o) + "=" + strconv.Itoa(s))
	case "pi":
	}
}
