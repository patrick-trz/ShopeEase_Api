package main

import (
	"ShopEase_API/config"
	"ShopEase_API/models"
	"ShopEase_API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	// Миграция
	config.DB.AutoMigrate(
		&models.Product{},
		&models.Customer{},
		&models.Order{},
		&models.OrderItem{},
	)

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
