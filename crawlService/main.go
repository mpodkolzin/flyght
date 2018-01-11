package main

import (
	"encoding/json"
	"flyght/adsb"
	"flyght/publisher"
	"flyght/publisher/kafkaPublisher"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world")
	w.Write([]byte("Hello world"))
	w.WriteHeader(http.StatusOK)

}

var latLonPartitions []string {
	"lat=33.433638&lng=-112.008113"
}


const (
	adsbExchangeURL = "http://public-api.adsbexchange.com/VirtualRadar/AircraftList.json?lat=33.433638&lng=-112.008113&fDstL=0&fDstU=100"
	adsbBaseURL     = "http://public-api.adsbexchange.com/VirtualRadar/AircraftList.json"
	adsbTcpEndpoint = "pub-vrs.adsbexchange.com:32030"
)


func main() {

	topic := "adsb_topic"
	adsbPublisher, err := publisher.NewPublisher(topic)
	defer adsbPublisher.Pro
	if err != nil {
		panic(err)
	} else {
		fmt.Println(err)
	}


}


func crawlTcp(publiser Publisher) error {

	conn, err := net.Dial("tcp", adsbTcpEndpoint)
	if err != nil {
		fmt.Println(err1.Error())
	}
	defer conn.Close()

	d := json.NewDecoder(conn)

    var msg adsb.AdsbResponse
	err := d.Decode(&msg)

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	for _, ac := range msg.AcList {
		fmt.Printf("Icao: %s, lat: %f \n", ac.Icao, ac.Lat);
	}

	fmt.Println("Ac count: ", len(msg.AcList))

	return nil

}

func crawlHttp(publisher Publisher) error {

	fmt.Println("Reading plane list")
	resp, err := http.Get(adsbExchangeURL)

	if err != nil {
		log.Fatal("Could not read adsb stream")
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var data adsb.AdsbResponse

	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err.Error())
		return err;
	}

	for _, ac := range data.AcList {
		acJson, _ := json.Marshal(&ac)
		fmt.Println("----------------------------Sending message to kafka queue-------------------------------")
		p.Publish(acJson)
	}

	return nil

}
