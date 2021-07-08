package main

import (
	"github.com/gin-gonic/gin"
	"github.com/LucasCarioca/go-service/pkg/routes"
)

func main() {
	app := gin.Default()
	routes.HealthRouter(app)
	app.Run()
}
