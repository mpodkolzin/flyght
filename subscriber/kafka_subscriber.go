package subscriber

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaSubscriber struct {
	Consumer *kafka.Consumer
}

func New() (*KafkaSubscriber, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "10.4.200.9",
		"group.id":          "my_group",
	}
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}

	return &KafkaSubscriber{Consumer: consumer}, nil
}

//Subscribe : subscribe to changes in ADSB feed
func Subscribe() (chan interface{}, error) {
	return nil, nil

}
