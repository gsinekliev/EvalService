package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gsinekliev/eval-service/service/eval"
	"github.com/gsinekliev/eval-service/service/models"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	errorStore := models.InitErrorStore()
	evaluator := eval.Evaluator{}
	r.POST("/evaluate", func(c *gin.Context) {
		request := models.Request{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": fmt.Sprintf("Invalid JSON: %v", err)})
			return
		}

		result, status := evaluator.ComputeExpression(request.Expression)
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
		request := models.Request{}
		evaluator := eval.Evaluator{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": fmt.Sprintf("Invalid JSON: %v", err)})
			return
		}

		status := evaluator.ValidateExpression(request.Expression)
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
