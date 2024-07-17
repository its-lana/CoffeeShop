package repository

import (
	"github.com/its-lana/coffee-shop/config"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/model"
	"github.com/veritrans/go-midtrans"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaymentRepository interface {
	CreatePayment(*dto.ReqPayment) (*dto.PayloadID, error)
	RetrieveAllPayment() ([]model.Payment, error)
	RetrievePaymentByID(int) (*model.Payment, error)
	RetrievePaymentByOrderUID(string) (*model.Payment, error)
	UpdatePayment(int, *model.Payment) (*dto.PayloadID, error)
	GetPaymentTokenSnap(*dto.ReqTransactionDetail) (*midtrans.SnapResponse, error)
}

type paymentRepository struct {
	DB        *gorm.DB
	MidClient *midtrans.Client
}

func NewPaymentRepository(config *config.GormDatabase, midCLient *config.MidClient) PaymentRepository {
	return &paymentRepository{
		DB:        config.DB,
		MidClient: midCLient.Client,
	}
}

func (cr *paymentRepository) CreatePayment(req *dto.ReqPayment) (*dto.PayloadID, error) {
	model := helper.ToPaymentModel(req)
	res := cr.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}

	return &dto.PayloadID{Id: model.ID}, nil
}

func (cr *paymentRepository) UpdatePayment(id int, req *model.Payment) (*dto.PayloadID, error) {
	res := cr.DB.Model(model.Payment{}).Where("id = ?", id).Updates(&req)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: id}, nil
}

func (cr *paymentRepository) GetPaymentTokenSnap(transaction *dto.ReqTransactionDetail) (*midtrans.SnapResponse, error) {

	snapGateway := midtrans.SnapGateway{
		Client: *cr.MidClient,
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.OrderUID,
			GrossAmt: int64(transaction.GrossAmount),
		},
		CustomerDetail: &midtrans.CustDetail{
			Email: transaction.Email,
			FName: transaction.FullName,
		},
		Expiry: &midtrans.ExpiryDetail{
			Duration: 60,
			Unit:     "minute",
		},
		EnabledPayments: []midtrans.PaymentType{midtrans.PaymentType(transaction.PaymentMethod)},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return nil, err
	}

	return &snapTokenResp, nil
}

func (cr *paymentRepository) RetrieveAllPayment() ([]model.Payment, error) {
	var payment []model.Payment
	err := cr.DB.Debug().Preload(clause.Associations).Find(&payment).Error
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (cr *paymentRepository) RetrievePaymentByID(id int) (*model.Payment, error) {
	var payment model.Payment
	err := cr.DB.Debug().Preload(clause.Associations).First(&payment, id).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (cr *paymentRepository) RetrievePaymentByOrderUID(uid string) (*model.Payment, error) {
	var payment model.Payment
	err := cr.DB.Debug().Preload(clause.Associations).First(&payment, "order_uid = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}
