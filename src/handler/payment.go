package handler

// import local package payment in go/src/payment
import (
	"go/src/payment"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
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

func (ps *paymentHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
