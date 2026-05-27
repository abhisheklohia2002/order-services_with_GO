package repository

import (
	ordersModel "example.com/m/v4/internal/models/orders"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *ordersModel.Order) (*ordersModel.Order, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) GetOrderById(orderID uint) (*ordersModel.Order, error) {
	var order ordersModel.Order

	err := r.db.
		Preload("Items").
		First(&order, orderID).Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}
