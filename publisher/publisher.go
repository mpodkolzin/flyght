package publisher

//Publisher : Interface to publish plane data to message bus
type Publisher interface {
	Publish(message interface{})
}
