package main

import (
	"fmt"

	"github.com/OverCV/go-automats/controllers"
)

func main() {
	controller := controllers.NewController()
	err := controller.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
