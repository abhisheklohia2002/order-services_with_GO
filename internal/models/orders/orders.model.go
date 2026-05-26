package ordersModel

import (
	"time"

	"example.com/m/v4/internal/enums"
	orderitems "example.com/m/v4/internal/models/orderItems"
)

type OrderStatus string
type PaymentStatus string

const (
	OrderPendingPayment OrderStatus = "PENDING_PAYMENT"
	OrderConfirmed      OrderStatus = "CONFIRMED"
	OrderProcessing     OrderStatus = "PROCESSING"
	OrderShipped        OrderStatus = "SHIPPED"
	OrderDelivered      OrderStatus = "DELIVERED"
	OrderCancelled      OrderStatus = "CANCELLED"
	OrderPaymentFailed  OrderStatus = "PAYMENT_FAILED"

	PaymentPending  PaymentStatus = "PENDING"
	PaymentPaid     PaymentStatus = "PAID"
	PaymentFailed   PaymentStatus = "FAILED"
	PaymentRefunded PaymentStatus = "REFUNDED"
)

type Order struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	OrderNumber string `json:"orderNumber" gorm:"type:varchar(50);uniqueIndex;not null"`

	UserID uint `json:"userId" gorm:"not null;index"`

	Items []orderitems.OrderItem `json:"items" gorm:"foreignKey:OrderID"`

	SubTotal      int `json:"subTotal" gorm:"not null"`
	DiscountTotal int `json:"discountTotal" gorm:"default:0"`
	ShippingFee   int `json:"shippingFee" gorm:"default:0"`
	GrandTotal    int `json:"grandTotal" gorm:"not null"`

	PaymentMethod enums.PaymentMethod `json:"paymentMethod" gorm:"type:varchar(20);not null"`
	Status        OrderStatus         `json:"status" gorm:"type:varchar(30);default:'PENDING_PAYMENT'"`
	PaymentStatus PaymentStatus       `json:"paymentStatus" gorm:"type:varchar(30);default:'PENDING'"`

	PaymentID string `json:"paymentId" gorm:"type:varchar(100)"`

	ShippingFullName     string `json:"shippingFullName" gorm:"type:varchar(100);not null"`
	ShippingPhone        string `json:"shippingPhone" gorm:"type:varchar(20);not null"`
	ShippingAddressLine1 string `json:"shippingAddressLine1" gorm:"type:text;not null"`
	ShippingAddressLine2 string `json:"shippingAddressLine2" gorm:"type:text"`
	ShippingCity         string `json:"shippingCity" gorm:"type:varchar(100);not null"`
	ShippingState        string `json:"shippingState" gorm:"type:varchar(100);not null"`
	ShippingCountry      string `json:"shippingCountry" gorm:"type:varchar(100);not null"`
	ShippingPincode      string `json:"shippingPincode" gorm:"type:varchar(20);not null"`

	IdempotencyKey string `json:"idempotencyKey" gorm:"type:varchar(150);uniqueIndex;not null"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
