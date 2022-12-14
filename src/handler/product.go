package handler

// import local package product in go/src/product
import (
	"go/src/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductResponse struct {
	Success bool
	Message string
	Data    interface{}
}

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

//TODO: Implement the handler methods here

func (ps *productHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
