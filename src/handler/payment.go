package handler

import (
	"go/src/payment"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentResponse struct {
	Success bool
	Message string
	Data    interface{}
}

type paymentHandler struct {
	paymentService payment.Service
}

func NewPaymentHandler(paymentService payment.Service) *paymentHandler {
	return &paymentHandler{paymentService}
}

//TODO: Implement the handler methods here

func (ph *paymentHandler) TestPayment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}

// func Create payment in bdd using PaymentResponse
func (ph *paymentHandler) Create(c *gin.Context) {
	var input payment.InputPayment
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, PaymentResponse{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		})
		return
	}

	// doesnt work because of the foreign key
	payment, err := ph.paymentService.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, PaymentResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, PaymentResponse{
		Success: true,
		Message: "Payment created",
		Data:    payment,
	})
}

func (ph *paymentHandler) GetAll(c *gin.Context) {
	payments, err := ph.paymentService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, PaymentResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, PaymentResponse{
		Success: true,
		Message: "All payments",
		Data:    payments,
	})
}
