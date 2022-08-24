package shipping

import (
	"github.com/medium-stories/go-rabbitmq/event"
	"github.com/sirupsen/logrus"
)

type local struct {
	pub event.Publisher
}

func NewShippingMethod(pub event.Publisher) *local {
	return &local{
		pub: pub,
	}
}

func (loc *local) Ship(orderId string) error {
	logrus.Infof("order id: %s shipped", orderId)

	go func(oId string) {
		if err := loc.pub.Publish(event.OrderShipped, oId); err != nil {
			logrus.Error(err)
		}
	}(orderId)

	return nil
}
