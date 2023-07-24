package main

import (
	"fmt"
	"log"

	"github.com/panbhatt/go-gin-crud-gorm/initializers"
	"github.com/panbhatt/go-gin-crud-gorm/models"
)

func init() {

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Cloud not load environment configuration", err)
	}

	initializers.ConnectDB(&config)

}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration Complete")
}
