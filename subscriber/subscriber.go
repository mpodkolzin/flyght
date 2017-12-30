package main

//Subscriber : interface for subscribing to changes for ADSB feed
type Subscriber interface {
	Subscribe() (chan interface{}, error)
}
