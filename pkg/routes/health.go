package routes

import (
	"net/http"

	"github.com/LucasCarioca/go-service/pkg/services"
	"github.com/gin-gonic/gin"
)

func get(context *gin.Context) {
	ps := services.PinetService{}
	pinetStatus := ps.Status()
	context.JSON(http.StatusOK, gin.H{
		"message": "warning",
		"services": gin.H{
			"something": "up",
			"pinet":     pinetStatus,
		},
	})
}

func HealthRouter(router *gin.Engine) {
	router.GET("/api/v1/health", get)
}
