// +build linux darwin !appengine

package tiny

const (
	U_TEMP    = "°C"
	U_HUMID   = "%"
	U_VOLTAGE = "V"
	U_CURRENT = "A"
	U_ENERGY  = "kwh"
	U_WATER   = "m³"
	U_STATUS  = "sw"
	U_AIRQ    = "μg/m³"
	U_DUSK    = "V/lux"
	U_PRESS   = "hPa"
)

type Coordinates struct {
	X string `json:"x,omitempty"`
	Y string `json:"y,omitempty"`
}

type Sensor struct {
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	ApiKey      string      `json:"api_key,omitempty"`
	Value       interface{} `json:"value,omitempty"`
	Unit        string      `json:"unit,omitempty"`
	SetPoint    interface{} `json:"set_point,omitempty"`
	Histerezis  interface{} `json:"histerezis,omitempty"`
	Divisor     int         `json:"divisor,omitempty"`
	Multiplier  int         `json:"multiplier,omitempty"`
	Ratio       string      `json:"ratio,omitempty"`
	Residual    float64     `json:"residual,omitempty"`
	Coordinates
	Flags
}

var divisors = map[string]int{U_ENERGY: 1000, U_CURRENT: 100, U_TEMP: 10, U_HUMID: 10, U_VOLTAGE: 10, U_DUSK: 1000}

// energy meters CT rates in Adeleq format (http://www.adeleq.gr)
var ct = map[string]int{
	"5-5":    1,
	"50-5":   10,
	"65-5":   13,
	"75-5":   15,
	"100-5":  20,
	"125-5":  25,
	"150-5":  30,
	"160-5":  32,
	"200-5":  40,
	"250-5":  50,
	"300-5":  60,
	"400-5":  80,
	"500-5":  100,
	"600-5":  120,
	"750-5":  150,
	"800-5":  160,
	"1000-5": 200,
	"1200-5": 240,
	"1250-5": 250,
	"1500-5": 300,
	"2000-5": 400,
	"2400-5": 480,
	"2500-5": 500,
	"3000-5": 600,
	"4000-5": 800,
	"5000-5": 1000,
	"6000-5": 1200,
	"7500-5": 1500,
}

// Return standard multiplier or specified.
func (s *Sensor) SetMultiplier() int {
	// Multiplier have sens only for indirect energy meters.
	if s.Multiplier = ct[s.Ratio]; s.Multiplier != 0 && s.Unit == U_ENERGY {
		return s.Multiplier
	}
	return 1
}

// Return standard divisor or specified.
func (s *Sensor) SetDivisor() int {
	// value of 0 mean that isn't specified.
	if s.Divisor == 0 {
		if s.Divisor = divisors[s.Unit]; s.Divisor != 0 {
			return s.Divisor
		} else {
			return 1
		}
	}
	return s.Divisor
}
