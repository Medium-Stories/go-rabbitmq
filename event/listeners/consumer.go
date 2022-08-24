package listeners

import (
	"context"
	"fmt"
	"github.com/gobackpack/rmq"
	"github.com/medium-stories/go-rabbitmq/event"
	"github.com/sirupsen/logrus"
)

func StartConsumer(ctx context.Context, hub *rmq.Hub, svc, ev string) *rmq.Consumer {
	conf := rmq.NewConfig()
	conf.Exchange = event.WooHooStoreBus
	conf.Queue = fmt.Sprintf("%s@%s", svc, ev)
	conf.ConsumerTag = fmt.Sprintf("%s@%s", svc, ev)
	conf.RoutingKey = ev
	conf.ExchangeKind = "topic"

	if err := hub.CreateQueue(conf); err != nil {
		logrus.Fatal(err)
	}

	return hub.StartConsumer(ctx, conf)
}
