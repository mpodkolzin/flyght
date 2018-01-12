package main

import (
	"encoding/json"
	"flyght/adsb"
	"flyght/publisher"
	"flyght/publisher/kafkaPublisher"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}

var latLonPartitions []string = []string{
	"lat=33.433638&lng=-112.008113",
}

const (
	adsbExchangeURL = "http://public-api.adsbexchange.com/VirtualRadar/AircraftList.json?lat=33.433638&lng=-112.008113&fDstL=0&fDstU=100"
	adsbBaseURL     = "http://public-api.adsbexchange.com/VirtualRadar/AircraftList.json"
	adsbTcpEndpoint = "pub-vrs.adsbexchange.com:32030"
	adsbTopic       = "adsb_topic"
)

func main() {

	port := os.Getenv("CRAWLER_PORT")
	if port == "" {
		port = "8080"
	}

	adsbPublisher, err := kafkaPublisher.NewPublisher(adsbTopic)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(err)
	}
	go crawlTcp(adsbPublisher)

	defer adsbPublisher.Producer.Close()

	http.HandleFunc("/ping", PingHandler)
	http.ListenAndServe(":"+port, nil)

}

func crawlTcp(publiser publisher.Publisher) error {

	conn, err := net.Dial("tcp", adsbTcpEndpoint)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	d := json.NewDecoder(conn)

	var msg adsb.AdsbResponse
	err = d.Decode(&msg)

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	for _, ac := range msg.AcList {
		fmt.Printf("Icao: %s, lat: %f \n", ac.Icao, ac.Lat)
	}

	fmt.Println("Ac count: ", len(msg.AcList))

	return nil

}

func crawlHttp(publisher publisher.Publisher) error {

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
		return err
	}

	for _, ac := range data.AcList {
		acJson, _ := json.Marshal(&ac)
		fmt.Println("----------------------------Sending message to kafka queue-------------------------------")
		publisher.Publish(acJson)
	}

	return nil

}
