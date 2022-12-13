package handler

// import local package payment in go/src/payment
import (
	"payment"
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
