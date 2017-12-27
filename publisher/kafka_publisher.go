package publisher

import (
	kafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

//KafkaPublisher is an implementation of published interface : Interface to publish plane data to message bus
type KafkaPublisher struct {
	Producer *kafka.Producer
}

func New() (*KafkaPublisher, error) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "10.4.200.9:9092"})
	if err != nil {
		return nil, err
	}
	return &KafkaPublisher{Producer: p}, nil
}

func (kp *KafkaPublisher) Publish(message interface{}) {

}
