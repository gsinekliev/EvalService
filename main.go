package main

import (
	"fmt"
	"github.com/gsinekliev/eval-service/service"
)

func main() {
	r := service.SetupRouter()
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		fmt.Printf("Error while running the service: %v", err)
	}
}
