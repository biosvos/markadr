package mqtt

import (
	"fmt"
	"github.com/biosvos/markadr/flow/broker"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
)

var _ broker.Broker = &Broker{}

func NewBroker() (*Broker, error) {
	options := MQTT.NewClientOptions()
	options.AddBroker(fmt.Sprintf("ws://%v:%v", "127.0.0.1", "9001"))
	client := MQTT.NewClient(options)
	token := client.Connect()
	token.Wait()
	if token.Error() != nil {
		return nil, errors.WithStack(token.Error())
	}
	return &Broker{
		client: client,
	}, nil
}

type Broker struct {
	client MQTT.Client
}

func (b *Broker) Publish(topic string, message string) error {
	token := b.client.Publish(topic, 0, false, message)
	token.Wait()
	if token.Error() != nil {
		return errors.WithStack(token.Error())
	}
	return nil
}
