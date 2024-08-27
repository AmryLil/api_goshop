package handler

import (
	"api_goshop/dto"
	"api_goshop/handleError"
	"api_goshop/helper"
	"api_goshop/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type useraccounthandler struct {
	service services.Service
}

func NewUserAccount(service services.Service) *useraccounthandler {
	return &useraccounthandler{service}
}
func (h *useraccounthandler) RegisterHandler(c *gin.Context) {
	var user dto.RegisterRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		handleError.HandleError(c, &handleError.BadRequestError{Message: err.Error()})
		return
	}
	if err := h.service.Register(user); err != nil {
		handleError.HandleError(c, err)
		return
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register Success, please login",
	})
	c.JSON(http.StatusCreated, response)

}

func (h *useraccounthandler) GetDataUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	data, err := h.service.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        data.Id,
		"firtname":  data.Firstname,
		"lastname":  data.Lastname,
		"email":     data.Email,
		"username":  data.Username,
		"create_at": data.CreatedAt,
	})
}

func (h *useraccounthandler) LoginHandler(c *gin.Context) {
	var user dto.LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		handleError.HandleError(c, &handleError.BadRequestError{Message: err.Error()})
		return
	}

	token, res, err := h.service.Login(&user)

	if err != nil {
		handleError.HandleError(c, err)
		return
	}

	// shddjhdahhjajdjv

	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",         // Pastikan path sudah benar
		Domain:   "localhost", // Pastikan domain sesuai
		Expires:  expiration,
		HttpOnly: false,
		Secure:   false, // Ubah ke true jika menggunakan HTTPS
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(c.Writer, &cookie)

	// Tambahkan logging
	log.Printf("Set Cookie: %v\n", cookie)

	response := dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Successfully Login",
		Data:       res,
	}
	c.JSON(http.StatusOK, response)
}

func (h *useraccounthandler) GetAllUser(c *gin.Context) {
	userAccount, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"useraccount": userAccount})

}
