package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func get() {
}

func main() {
	engine := gin.Default()

	engine.GET("/ping", func(context *gin.Context) {
		fmt.Println("GET /")
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	engine.Run()

	fmt.Println("Launched")
}
