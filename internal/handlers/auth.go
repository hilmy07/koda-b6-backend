package handlers

import (
	"net/http"

	"backend/internal/models"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) AuthLogin(ctx *gin.Context) {

	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := Request{}

	ctx.ShouldBindJSON(&req)

	token, err := h.service.Login(req.Email, req.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "login failed",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "Login success",
		"token": token,
	})
}

func (h *AuthHandler) AuthRegister(ctx *gin.Context) {

	// type Request struct {
	// 	Email    string `json:"email"`
	// 	Password string `json:"password"`
	// }

	req := models.CreateUserRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	err := h.service.Register(req)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"success": true,
		"message": "register success",
	})
}