package main

import (
	"log"

	"github.com/duglas43/calculator/internal/executor"
	"github.com/duglas43/calculator/internal/handler"
	"github.com/gin-gonic/gin"

	_ "github.com/duglas43/calculator/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	exec := executor.NewExecutor()
	router := gin.Default()
	h := handler.NewHTTPHandler(exec)
	router.POST("/execute", h.HandleExecute)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("HTTP  listen on port 3000")
	if err := router.Run(":3000"); err != nil {
		log.Fatalf("Error HTTP: %v", err)
	}
}
