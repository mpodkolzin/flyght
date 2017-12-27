package main

import (
	"encoding/json"
	"flyght/adsb"
	"flyght/publisher"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world")
	w.Write([]byte("Hello world"))
	w.WriteHeader(http.StatusOK)

}

const (
	adsbExchangeURL = "http://public-api.adsbexchange.com/VirtualRadar/AircraftList.json?lat=33.433638&lng=-112.008113&fDstL=0&fDstU=100"
)

func main() {
	fmt.Println("Reading plane list")
	resp, err := http.Get(adsbExchangeURL)

	if err != nil {

		panic(err)

	} else {

		fmt.Println("Error occured")
		fmt.Println("Error occ")

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var data adsb.AdsbResponse

	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err.Error())
		panic(err)

	} else {
	}

	//p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "10.4.200.9:9092"})
	p, err := publisher.New()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(err)
	}

	topic := "test_go"

	//fmt.Println(data["acList"])
	for _, ac := range data.AcList {

		acJson, _ := json.Marshal(&ac)

		fmt.Println("----------------------------Sending message to kafka queue-------------------------------")

		p.Producer.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(acJson)}
		fmt.Println(ac)
	}
	p.Producer.Close()

}
