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

// Create creates a new product
// @Summary      Create a new product
// @Description  Creates a new product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product        body    product.InputProduct  true  "Product information"
// @Success      200  {object}  product.Product
// @Failure      400  {object}  ProductResponse
// @Failure      404  {object}  ProductResponse
// @Failure      500  {object}  ProductResponse
// @Router       /products [post]
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

// GetAll returns all products
// @Summary      Get all products
// @Description  Returns all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {object}  product.Product
// @Failure      400  {object}  ProductResponse
// @Failure      404  {object}  ProductResponse
// @Failure      500  {object}  ProductResponse
// @Router       /products [get]
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

// GetByID returns a product by ID
// @Summary      Get a product by ID
// @Description  Returns a product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  product.Product
// @Failure      400  {object}  ProductResponse
// @Failure      404  {object}  ProductResponse
// @Failure      500  {object}  ProductResponse
// @Router       /products/{id} [get]
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

// Update updates a product
// @Summary      Update a product
// @Description  Updates a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id          path    int     true  "Product ID"
// @Success      200  {object}  product.Product
// @Failure      400  {object}  ProductResponse
// @Failure      404  {object}  ProductResponse
// @Failure      500  {object}  ProductResponse
// @Router       /products/{id} [patch]
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

// Delete deletes a product
// @Summary      Delete a product
// @Description  Deletes a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id          path    int     true  "Product ID"
// @Param        product 	 body    product.InputProduct  true  "Product"
// @Success      200  {object}  product.Product
// @Failure      400  {object}  ProductResponse
// @Failure      404  {object}  ProductResponse
// @Failure      500  {object}  ProductResponse
// @Router       /products/{id} [delete]
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
