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
		result, status := computeExpression(request.Expression)
		if status == NoError {
			c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("%.2f", result)})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"result": fmt.Sprintf("Bad Request:%s", status)})
		}
	})

	r.POST("/validate", func(c *gin.Context) {
		request := Request{}
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)

		json.Unmarshal(buf.Bytes(), &request)

		fmt.Printf("[Evaluate] Body: %s", request.Expression)
		status := validateExpression(request.Expression)
		if status == NoError {
			c.JSON(http.StatusOK, gin.H{"valid": "true"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"valid": "false", "reason": status})
		}
	})
	return r
}

func main() {

	r := SetupRouter()

	r.Run("localhost:8080")
}
