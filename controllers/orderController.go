package controllers

import (
	"ShopEase_API/config"
	"ShopEase_API/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateOrder handles the POST request to create a new order
func CreateOrder(c *gin.Context) {
	var order models.Order
	// Bind JSON body to order struct
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create new order in the database
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GetOrders handles the GET request to fetch all orders
func GetOrders(c *gin.Context) {
	var orders []models.Order
	// Preload related data (e.g., items in the order) and find all orders
	if err := config.DB.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetOrder handles the GET request to fetch an order by ID
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	// Find the order by ID
	if err := config.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// UpdateOrder handles the PUT request to update an order
func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	// Find the order by ID
	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Bind new data to order struct
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the order
	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// DeleteOrder handles the DELETE request to delete an order
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	// Delete the order by ID
	if err := config.DB.Delete(&models.Order{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
