package main

import (
	"./adsb"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world")
	w.Write([]byte("Hello world"))
	w.WriteHeader(http.StatusOK)

}

func main() {
	url := "http://public-api.adsbexchange.com/VirtualRadar/AircraftList.json?lat=33.433638&lng=-112.008113&fDstL=0&fDstU=100"
	fmt.Println("Reading plane list")
	resp, err := http.Get(url)

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

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "10.4.200.9:9092"})
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

		p.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(acJson)}
		fmt.Println(ac)
	}
	p.Close()

}
