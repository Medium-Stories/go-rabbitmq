package main

import (
	"flag"
	"github.com/medium-stories/go-rabbitmq/internal/web"
	"github.com/medium-stories/go-rabbitmq/order"
)

var (
	httpAddr  = flag.String("http", ":8000", "Http address")
	orderAddr = flag.String("order_uri", ":8001", "Order Service address")
)

func main() {
	flag.Parse()

	router := web.NewRouter()

	orderApi := order.NewApi(*orderAddr)

	router.POST("order", orderApi.CreateHandler())
	router.GET("order/:identifier", orderApi.GetHandler())

	web.ServeHttp(*httpAddr, "gateway", router)
}
