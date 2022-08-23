package shipping

import "github.com/medium-stories/go-rabbitmq/event"

type local struct {
	pub event.Publisher
}

func (loc *local) Ship(orderId string) error {
	loc.pub.Publish(event.OrderShipped, orderId)
	return nil
}
