package routes

import (
	"ShopEase_API/controllers"
	"ShopEase_API/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	product := r.Group("/products")
	product.Use(middleware.AuthMiddleware())
	{
		product.GET("/", controllers.GetProducts)
		product.GET("/:id", controllers.GetProduct)
		product.POST("/", controllers.CreateProducts)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
	}

	order := r.Group("/orders")
	order.Use(middleware.AuthMiddleware())
	{
		order.GET("/dashboard", controllers.Dashboard) // ✅ dashboard (current user info)
		order.GET("/", controllers.GetOrders)
		order.GET("/:id", controllers.GetOrder)
		order.POST("/", controllers.CreateOrder)
		order.PUT("/:id", controllers.UpdateOrder)
		order.DELETE("/:id", controllers.DeleteOrder)
	}
}
