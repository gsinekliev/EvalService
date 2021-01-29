package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/evaluate", func(c *gin.Context) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)
		c.String(http.StatusOK, buf.String())
	})
	return r
}

func main() {

	r := SetupRouter()

	r.Run("localhost:8080")
}
