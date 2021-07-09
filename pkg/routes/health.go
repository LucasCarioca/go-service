package routes

import (
	"net/http"
	"github.com/LucasCarioca/go-service/pkg/services"
	"github.com/gin-gonic/gin"
)

func get(context *gin.Context) {
	message := "up"
	hs := services.HealthService{}

	pinetHub := hs.CheckStatus("http://localhost:8000/api/v1/health")
	pinetNode := hs.CheckStatus("http://localhost:8000/api/v1/health")
	list := [2]services.Status{pinetHub, pinetNode}

	for _, s := range  list{
		if s.Message != "up" {
			message = "outage"
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"message": message,
		"services": gin.H{
			"pinet-hub": pinetHub.Message,
			"pinet-node-01": pinetNode.Message,
		},
	})
}

func HealthRouter(router *gin.Engine) {
	router.GET("/api/v1/health", get)
}
