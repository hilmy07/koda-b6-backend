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

func getUserID(ctx *gin.Context) (int, bool) {
	userID, exists := ctx.Get("user_id")
	if !exists || userID == nil {
		return 0, false
	}

	uidFloat, ok := userID.(float64)
	if !ok {
		return 0, false
	}

	return int(uidFloat), true
}

func (h *CartHandler) GetCartByUser(ctx *gin.Context) {

	uid, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	data, err := h.service.GetCartByUserId(uid)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	if len(data) == 0 {
		ctx.JSON(404, gin.H{
			"data":    []models.Cart{},
			"success": false,
			"message": "cart is empty",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data":    data,
		"success": true,
	})
}
