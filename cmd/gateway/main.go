package main

import (
	"flag"
	"github.com/medium-stories/go-rabbitmq/cmd/gateway/publisher"
	"github.com/medium-stories/go-rabbitmq/cmd/gateway/repository"
	"github.com/medium-stories/go-rabbitmq/internal/web"
	"github.com/medium-stories/go-rabbitmq/order"
)

var (
	httpAddr = flag.String("http", ":8000", "Http address")
)

func main() {
	flag.Parse()

	router := web.NewRouter()

	orderApi := order.NewApi(order.NewService(repository.NewInMemory(), publisher.NewMocked()))

	router.POST("order", orderApi.CreateHandler())
	router.GET("order/:identifier", orderApi.GetHandler())

	web.ServeHttp(*httpAddr, "gateway", router)
}
