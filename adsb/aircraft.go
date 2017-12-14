package adsb

type Aircraft struct {
	Id           int
	Rcvr         int
	HasSig       bool
	Icao         string
	Bad          bool
	Reg          string
	FSeen        string
	TSecs        float32
	CMsgs        float32
	Alt          float32
	GAlt         float32
	InHg         float32
	AltT         float32
	Call         string
	Lat          float32
	Long         float32
	PosTime      uint64
	Mlat         bool
	Tisb         bool
	Spd          float32
	Trak         float32
	TrkH         bool
	Type         string
	Mdl          string
	Man          string
	CNun         string
	Op           string
	OpIcao       string
	Sqk          string
	Help         bool
	Vsi          int //???
	VsiT         int //???
	Dst          float32
	Brng         float32
	WTC          int
	Species      float32
	Engines      string
	EngType      int
	EngMount     int
	Mil          bool
	Cou          string
	HasPic       bool
	Interested   bool
	FlightsCount int
	Gnd          bool
	SpdTyp       int
	CallSus      bool
	Trt          int
	Year         string
}
