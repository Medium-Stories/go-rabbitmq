package event

import (
	"context"
)

const (
	WooHooStoreBus = "woohoo_store"
	OrderCreated   = "order_created"
	OrderPaid      = "order_paid"
	OrderShipped   = "order_shipped"
)

type Publisher interface {
	Publish(event string, payload interface{}) error
}

type Listener interface {
	Listen(ctx context.Context)
}

func Listen(ctx context.Context, listeners ...Listener) {
	for _, listener := range listeners {
		listener.Listen(ctx)
	}
}
