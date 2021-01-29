package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gsinekliev/eval-service/service/eval"
	"github.com/gsinekliev/eval-service/service/models"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	errorStore := models.InitErrorStore()
	r.POST("/evaluate", func(c *gin.Context) {

		request := Request{}
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)

		json.Unmarshal(buf.Bytes(), &request)

		fmt.Printf("[Evaluate] Body: %s", request.Expression)
		result, status := eval.ComputeExpression(request.Expression)
		if status == eval.NoError {
			c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("%.2f", result)})
		} else {
			modelError := models.Error{
				Expression: request.Expression,
				Endpoint:   c.Request.URL.Path,
				Frequency:  1,
				ErrorType:  status,
			}
			errorStore.AddError(modelError)
			c.JSON(http.StatusBadRequest, gin.H{"result": fmt.Sprintf("Bad Request:%s", status)})
		}
	})

	r.POST("/validate", func(c *gin.Context) {
		request := Request{}
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)

		json.Unmarshal(buf.Bytes(), &request)

		fmt.Printf("[Evaluate] Body: %s", request.Expression)
		status := eval.ValidateExpression(request.Expression)
		if status == eval.NoError {
			c.JSON(http.StatusOK, gin.H{"valid": "true"})
		} else {
			modelError := models.Error{
				Expression: request.Expression,
				Endpoint:   c.Request.URL.Path,
				Frequency:  1,
				ErrorType:  status,
			}
			errorStore.AddError(modelError)
			c.JSON(http.StatusBadRequest, gin.H{"valid": "false", "reason": status})
		}
	})

	r.GET("/errors", func(c *gin.Context) {

		errorsList := make([]models.Error, 0, len(errorStore))

		for _, value := range errorStore {
			errorsList = append(errorsList, value)
		}

		c.JSON(http.StatusOK, errorsList)
	})

	return r
}
