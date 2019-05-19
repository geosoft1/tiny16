// +build linux darwin !appengine

package tiny

// Runtime structure for v2 controllers.
type st0 struct {
	Temp   float64 `xml:"ia0,string"`  // onboard temp sensor
	Vcc    float64 `xml:"ia1,string"`  // onboard power supply
	Inp6   float64 `xml:"ia7,string"`  // DS18B20
	Inp7   float64 `xml:"ia8,string"`  // DS18B20
	Inp8   float64 `xml:"ia9,string"`  // DS18B20
	Inp9   float64 `xml:"ia10,string"` // DS18B20
	Inp10  float64 `xml:"ia11,string"` // DS18B20
	Inp11  float64 `xml:"ia12,string"` // DS18B20
	DTH22t float64 `xml:"ia13,string"` //
	DTH22h float64 `xml:"ia14,string"` //
	DIFF   float64 `xml:"ia19,string"` //
	Di0    string  `xml:"di0"`         // digital input 0
	Di1    string  `xml:"di1"`         // digital input 1
	Di2    string  `xml:"di2"`         // digital input 2
	Di3    string  `xml:"di3"`         // digital input 3
	INP4D  float64 `xml:"ia17,string"` // energy input
}

// Status structure for v2 controllers.
type st2 struct {
	Na  string `xml:"na"`  // board name
	Ver string `xml:"ver"` // software version
}

// Configuration structure for v2 controllers.
type board struct {
	HTTP_time int `xml:"e3,string"`
	DsID      int `xml:"ds,string"`
}

// Second generation controller.
type Lk2 struct {
	st0
	st2
	board
}

func (c *Lk2) GetName() string {
	return c.Na
}

func (c *Lk2) GetVersion() string {
	return c.Ver
}
