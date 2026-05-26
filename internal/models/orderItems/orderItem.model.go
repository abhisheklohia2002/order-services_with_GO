package orderitems

import "time"
type OrderItem struct {
	ID      uint `json:"id" gorm:"primaryKey"`
	OrderID uint `json:"orderId" gorm:"not null;index"`

	ProductID uint `json:"productId" gorm:"not null"`

	ProductName  string `json:"productName" gorm:"type:varchar(150);not null"`
	ProductSKU   string `json:"productSku" gorm:"type:varchar(100);not null"`
	ProductImage string `json:"productImage" gorm:"type:text"`

	UnitPrice     int `json:"unitPrice" gorm:"not null"`
	DiscountPrice int `json:"discountPrice" gorm:"default:0"`
	Quantity      int `json:"quantity" gorm:"not null"`
	Total         int `json:"total" gorm:"not null"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}