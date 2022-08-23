package payment

import "github.com/medium-stories/go-rabbitmq/event"

type gateway struct {
	pub event.Publisher
}

func (gtw *gateway) Pay(orderId string) error {
	gtw.pub.Publish(event.OrderPaid, orderId)
	return nil
}
