package handlers

import (
	"backend/internal/models"
	"backend/internal/service"
	"fmt"
	"log"
	"net/http"
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

	pageQuery := ctx.Query("page")
	page := 1
	if pageQuery != "" {
		if p, err := strconv.Atoi(pageQuery); err == nil && p > 0 {
			page = p
		}
	}

	paginatedProducts, err := h.service.GetProductList(page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, paginatedProducts)
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {

	req := models.Product{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	err := h.service.CreateProduct(req)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"success": true,
		"message": "product created",
	})
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := h.service.DeleteProduct(id)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "product deleted",
	})
}

func (h *ProductHandler) GetRecommendedProduct(ctx *gin.Context) {

	products, err := h.service.GetRecommendedProduct()

	if err != nil {
		log.Println("ERROR GetRecommendedProduct:", err)

		ctx.JSON(500, gin.H{
			"message": "failed get products",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"data": products,
	})
}

func (h *ProductHandler) GetAllProduct(ctx *gin.Context) {

	products, err := h.service.GetAllProduct()

	if err != nil {
		log.Println("ERROR GetAllProduct:", err)

		ctx.JSON(500, gin.H{
			"message": "failed get all products",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"data": products,
	})
}

func (h *ProductHandler) GetProductReview(ctx *gin.Context) {

	reviews, err := h.service.GetProductReview()

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "failed get products",
		})
		return
	}

	ctx.IndentedJSON(200, gin.H{
		"success": true,
		"data": reviews,
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
		fmt.Println("ERROR:", err)
		ctx.JSON(500, gin.H{
			"message": "failed get product detail",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"data": []product,
	})
}