package handlers

import (
	"backend/internal/models"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(ctx *gin.Context) {

	uid, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	req := models.Order{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	err := h.service.CreateOrder(req, uid)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"success": true,
		"message": "order created",
	})
}