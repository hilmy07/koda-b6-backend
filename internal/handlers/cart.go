package handlers

import (
	"backend/internal/models"
	"backend/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service *service.CartService
}

func NewCartHandler(service *service.CartService) *CartHandler {
	return &CartHandler{service: service}
}


func (h *CartHandler) CreateCart(ctx *gin.Context) {

	req := models.Cart{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	err := h.service.CreateCart(req)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"success": true,
		"message": "cart item created",
	})
}

func (h *CartHandler) GetCartList(ctx *gin.Context) {

	carts, err := h.service.GetCartList()

	if err != nil {
		log.Println("ERROR GetCartList:", err)

		ctx.JSON(500, gin.H{
			"message": "failed get carts",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"data": carts,
	})
}


