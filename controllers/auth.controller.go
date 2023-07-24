package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/panbhatt/go-gin-crud-gorm/models"
	"github.com/panbhatt/go-gin-crud-gorm/utils"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{
		DB: DB,
	}
}

func (ac *AuthController) SignUpUser(ctx *gin.Context) {

	var payload *models.SignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Password is not equal"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	now := time.Now()
	newUser := models.User{
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		Role:      "user",
		Verified:  true,
		Photo:     payload.Photo,
		Provider:  "local",
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User already exists "})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "SOmething bad happened"})
		return
	}

	userResponse := &models.UserResponse{

		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		Photo:     newUser.Photo,
		Role:      newUser.Role,
		Provider:  newUser.Provider,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
	return
}
