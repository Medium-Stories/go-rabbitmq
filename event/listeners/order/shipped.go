package order

import (
	"context"
	"encoding/json"
	"github.com/gobackpack/rmq"
	"github.com/medium-stories/go-rabbitmq/event"
	"github.com/medium-stories/go-rabbitmq/event/listeners"
	"github.com/medium-stories/go-rabbitmq/order"
	"github.com/sirupsen/logrus"
)

type orderShipped struct {
	hub  *rmq.Hub
	repo order.Repository
}

func NewOrderShippedListener(hub *rmq.Hub, repo order.Repository) *orderShipped {
	return &orderShipped{
		hub:  hub,
		repo: repo,
	}
}

func (ev *orderShipped) Listen(ctx context.Context) {
	consumer := listeners.StartConsumer(ctx, ev.hub, "order", event.OrderShipped)
	ev.handleMessages(ctx, consumer, event.OrderShipped)
}

func (ev *orderShipped) handleMessages(ctx context.Context, cons *rmq.Consumer, name string) {
	logrus.Infof("%s started", name)

	defer logrus.Warnf("%s closed", name)

	for {
		select {
		case msg := <-cons.OnMessage:
			var identifier string
			json.Unmarshal(msg, &identifier)

			logrus.Infof("[%s] - %s", name, identifier)

			if err := ev.repo.UpdateStatus(identifier, order.StatusShipped); err != nil {
				logrus.Error(err)
			}
		case err := <-cons.OnError:
			logrus.Error(err)
		case <-ctx.Done():
			return
		}
	}
}
