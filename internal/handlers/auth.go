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

type ForgotPasswordHandler struct {
	service *service.ForgotPasswordService
}

func NewForgotPasswordHandler(service *service.ForgotPasswordService) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{
		service: service,
	}
}

func (h *AuthHandler) AuthLogin(ctx *gin.Context) {

	req := models.RequestLogin{}

	ctx.ShouldBindJSON(&req)

	user, token, err := h.service.Login(req.Email, req.Password)

	if err != nil {
		
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "login failed",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "Login success",
		"email": req.Email,
		"token": token,
		"id": user.Id,
		"fullName": user.Fullname,
		"profileImg": user.Profile_img,
	})
}

func (h *AuthHandler) AuthRegister(ctx *gin.Context) {

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

func (h *AuthHandler) AuthProfile(ctx *gin.Context) {

	req := models.CreateUserRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	err := h.service.ProfileRegister(req)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"success": true,
		"message": "update profile success",
	})
}

func (h *ForgotPasswordHandler) RequestForgotPassword(ctx *gin.Context) {

	var req models.User

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	err := h.service.RequestForgotPassword(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OTP sent successfully",
	})
}

func (h *ForgotPasswordHandler) ResetPassword(ctx *gin.Context) {

	var req struct {
		Email       string `json:"email"`
		Code        string `json:"code"`
		NewPassword string `json:"new_password"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	forgot := models.ForgotPassword{
		Email: req.Email,
		Code:  req.Code,
	}

	user := models.User{
		Email:    req.Email,
		Password: req.NewPassword,
	}

	err := h.service.ResetPassword(forgot, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "password reset success",
	})
}
