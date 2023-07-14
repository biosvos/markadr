package broker

type Broker interface {
	Publish(topic string, message string) error
}
