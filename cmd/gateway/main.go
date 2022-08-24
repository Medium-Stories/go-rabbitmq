package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/medium-stories/go-rabbitmq/cmd/gateway/publisher"
	"github.com/medium-stories/go-rabbitmq/cmd/gateway/repository"
	"github.com/medium-stories/go-rabbitmq/internal/web"
	"github.com/medium-stories/go-rabbitmq/order"
	"github.com/medium-stories/go-rabbitmq/payment"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	httpAddr = flag.String("http", ":8000", "Http address")
)

func main() {
	flag.Parse()

	router := web.NewRouter()

	orderApi := order.NewApi(order.NewService(repository.NewInMemory(), publisher.NewMocked()))
	paymentGateway := payment.NewPaymentGateway(publisher.NewMocked())

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
