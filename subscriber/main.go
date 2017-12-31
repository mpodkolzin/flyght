package main

import (
	//"github.com/olivere/elastic"
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gopkg.in/olivere/elastic.v5"
)

const (
	timeout = 100
)

var (
	config = &kafka.ConfigMap{
		//"bootstrap.servers": "10.4.200.9",
		"bootstrap.servers": "127.0.0.1:9092",
		"group.id":          "my_group2",
		"auto.offset.reset": "earliest",
	}
	elasticUrl   = "http://172.18.0.4:9200"
	elasticIndex = "adsb"
	topics       = []string{"adsb_topic"}
)

func NewElasticClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL(elasticUrl))

	if err != nil {
		log.Fatal("could not connect to elastic search client")
		return nil, err
	}

	return client, nil

}

func IndexDoc(client *elastic.Client, doc interface{}) error {

	_, err := client.Index().
		Index(elasticIndex).
		Type(elasticIndex).
		BodyJson(doc).
		Do(context.Background())
	if err != nil {
		log.Print(err)
		log.Println("cound not index the doc")
		return err
	}

	return nil
}

func CreateIndex(client *elastic.Client) bool {

	if exists, _ := client.IndexExists(elasticIndex).Do(context.Background()); exists {

		return true

	} else {

		//TODO add error handling here
		client.CreateIndex(elasticIndex).Do(context.Background())
		return true
	}

}

func main() {

	c, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatal("cannot create kafka subscriber")
	}

	elClient, _ := NewElasticClient()
	CreateIndex(elClient)
	//TODO error handling

	err = c.SubscribeTopics(topics, nil)
	for {

		ev := c.Poll(timeout)
		if ev == nil {
			continue
		}

		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("Received message, %s\n", string(e.Value))
			IndexDoc(elClient, string(e.Value))

		case *kafka.PartitionEOF:
		case *kafka.Error:
			break
		}

	}

}
