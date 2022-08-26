package shipping

import (
	"github.com/medium-stories/go-rabbitmq/event"
	"github.com/sirupsen/logrus"
)

type localMethod struct {
	pub event.Publisher
}

func NewShippingMethod(pub event.Publisher) *localMethod {
	return &localMethod{
		pub: pub,
	}
}

func (loc *localMethod) Ship(orderId string) error {
	logrus.Infof("order id: %s shipped", orderId)

	go func(oId string) {
		if err := loc.pub.Publish(event.OrderShipped, oId); err != nil {
			logrus.Error(err)
		}
	}(orderId)

	return nil
}
