package main

import (
	"example.com/m/v4/internal/clients"
	"example.com/m/v4/internal/config"
	"example.com/m/v4/internal/db"
	"example.com/m/v4/internal/handlers"
	orderitems "example.com/m/v4/internal/models/orderItems"
	ordersModel "example.com/m/v4/internal/models/orders"
	"example.com/m/v4/internal/repository"
	"example.com/m/v4/internal/routes"
	"example.com/m/v4/internal/services"
	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "server is healthy",
	})
}
func main() {
	database := db.SetupDB()
	err := database.AutoMigrate(
		&ordersModel.Order{},
		&orderitems.OrderItem{},
	)
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.GET("/health", health)
	//order routes
	newOrderRepository := repository.NewOrderRepository(database)
	catalogService := clients.NewCatalogHTTPClient("http://localhost:8080")
	newOrderService := services.NewOrderService(newOrderRepository, catalogService)
	newOrderHandler := handlers.NewOrderHandlers(newOrderService)
	routes.SetupRoutes(router, newOrderHandler)
	router.Run(":" + config.ConfigLoadEnv().PORT)

}
