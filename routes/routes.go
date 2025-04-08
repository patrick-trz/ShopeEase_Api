package routes

import (
	"ShopEase_API/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	product := r.Group("/products")
	{
		product.GET("/", controllers.GetProducts)
		product.POST("/", controllers.CreateProduct)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
	}

	order := r.Group("/orders")
	{
		order.GET("/", controllers.GetOrders)
		order.POST("/", controllers.CreateOrder)
		order.PUT("/:id", controllers.UpdateOrder)
		order.DELETE("/:id", controllers.DeleteOrder)
	}
}
