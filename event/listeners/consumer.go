package listeners

import (
	"context"
	"github.com/gobackpack/rmq"
	"github.com/medium-stories/go-rabbitmq/event"
	"github.com/sirupsen/logrus"
)

func StartConsumer(ctx context.Context, hub *rmq.Hub, ev string) *rmq.Consumer {
	conf := rmq.NewConfig()
	conf.Exchange = event.WooHooStoreBus
	conf.Queue = ev
	conf.RoutingKey = ev

	if err := hub.CreateQueue(conf); err != nil {
		logrus.Fatal(err)
	}

	return hub.StartConsumer(ctx, conf)
}
