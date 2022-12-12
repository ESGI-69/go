package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID        int
	Name      string
	Price     float64
	timestamp time.Time
}

func CreateProduct(context *gin.Context) {
	if context.Query("name") == "" || context.Query("price") == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	// Product := Product{
	// 	Name:      context.Query("name"),
	// 	Price:     context.Query("price"),
	// 	timestamp: time.Now(),
	// }

	// context.JSON(http.StatusCreated, gin.H{
	// 	"message": "Product created !",
	// 	"data":    "dsqd",
	// })
}
