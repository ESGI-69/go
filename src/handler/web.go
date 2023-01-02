package handler

import (
	"go/src/payment"
	"go/src/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type webHandler struct{}

func NewWebHandler() *webHandler {
	return &webHandler{}
}

func (webHandler *webHandler) Home(products []product.Product, payments []payment.Payment) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"products": products,
			"payments": payments,
		})
	}
}

func (webHandler *webHandler) CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "productCreation.tmpl", gin.H{})
	}
}

func (webHandler *webHandler) CreatePayment(products []product.Product) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "paymentCreation.tmpl", gin.H{
			"products": products,
		})
	}
}

func (webHandler *webHandler) EditProduct(productService product.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)
		product, _ := productService.GetById(idInt)
		c.HTML(http.StatusOK, "productEdit.tmpl", gin.H{
			"product": product,
		})
	}
}

func (webHandler *webHandler) EditPayment(paymentService payment.Service, productService product.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)
		payment, _ := paymentService.GetById(idInt)
		products, _ := productService.GetAll()
		c.HTML(http.StatusOK, "paymentEdit.tmpl", gin.H{
			"payment":  payment,
			"products": products,
		})
	}
}
