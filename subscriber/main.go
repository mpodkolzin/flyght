package main

import (
	//"github.com/olivere/elastic"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "gopkg.in/olivere/elastic.v5"
)

const (
	timeout = 100
)

var (
	config = &kafka.ConfigMap{
		"bootstrap.servers": "10.4.200.9",
		"group.id":          "my_group",
	}
	topics = []string{"adsb_topic"}
)

func main() {

	c, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatal("cannot create kafka subscriber")
	}

	err = c.SubscribeTopics(topics, nil)
	for {

		ev := c.Poll(timeout)
		if ev == nil {
			continue
		}

		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Println("Received message, %s\n", string(e.Value))

		case *kafka.PartitionEOF:
		case *kafka.Error:
			break
		}

	}

}
