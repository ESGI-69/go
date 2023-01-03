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

// Create creates a new payment
// @Summary      Create a new payment
// @Description  Creates a new payment and returns the created payment
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        payment   body      payment.InputPayment  true  "Payment information"
// @Success      200  {object}  payment.Payment
// @Failure      400  {object}  PaymentResponse
// @Failure      404  {object}  PaymentResponse
// @Failure      500  {object}  PaymentResponse
// @Router       /payments [post]
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

// GetAll get all payments
// @Summary      Get all payments
// @Description  Get all payments
// @Tags         payments
// @Accept       json
// @Produce      json
// @Success      200  {object}  payment.Payment
// @Failure      400  {object}  PaymentResponse
// @Failure      404  {object}  PaymentResponse
// @Failure      500  {object}  PaymentResponse
// @Router       /payments [get]
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

// GetById get payment by id
// @Summary      Get payment by id
// @Description  Get payment by id
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Payment id"
// @Success      200  {object}  payment.Payment
// @Failure      400  {object}  PaymentResponse
// @Failure      404  {object}  PaymentResponse
// @Failure      500  {object}  PaymentResponse
// @Router       /payments/{id} [get]
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

// Update updates a payment
// @Summary      Update a payment
// @Description  Updates a payment and returns the updated payment
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Payment id"
// @Param        payment   body      payment.InputPayment  true  "Payment information"
// @Success      200  {object}  payment.Payment
// @Failure      400  {object}  PaymentResponse
// @Failure      404  {object}  PaymentResponse
// @Failure      500  {object}  PaymentResponse
// @Router       /payments/{id} [patch]
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

// Delete deletes a payment
// @Summary      Delete a payment
// @Description  Deletes a payment
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Payment id"
// @Success      200  {object}  PaymentResponse
// @Failure      400  {object}  PaymentResponse
// @Failure      404  {object}  PaymentResponse
// @Failure      500  {object}  PaymentResponse
// @Router       /payments/{id} [delete]
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

func (ph *paymentHandler) Sse(c *gin.Context) {
	channel := make(chan interface{})
	ph.broadcaster.Register(channel)
	defer ph.broadcaster.Unregister(channel)
	c.Stream(func(w io.Writer) bool {
		for {
			select {
			case payment := <-channel:
				c.SSEvent("message", payment)
				return true
			}
		}
	})
}
