// +build linux darwin !appengine

package tiny

// Runtime structure for v3 controllers.
type ix struct {
	Tem     float64 `xml:"tem,string"`     // onboard temp sensor
	Vin     float64 `xml:"vin,string"`     // onboard power supply
	Ds1     float64 `xml:"ds1,string"`     // DS18B20
	Ds2     float64 `xml:"ds2,string"`     // DS18B20
	Ds3     float64 `xml:"ds3,string"`     // DS18B20
	Ds4     float64 `xml:"ds4,string"`     // DS18B20
	Ds5     float64 `xml:"ds5,string"`     // DS18B20
	Ds6     float64 `xml:"ds6,string"`     // DS18B20
	Ind     int     `xml:"ind,string"`     // digital inputs bitmask Inpd4=7,3=11,2=13,1=14 else 15
	Dth0    float64 `xml:"dth0,string"`    //
	Dth1    float64 `xml:"dth1,string"`    //
	Bm280p  float64 `xml:"bm280p,string"`  //
	Energy1 float64 `xml:"energy1,string"` // energy input 1
	Energy2 float64 `xml:"energy2,string"` // energy input 2
	Energy3 float64 `xml:"energy3,string"` // energy input 3
	Energy4 float64 `xml:"energy4,string"` // energy input 4
	Pm10    float64 `xml:"pm10,string"`    // air quality sensor
	Pm25    float64 `xml:"pm25,string"`    // air quality sensor
	Inp1    float64 `xml:"inp1,string"`    // INPA1
	Inp2    float64 `xml:"inp2,string"`    // INPA2
	Inp3    float64 `xml:"inp3,string"`    // INPA3
	Inp4    float64 `xml:"inp4,string"`    // INPA4
	Inp5    float64 `xml:"inp5,string"`    // INPA5
	Inp6    float64 `xml:"inp6,string"`    // INPA6
	Inpp1   float64 `xml:"inpp1,string"`   //
	Inpp2   float64 `xml:"inpp2,string"`   //
	Inpp3   float64 `xml:"inpp3,string"`   //
	Inpp4   float64 `xml:"inpp4,string"`   //
}

// Status structure for v3 controllers. Just what's in addition to v2!
type stat struct {
	Sname string `xml:"sname"` // board name
	Sw    string `xml:"sw"`    // software version
}

// Configuration structure for v3 controllers.
type st struct {
	HTTPTime int `xml:"ht3,string"`
	DsID     int `xml:"dsid,string"`
}

// Third generation controller.
type Lk3 struct {
	ix
	stat
	//st
}

func (c *Lk3) GetName() string {
	return c.Sname
}

func (c *Lk3) GetVersion() string {
	return c.Sw
}
