package main

import (
	"example.com/m/v4/internal/config"
	"example.com/m/v4/internal/db"
	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "server is healthy",
	})
}
func main() {
	database := db.SetupDB()
	err := database.AutoMigrate()
	if err != nil {
		panic(err)

	}
	router := gin.Default()
	router.GET("/health", health)
	router.Run(":" + config.ConfigLoadEnv().PORT)

}
