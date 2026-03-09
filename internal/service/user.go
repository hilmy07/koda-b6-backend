package service

import (
	"backend/internal/lib"
	"backend/internal/repository"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UploadUserPicture(ctx *gin.Context) (string, bool) {
	
	file, _ := ctx.FormFile("picture")

	if file != nil {
		if x := lib.AcceptedMime(file.Header.Get("Content-Type"), "picture"); x {
			ctx.SaveUploadedFile(file, "./uploads/"+file.Filename)
			return file.Filename, true
		}
	}

	return "", false
}



