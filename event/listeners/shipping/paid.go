package shipping

import (
	"context"
	"fmt"
	"github.com/gobackpack/rmq"
	"github.com/medium-stories/go-rabbitmq/event"
	"github.com/medium-stories/go-rabbitmq/event/listeners"
	"github.com/medium-stories/go-rabbitmq/order"
	"github.com/sirupsen/logrus"
	"time"
)

type orderPaid struct {
	hub  *rmq.Hub
	repo order.Repository
}

func NewOrderPaidListener(hub *rmq.Hub, repo order.Repository) *orderPaid {
	return &orderPaid{
		hub:  hub,
		repo: repo,
	}
}

func (ev *orderPaid) Listen(ctx context.Context) {
	consumer := listeners.StartConsumer(ctx, ev.hub, event.OrderPaid)
	go ev.handleMessages(ctx, consumer, fmt.Sprintf("shipping[%s]", event.OrderPaid))
}

func (ev *orderPaid) handleMessages(ctx context.Context, cons *rmq.Consumer, name string) {
	logrus.Infof("%s started", name)

	defer logrus.Warnf("%s closed", name)

	for {
		select {
		case msg := <-cons.OnMessage:
			logrus.Infof("[%s] %s - %s", time.Now().UTC(), name, msg)
		case err := <-cons.OnError:
			logrus.Error(err)
		case <-ctx.Done():
			return
		}
	}
}
