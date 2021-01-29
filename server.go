package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result"`
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/evaluate", func(c *gin.Context) {

		request := Request{}
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)

		json.Unmarshal(buf.Bytes(), &request)

		fmt.Printf("[Evaluate] Body: %s", request.Expression)

		c.JSON(http.StatusOK, gin.H{"result": request.Expression})
	})
	return r
}

func main() {

	r := SetupRouter()

	r.Run("localhost:8080")
}
