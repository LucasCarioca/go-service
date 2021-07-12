package main

import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/LucasCarioca/go-service/pkg/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
)

func main() {
	app := gin.Default()
	app.Use(cors.Default())

	app.Use(static.Serve("/", static.LocalFile("./public", false)))

	routes.HealthRouter(app)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	addr := fmt.Sprintf(":%s", port)
	app.Run(addr)
}
