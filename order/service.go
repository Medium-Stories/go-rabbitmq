package order

import "github.com/medium-stories/go-rabbitmq/event"

type service struct {
	pub event.Publisher
}

func (svc *service) Create(item Item) error {
	svc.pub.Publish(event.OrderCreated, item.Id)
	return nil
}
