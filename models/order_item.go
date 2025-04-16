package models

type OrderItem struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
	//quantity - how much?
}

// (3, 3, 2)

//order_id = 3 — 3-нөмірлі заказ
//product_id = 3 — 3-нөмірлі өнім (мысалы, "Headphones")
//quantity = 2 — клиент бұл өнімді 2 дана алған

//order_items — бұл "аралық кесте" (many-to-many қатынас)
//Бір заказда бірнеше өнім болуы мүмкін
//Бір өнім бірнеше заказда кездесуі мүмкін
