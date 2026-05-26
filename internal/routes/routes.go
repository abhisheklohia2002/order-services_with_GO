package routes

import (
	"example.com/m/v4/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	OrderHandler *handlers.OrderHandlers
}

func SetupRoutes(router *gin.Engine, OrderHandler *handlers.OrderHandlers) {

	api := router.Group("/api")
	orders := api.Group("/order")
	{
		orders.POST("/order", OrderHandler.CreateOrder)
	}

}
