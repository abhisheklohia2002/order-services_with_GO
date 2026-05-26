package dto

import "example.com/m/v4/internal/enums"

type CreateOrderDTO struct {
	PaymentMethod   enums.PaymentMethod  `json:"paymentMethod" binding:"required,oneof=ONLINE COD"`
	Items           []CreateOrderItemDTO `json:"items" binding:"required,min=1"`
	ShippingAddress ShippingAddressDTO   `json:"shippingAddress" binding:"required"`
}

type CreateOrderItemDTO struct {
	ProductID uint `json:"productId" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type ShippingAddressDTO struct {
	FullName     string `json:"fullName" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	AddressLine1 string `json:"addressLine1" binding:"required"`
	AddressLine2 string `json:"addressLine2"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	Country      string `json:"country" binding:"required"`
	Pincode      string `json:"pincode" binding:"required"`
}
