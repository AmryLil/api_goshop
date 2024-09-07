package services

import (
	"api_goshop/dto"
	"api_goshop/handleError"
	"api_goshop/models"
	"api_goshop/repositories"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentService interface {
	CreatePayment(itemDetailsPayment []dto.ItemDetailsRequest, userID *int) (*snap.Response, error)
}

type payment_service struct {
	repository repositories.PaymentRepo
	snapClient *snap.Client
}

func NewPaymentService(repository repositories.PaymentRepo) *payment_service {
	client := snap.Client{}
	client.New("SB-Mid-server-0BPSoq6ZokfN6eJlm49W38mF", midtrans.Sandbox) // Ganti dengan kunci rahasia kamu
	// Jika ingin menggunakan lingkungan produksi, ganti midtrans.Sandbox dengan midtrans.Production

	return &payment_service{
		repository: repository,
		snapClient: &client,
	}
}

func (service *payment_service) CreatePayment(itemDetailsPayment []dto.ItemDetailsRequest, userID *int) (*snap.Response, error) {

	// Ambil data pengguna dari repository
	user, err := service.repository.FindUserByID(userID)
	if err != nil {
		return nil, &handleError.InternalServerError{Message: err.Error()}
	}
	var amount int
	var items []midtrans.ItemDetails
	for _, itemDetails := range itemDetailsPayment {
		item := midtrans.ItemDetails{
			ID:    strconv.Itoa(itemDetails.ProductID),
			Name:  itemDetails.Name,
			Price: int64(itemDetails.Price),
			Qty:   int32(itemDetails.QTY),
		}
		items = append(items, item)
		amount += itemDetails.Price * itemDetails.QTY
	}

	// Buat parameter permintaan untuk transaksi Snap
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  uuid.NewString(), // Atur OrderID berdasarkan UUID baru
			GrossAmt: int64(amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		Items: &items,
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Firstname,
			Email: user.Email,
			Phone: user.PhoneNumber,
		},
	}
	// Buat transaksi Snap
	snapResp, err := service.snapClient.CreateTransaction(req)
	fmt.Println("id :", userID)
	fmt.Println("total : ", amount)
	payment := models.Payment{
		Amount: float64(amount),
		UserID: *userID,
		Status: "pending",
	}
	err = service.repository.CreatePayment(payment)
	if err != nil {
		return nil, &handleError.InternalServerError{Message: err.Error()}
	}
	var paymentDetails []models.ItemDetails
	for _, itemDetails := range itemDetailsPayment {
		itemDetail := models.ItemDetails{
			ProudctID: itemDetails.ProductID,
			QTY:       itemDetails.QTY,
			Name:      itemDetails.Name,
			Price:     itemDetails.Price,
		}
		paymentDetails = append(paymentDetails, itemDetail)
		err = service.repository.CreateItemDetailsPayment(paymentDetails)
		if err != nil {
			return nil, &handleError.InternalServerError{Message: err.Error()}
		}
	}
	if err != nil {
		return nil, &handleError.InternalServerError{Message: err.Error()}
	}

	return snapResp, nil
}
