package handler

import (
	"go/src/broadcaster"
	"go/src/payment"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentResponse struct {
	Success bool
	Message string
	Data    interface{}
}

type paymentHandler struct {
	paymentService payment.Service
	broadcaster    broadcaster.Broadcaster
}

func NewPaymentHandler(paymentService payment.Service, broadcaster broadcaster.Broadcaster) *paymentHandler {
	return &paymentHandler{
		paymentService,
		broadcaster,
	}
}

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
	ph.broadcaster.Submit(payment)
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

func (ph *paymentHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, PaymentResponse{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}
	payment, err := ph.paymentService.GetById(id)
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
		Data:    payment,
	})
}

// update
func (ph *paymentHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, PaymentResponse{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	var input payment.InputPayment
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, PaymentResponse{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		})
		return
	}

	uPayment, err := ph.paymentService.Update(id, input)
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
		Message: "Payment updated",
		Data:    uPayment,
	})
}

func (ph *paymentHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, PaymentResponse{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}
	err = ph.paymentService.Delete(id)
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
		Message: "Payment deleted",
	})
}

// Conect the brodcaster to the /sse endpoint
func (ph *paymentHandler) Sse(c *gin.Context) {
	channel := make(chan interface{})
	ph.broadcaster.Register(channel)
	defer ph.broadcaster.Unregister(channel)
	c.Stream(func(w io.Writer) bool {
		for {
			select {
			case payment := <-channel:
				c.SSEvent("payment", payment)
			}
		}
	})
}
