package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	setupServer().Run(":2001")
	fmt.Println("running at localhost:2001...")

}

func setupServer() *gin.Engine {

	router := *gin.Default()

	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true
	cfg.AllowMethods = []string{"GET", "POST"}
	cfg.AllowHeaders = []string{"Authorization", "Origin", "Accept", "X-Requested-With", " Content-Type", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	router.Use(cors.New(cfg))

	router.POST("/api/calculator", Calculator)

	return router
}
