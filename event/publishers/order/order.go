package order

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gobackpack/rmq"
	"github.com/medium-stories/go-rabbitmq/event"
)

type orderPublisher struct {
	hub  *rmq.Hub
	conf map[string]*rmq.Publisher
}

func NewOrderPublisher(ctx context.Context, hub *rmq.Hub) *orderPublisher {
	pub := &orderPublisher{
		hub:  hub,
		conf: make(map[string]*rmq.Publisher),
	}

	pub.setupEvents(ctx, []string{event.OrderCreated, event.OrderPaid, event.OrderShipped})

	return pub
}

func (pub *orderPublisher) Publish(event string, msg interface{}) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	rmqPub := pub.conf[event]
	if rmqPub == nil {
		return errors.New("invalid account event")
	}

	rmqPub.Publish(b)

	return nil
}

func (pub *orderPublisher) setupEvents(ctx context.Context, events []string) {
	for _, ev := range events {
		conf := rmq.NewConfig()
		conf.Exchange = event.WooHooStoreBus
		conf.RoutingKey = ev
		conf.ExchangeKind = "topic"

		pub.conf[ev] = pub.hub.CreatePublisher(ctx, conf)
	}
}
