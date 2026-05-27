package handlers

import (
	"net/http"

	"example.com/m/v4/internal/dto"
	"example.com/m/v4/internal/services"
	"github.com/gin-gonic/gin"
)

type OrderHandlers struct {
	orderService *services.OrderService
}

func NewOrderHandlers(orderService *services.OrderService) *OrderHandlers {
	return &OrderHandlers{orderService: orderService}
}

func (h *OrderHandlers) CreateOrder(c *gin.Context) {
	idempotencyKey := c.GetHeader("Idempotency-Key")
	if idempotencyKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Idempotency-Key header is required",
		})
		return
	}

	var req dto.CreateOrderDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userID := uint(1)

	result, err := h.orderService.CreateOrder(
		c.Request.Context(),
		userID,
		req,
		idempotencyKey,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "order created successfully",
		"data":    result,
	})
}
