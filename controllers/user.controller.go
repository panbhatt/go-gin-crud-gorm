package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/panbhatt/go-gin-crud-gorm/models"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) UserController {
	return UserController{DB: db}
}

func (uc *UserController) GetUserData(ctx *gin.Context) {

	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.UserResponse{

		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Email:     currentUser.Email,
		Photo:     currentUser.Photo,
		Role:      currentUser.Role,
		Provider:  currentUser.Provider,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})

}
