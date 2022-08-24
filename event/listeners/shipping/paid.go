package shipping

import (
	"context"
	"fmt"
	"github.com/gobackpack/rmq"
	"github.com/medium-stories/go-rabbitmq/event"
	"github.com/medium-stories/go-rabbitmq/event/listeners"
	"github.com/sirupsen/logrus"
	"time"
)

type orderPaid struct {
	hub    *rmq.Hub
	method Method
}

type Method interface {
	Ship(orderId string) error
}

func NewOrderPaidListener(hub *rmq.Hub, method Method) *orderPaid {
	return &orderPaid{
		hub:    hub,
		method: method,
	}
}

func (ev *orderPaid) Listen(ctx context.Context) {
	consumer := listeners.StartConsumer(ctx, ev.hub, "shipping", event.OrderPaid)
	ev.handleMessages(ctx, consumer, fmt.Sprintf("shipping[%s]", event.OrderPaid))
}

func (ev *orderPaid) handleMessages(ctx context.Context, cons *rmq.Consumer, name string) {
	logrus.Infof("%s started", name)

	defer logrus.Warnf("%s closed", name)

	for {
		select {
		case msg := <-cons.OnMessage:
			logrus.Infof("[%s] %s - %s", time.Now().UTC(), name, msg)

			if err := ev.method.Ship(string(msg)); err != nil {
				logrus.Error(err)
			}
		case err := <-cons.OnError:
			logrus.Error(err)
		case <-ctx.Done():
			return
		}
	}
}
