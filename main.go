package main

import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/LucasCarioca/go-service/pkg/routes"
)

func main() {
	app := gin.Default()

	routes.HealthRouter(app)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	app.Run(addr)
}
