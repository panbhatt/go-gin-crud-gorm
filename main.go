package main

import (
	"fmt"

	"github.com/panbhatt/go-gin-crud-gorm/initializers"
)

func main() {

	config, err := initializers.LoadConfig(".")

	if err != nil {
		fmt.Println("An Error Ocured : %v", err)
	}

	initializers.ConnectDB(&config)
	fmt.Println(" DB host = " + config.DBHost)

}
