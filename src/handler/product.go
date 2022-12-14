package handler

// import local package product in go/src/product
import (
	"go/src/product"
	"net/http"
	"strconv"

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

func (productHandler *productHandler) GetAll(c *gin.Context) {
	products, err := productHandler.productService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, ProductResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ProductResponse{
		Success: true,
		Message: "All products",
		Data:    products,
	})
}

func (productHandler *productHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	product, err := productHandler.productService.GetById(idInt)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, ProductResponse{
				Success: false,
				Message: "Product not found",
				Data:    err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, ProductResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ProductResponse{
		Success: true,
		Message: "Product by ID",
		Data:    product,
	})
}

func (productHandler *productHandler) Create(c *gin.Context) {
	var input product.InputProduct
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, ProductResponse{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		})
		return
	}
	product, err := productHandler.productService.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, ProductResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ProductResponse{
		Success: true,
		Message: "Product created",
		Data:    product,
	})
}

func (productHandler *productHandler) Update(c *gin.Context) {
	var input product.InputProduct
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, ProductResponse{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		})
		return
	}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	product, err := productHandler.productService.Update(idInt, input)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, ProductResponse{
				Success: false,
				Message: "Product not found",
				Data:    err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, ProductResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ProductResponse{
		Success: true,
		Message: "Product updated",
		Data:    product,
	})
}

func (productHandler *productHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	err = productHandler.productService.Delete(idInt)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, ProductResponse{
				Success: false,
				Message: "Product not found",
				Data:    err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, ProductResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ProductResponse{
		Success: true,
		Message: "Product deleted",
	})
}
