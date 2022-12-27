package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"go/src/handler"
	"go/src/payment"
	"go/src/product"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 404 custom
func notFound(context *gin.Context) {
	context.JSON(404, gin.H{
		"message": "❌ Page not found ❌",
	})
}

func main() {
	// Connection DB
	dbURL := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Println("❌ Connection to DB failed")
		fmt.Println(dbURL)
		log.Fatal(err.Error())
	} else {
		fmt.Println("🔗 Connection to DB OK")
	}

	db.AutoMigrate(&payment.Payment{}, &product.Product{})

	router := gin.Default()

	// Load the pages templates https://gin-gonic.com/docs/examples/html-rendering/
	router.LoadHTMLGlob("src/templates/*")

	// Add the static files
	router.Static("static/", "./src/js")

	// Create the api
	api := router.Group("/api")
	web := router.Group("/")

	router.NoRoute(notFound)

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	api.GET("/products", productHandler.GetAll)
	api.GET("/products/:id", productHandler.GetByID)
	api.POST("/products", productHandler.Create)
	api.PATCH("/products/:id", productHandler.Update)
	api.DELETE("/products/:id", productHandler.Delete)

	paymentRepository := payment.NewRepository(db)
	paymentService := payment.NewService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService)
	api.POST("/payments", paymentHandler.Create)
	api.GET("/payments", paymentHandler.GetAll)
	api.GET("/payments/:id", paymentHandler.GetById)
	api.PATCH("/payments/:id", paymentHandler.Update)
	api.DELETE("/payments/:id", paymentHandler.Delete)

	web.GET("/", func(c *gin.Context) {
		products, _ := productService.GetAll()
		payments, _ := paymentService.GetAll()
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"products": products,
			"payments": payments,
		})
	})

	web.GET("/products/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "productCreation.tmpl", gin.H{})
	})

	web.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)
		product, _ := productService.GetById(idInt)
		c.HTML(http.StatusOK, "product.tmpl", gin.H{
			"product": product,
		})
	})

	web.GET("/payments/create", func(c *gin.Context) {
		products, _ := productService.GetAll()
		c.HTML(http.StatusOK, "paymentCreation.tmpl", gin.H{
			"products": products,
		})
	})

	web.GET("/payments/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)
		payment, _ := paymentService.GetById(idInt)
		c.HTML(http.StatusOK, "payment.tmpl", gin.H{
			"payment": payment,
		})
	})

	router.Run()
}
