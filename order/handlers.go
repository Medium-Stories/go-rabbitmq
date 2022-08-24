package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type api struct {
	svcAddr string
}

func NewApi(orderServiceAddr string) *api {
	return &api{
		svcAddr: orderServiceAddr,
	}
}

func (api *api) CreateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "order created")
	}
}

func (api *api) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "order details")
	}
}
