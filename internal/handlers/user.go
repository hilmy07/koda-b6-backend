package handlers

import (
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) UploadPicture (ctx *gin.Context) {
	
	filename, uploadStatus := h.service.UploadUserPicture(ctx)

	if uploadStatus {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Upload success",
			"results": filename,
		})
	}else{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Image doesnt accepted",
		})
	}
}