package main

import (
	"github.com/gsinekliev/eval-service/service"
)

func main() {
	r := service.SetupRouter()
	r.Run("localhost:8080")
}
