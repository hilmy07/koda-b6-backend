package handlers

import (
	"backend/internal/service"
	"net/http"
	"strconv"

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

func (h *UserHandler) GetUser(ctx *gin.Context) {

	users, err := h.service.GetUsers()

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "failed get users",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"data": users,
	})
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := h.service.DeleteUser(id)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "user deleted",
	})
}
