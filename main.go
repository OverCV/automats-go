package main

import (
	"fmt"

	"github.com/OverCV/go-automats/controllers"
	"github.com/OverCV/go-automats/views"
)

func main() {
	// Return statement without specify variable name
	controller := controllers.NewController()
	uinterface := views.NewUI(controller)
	err := uinterface.RunUI()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
