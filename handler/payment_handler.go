package handler

import (
	"api_goshop/dto"
	"api_goshop/handleError"
	"api_goshop/helper"
	"api_goshop/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type payment_handler struct {
	service services.PaymentService
}

func NewPaymentHandler(service services.PaymentService) *payment_handler {
	return &payment_handler{service}
}

func (h payment_handler) CreateTransaction(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
		return
	}
	// Mengonversi userID menjadi tipe yang sesuai
	id, ok := userID.(*int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	var itemDetails []dto.ItemDetailsRequest
	if err := c.ShouldBindJSON(&itemDetails); err != nil {
		handleError.HandleError(c, &handleError.BadRequestError{Message: err.Error()})
		return
	}
	snapRes, err := h.service.CreatePayment(itemDetails, id)
	if err != nil {
		handleError.HandleError(c, &handleError.BadRequestError{Message: err.Error()})
		return
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Succesfully to create payment transaction",
		Data:       snapRes,
	})
	c.JSON(http.StatusOK, res)
}
