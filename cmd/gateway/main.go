package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobackpack/rmq"
	orderPublisher "github.com/medium-stories/go-rabbitmq/event/publishers/order"
	"github.com/medium-stories/go-rabbitmq/internal/web"
	"github.com/medium-stories/go-rabbitmq/order"
	"github.com/medium-stories/go-rabbitmq/order/repository"
	"github.com/medium-stories/go-rabbitmq/payment"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	httpAddr = flag.String("http", ":8000", "Http address")
	rmqHost  = flag.String("rmq_host", "localhost", "RabbitMQ host address")
)

func main() {
	flag.Parse()

	cred := rmq.NewCredentials()
	cred.Host = *rmqHost
	hub := rmq.NewHub(cred)

	hubCtx, hubCancel := context.WithCancel(context.Background())
	defer hubCancel()

	_, err := hub.Connect(hubCtx)
	if err != nil {
		logrus.Fatal(err)
	}

	router := web.NewRouter()

	pub := orderPublisher.NewOrderPublisher(hubCtx, hub)
	orderApi := order.NewApi(order.NewService(repository.NewInMemory(), pub))
	paymentGateway := payment.NewPaymentGateway(pub)

	router.POST("order", orderApi.CreateHandler())
	router.GET("order/:identifier", orderApi.GetHandler())
	router.POST("order/:identifier/checkout", func(c *gin.Context) {
		identifier := c.Param("identifier")
		if err := paymentGateway.Pay(identifier); err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, fmt.Sprintf("order %s paid", identifier))
	})

	web.ServeHttp(*httpAddr, "gateway", router)
}
