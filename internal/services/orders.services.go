package services

import (
	"context"
	"fmt"
	"time"

	"example.com/m/v4/internal/dto"
	"example.com/m/v4/internal/enums"
	orderitems "example.com/m/v4/internal/models/orderItems"
	ordersModel "example.com/m/v4/internal/models/orders"
	"example.com/m/v4/internal/repository"
	"github.com/google/uuid"
)

type OrderService struct {
	orderRepo     *repository.OrderRepository
	catalogClient CatalogClient
}

func NewOrderService(orderRepo *repository.OrderRepository, catalogClient CatalogClient) *OrderService {
	return &OrderService{orderRepo: orderRepo, catalogClient: catalogClient}
}

func (s *OrderService) CreateOrder(
	ctx context.Context,
	userID uint,
	req dto.CreateOrderDTO,
	idempotencyKey string,
) (*dto.CreateOrderResponse, error) {
	if idempotencyKey == "" {
		return nil, fmt.Errorf("idempotency key is required")
	}

	var orderItems []orderitems.OrderItem

	subTotal := 0
	discountTotal := 0

	for _, item := range req.Items {
		product, err := s.catalogClient.GetProductByID(ctx, item.ProductID)
		if err != nil {
			return nil, err
		}

		if !product.Status {
			return nil, fmt.Errorf("product %d is not active", product.ID)
		}

		if product.StockQuantity < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s", product.Name)
		}

		finalUnitPrice := product.Price
		if product.DiscountPrice > 0 && product.DiscountPrice < product.Price {
			finalUnitPrice = product.DiscountPrice
		}

		itemTotal := finalUnitPrice * item.Quantity

		subTotal += product.Price * item.Quantity
		discountTotal += (product.Price - finalUnitPrice) * item.Quantity

		orderItems = append(orderItems, orderitems.OrderItem{
			ProductID:     product.ID,
			ProductName:   product.Name,
			ProductSKU:    product.SKU,
			ProductImage:  product.Image,
			UnitPrice:     product.Price,
			DiscountPrice: product.DiscountPrice,
			Quantity:      item.Quantity,
			Total:         itemTotal,
		})
	}

	shippingFee := 0
	grandTotal := subTotal - discountTotal + shippingFee

	orderStatus := ordersModel.OrderPendingPayment
	paymentStatus := ordersModel.PaymentPending

	if req.PaymentMethod == enums.PaymentCOD {
		orderStatus = ordersModel.OrderConfirmed
	}

	order := &ordersModel.Order{
		OrderNumber:    generateOrderNumber(),
		UserID:         userID,
		Items:          orderItems,
		SubTotal:       subTotal,
		DiscountTotal:  discountTotal,
		ShippingFee:    shippingFee,
		GrandTotal:     grandTotal,
		PaymentMethod:  req.PaymentMethod,
		Status:         orderStatus,
		PaymentStatus:  paymentStatus,
		IdempotencyKey: idempotencyKey,

		ShippingFullName:     req.ShippingAddress.FullName,
		ShippingPhone:        req.ShippingAddress.Phone,
		ShippingAddressLine1: req.ShippingAddress.AddressLine1,
		ShippingAddressLine2: req.ShippingAddress.AddressLine2,
		ShippingCity:         req.ShippingAddress.City,
		ShippingState:        req.ShippingAddress.State,
		ShippingCountry:      req.ShippingAddress.Country,
		ShippingPincode:      req.ShippingAddress.Pincode,
	}

	createdOrder, err := s.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateOrderResponse{
		Order: createdOrder,
	}

	// Later: if ONLINE, call payment service and set PaymentURL.
	// For now, leave it empty until Stripe/payment service is ready.

	return response, nil
}

func generateOrderNumber() string {
	return fmt.Sprintf("ORD-%d-%s", time.Now().Unix(), uuid.NewString()[:8])
}
