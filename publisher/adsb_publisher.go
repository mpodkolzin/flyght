package adsb_publisher

type Publisher interface {
	Publish(message interface{})
}

type KafkaPublisher struct {
}

func (kp *KafkaPublisher) Publish(message interface{}) {

}
