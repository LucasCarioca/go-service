package main

import (
	"github.com/gin-gonic/gin"
	"github.com/LucasCarioca/status-service/pkg/routes"
)

func main() {
	app := gin.Default()
	routes.HealthRouter(app)
	app.Run()
}
