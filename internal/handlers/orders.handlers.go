package handlers

import (
	"net/http"

	"example.com/m/v4/internal/services"
	"github.com/gin-gonic/gin"
)

type OrderHandlers struct {
	orderService *services.OrderService
}

func NewOrderHandlers(orderService *services.OrderService) *OrderHandlers {
	return &OrderHandlers{orderService: orderService}
}

func (o *OrderHandlers) CreateOrder(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"msg": "hii",
	})
}
