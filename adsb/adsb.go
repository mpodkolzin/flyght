package adsb

type AdsbResponse struct {
	Src     int
	SrcFeed int
	AcList  []Aircraft
}
