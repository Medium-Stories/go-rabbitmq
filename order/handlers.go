package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type api struct {
	svc *service
}

func NewApi(svc *service) *api {
	return &api{
		svc: svc,
	}
}

func (api *api) CreateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := uuid.NewString()

		if err := api.svc.Create(&Bucket{
			Identifier: orderId,
		}); err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, fmt.Sprintf("order created: %s", orderId))
	}
}

func (api *api) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		identifier := c.Param("identifier")
		c.JSON(http.StatusOK, api.svc.CheckStatus(identifier))
	}
}
