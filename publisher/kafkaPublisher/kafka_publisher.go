package kafkaPublisher

import (
	"fmt"

	kafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

//KafkaPublisher is an implementation of published interface : Interface to publish plane data to message bus
type KafkaPublisher struct {
	Producer *kafka.Producer
	Topic    string
}

func NewPublisher(topic string, bootstrapServers string) (*KafkaPublisher, error) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		return nil, err
	}
	return &KafkaPublisher{Producer: p, Topic: topic}, nil
}

func (kp *KafkaPublisher) Publish(message []byte) {

	fmt.Printf("producing message, topic %s \n", kp.Topic)
	kp.Producer.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &kp.Topic, Partition: kafka.PartitionAny}, Value: []byte(message)}

}
