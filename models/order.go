package models

type Order struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	CustomerID uint        `json:"customer_id"`
	Status     string      `json:"status"`
	Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}
