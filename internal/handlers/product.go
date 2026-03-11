package handlers

import (
	"backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProductList(ctx *gin.Context) {

	products, err := h.service.GetProductList()

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "failed get products",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"data": products,
	})
}

func (h *ProductHandler) GetProductDetail(ctx *gin.Context) {

	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid product id",
		})
		return
	}

	product, err := h.service.GetProductDetail(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "failed get product detail",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"data": product,
	})
}