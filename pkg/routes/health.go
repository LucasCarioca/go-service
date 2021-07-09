package routes

import (
	"net/http"
	"github.com/LucasCarioca/go-service/pkg/services"
	"github.com/gin-gonic/gin"
)

func get(context *gin.Context) {
	message := "up"
	hs := services.HealthService{}

	pinetHub := hs.CheckStatus("192.168.1.150:30101/api/v1/health")
	pinetNode01 := hs.CheckStatus("http://192.168.1.200:8000/api/v1/health")
	pinetNode02 := hs.CheckStatus("http://192.168.1.201:8000/api/v1/health")
	pinetNode03 := hs.CheckStatus("http://192.168.1.203:8000/api/v1/health")
	pinetNode04 := hs.CheckStatus("http://192.168.1.202:8000/api/v1/health")
	list := [5]services.Status{
		pinetHub, 
		pinetNode01,
		pinetNode02,
		pinetNode03,
		pinetNode04,
	}

	for _, s := range  list{
		if s.Message != "up" {
			message = "outage"
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"message": message,
		"services": gin.H{
			"pinet-hub": pinetHub.Message,
			"pinet-node-01": pinetNode01.Message,
			"pinet-node-02": pinetNode02.Message,
			"pinet-node-03": pinetNode03.Message,
			"pinet-node-04": pinetNode04.Message,
		},
	})
}

func HealthRouter(router *gin.Engine) {
	router.GET("/api/v1/health", get)
}
