package payment

import (
	"github.com/medium-stories/go-rabbitmq/event"
	"github.com/sirupsen/logrus"
)

type gateway struct {
	pub event.Publisher
}

func NewPaymentGateway(pub event.Publisher) *gateway {
	return &gateway{
		pub: pub,
	}
}

func (gtw *gateway) Pay(orderId string) error {
	logrus.Infof("payment for order id: %s completed", orderId)

	go func(oId string) {
		if err := gtw.pub.Publish(event.OrderPaid, oId); err != nil {
			logrus.Error(err)
		}
	}(orderId)

	return nil
}
