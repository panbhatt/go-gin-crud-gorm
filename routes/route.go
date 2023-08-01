package routes

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/panbhatt/go-gin-crud-gorm/controllers"
	"github.com/panbhatt/go-gin-crud-gorm/initializers"
	"github.com/panbhatt/go-gin-crud-gorm/middleware"
)

func AddRoutes(rg *gin.RouterGroup, config initializers.Config) {

	router := rg.Group("/auth")

	authController := ctrl.NewAuthController(initializers.DB)

	router.POST("/register", authController.SignUpUser)
	router.POST("/login", authController.SignInUser)
	router.GET("/refresh", authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(), authController.Logout)

	userRouter := rg.Group("/users")
	userController := ctrl.NewUserController(initializers.DB)
	userRouter.GET("/me", middleware.DeserializeUser(), userController.GetUserData)

}
