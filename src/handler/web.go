package handler

import (
	"go/src/broadcaster"
	"go/src/payment"
	"go/src/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type webHandler struct {
	productService product.Service
	paymentService payment.Service
	broadcaster    broadcaster.Broadcaster
}

func NewWebHandler(productService product.Service, paymentService payment.Service, broadcaster broadcaster.Broadcaster) *webHandler {
	return &webHandler{
		productService: productService,
		paymentService: paymentService,
		broadcaster:    broadcaster,
	}
}

func (webHandler *webHandler) Home(c *gin.Context) {
	products, _ := webHandler.productService.GetAll()
	payments, _ := webHandler.paymentService.GetAll()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"products": products,
		"payments": payments,
	})
}

func (webHandler *webHandler) CreateProduct(c *gin.Context) {
	c.HTML(http.StatusOK, "productCreation.tmpl", gin.H{})
}

func (webHandler *webHandler) CreatePayment(c *gin.Context) {
	products, _ := webHandler.productService.GetAll()
	c.HTML(http.StatusOK, "paymentCreation.tmpl", gin.H{
		"products": products,
	})
}

func (webHandler *webHandler) EditProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	product, _ := webHandler.productService.GetById(idInt)
	c.HTML(http.StatusOK, "productEdit.tmpl", gin.H{
		"product": product,
	})
}

func (webHandler *webHandler) EditPayment(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	payment, _ := webHandler.paymentService.GetById(idInt)
	products, _ := webHandler.productService.GetAll()
	c.HTML(http.StatusOK, "paymentEdit.tmpl", gin.H{
		"payment":  payment,
		"products": products,
	})
}
