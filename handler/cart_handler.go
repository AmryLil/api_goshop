package handler

import (
	"api_goshop/dto"
	"api_goshop/handleError"
	"api_goshop/helper"
	"api_goshop/models"
	"api_goshop/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type cart_handler struct {
	service services.CartService
}

func NewCartHandler(service services.CartService) *cart_handler {
	return &cart_handler{service}
}

func (h cart_handler) AddtoCartHandler(c *gin.Context) {
	var dataProduct dto.CartRequest

	if err := c.ShouldBindJSON(&dataProduct); err != nil {
		handleError.HandleError(c, &handleError.BadRequestError{Message: err.Error()})
		return
	}
	if err := h.service.AddtoCart(dataProduct); err != nil {
		handleError.HandleError(c, err)
		return
	}
	response := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Product added to cart",
	})
	c.JSON(http.StatusCreated, response)

}

func (h *cart_handler) UpdateCartHandler(c *gin.Context) {
	var dataProduct dto.CartRequest

	if err := c.ShouldBindJSON(&dataProduct); err != nil {
		handleError.HandleError(c, &handleError.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Update(dataProduct); err != nil {
		handleError.HandleError(c, err)
		return
	}
	response := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Cart successfully update",
	})
	c.JSON(http.StatusCreated, response)
}

func (h *cart_handler) DeleteCartHandler(c *gin.Context) {
	id := c.Param("id")
	var dataProduct models.Cart

	param_id, _ := strconv.Atoi(id)

	if err := h.service.Delete(param_id, dataProduct); err != nil {
		handleError.HandleError(c, err)
		return
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Product successfully deleted",
	})
	c.JSON(http.StatusOK, response)
}
func (h *cart_handler) ReadCartHandler(c *gin.Context) {
	cart_data, err := h.service.ReadCart()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cart": cart_data})

}
