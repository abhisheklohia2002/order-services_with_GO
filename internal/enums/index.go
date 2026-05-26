package enums

type OrderStatus string

const (
	OrderPendingPayment OrderStatus = "PENDING_PAYMENT"
	OrderConfirmed      OrderStatus = "CONFIRMED"
	OrderProcessing     OrderStatus = "PROCESSING"
	OrderShipped        OrderStatus = "SHIPPED"
	OrderDelivered      OrderStatus = "DELIVERED"
	OrderCancelled      OrderStatus = "CANCELLED"
	OrderPaymentFailed  OrderStatus = "PAYMENT_FAILED"
)

type PaymentStatus string

const (
	PaymentPending  PaymentStatus = "PENDING"
	PaymentPaid     PaymentStatus = "PAID"
	PaymentFailed   PaymentStatus = "FAILED"
	PaymentRefunded PaymentStatus = "REFUNDED"
)

type PaymentMethod string

const (
	PaymentOnline PaymentMethod = "ONLINE"
	PaymentCOD    PaymentMethod = "COD"
)
