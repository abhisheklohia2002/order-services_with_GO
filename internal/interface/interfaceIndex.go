package interfaceIndex

import ordersModel "example.com/m/v4/internal/models/orders"

type OrderRepositoryInterface interface {
	CreateOrder(order *ordersModel.Order) (*ordersModel.Order, error)
	GetOrderById(orderID uint) (*ordersModel.Order, error)
}
