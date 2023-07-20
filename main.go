package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/panbhatt/go-gin-crud-gorm/initializers"
)

var (
	server *gin.Engine
)

func main() {

	config, err := initializers.LoadConfig(".")

	if err != nil {
		fmt.Printf("An Error Ocured : %v", err)
	}

	initializers.ConnectDB(&config)

	fmt.Println("======DOING DB Migrations if Any = ===========")
	initializers.DB.AutoMigrate()

	fmt.Println(" DB host = " + config.DBHost)

	fmt.Println("===============Starting GIN things now================")
	server = gin.Default()
	apiRouter := server.Group("/api")

	apiRouter.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "Health is ok "})

	})

	log.Fatal(server.Run(":" + config.ServerPort))

}
