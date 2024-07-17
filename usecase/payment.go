package usecase

import (
	"time"

	"github.com/its-lana/coffee-shop/common"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/model"
	"github.com/its-lana/coffee-shop/repository"
)

type PaymentUseCase interface {
	RetrieveAllPayment() ([]dto.RespPayment, error)
	PaymentNotification(*dto.ReqTransactionNotification) (*dto.RespPayment, error)
}

type paymentUseCase struct {
	paymentRepository repository.PaymentRepository
	orderRepository   repository.OrderRepository
}

func NewPaymentUseCase(pRepo repository.PaymentRepository, oRepo repository.OrderRepository) PaymentUseCase {
	return &paymentUseCase{
		paymentRepository: pRepo,
		orderRepository:   oRepo,
	}
}

func (pu *paymentUseCase) RetrieveAllPayment() ([]dto.RespPayment, error) {
	payments, err := pu.paymentRepository.RetrieveAllPayment()
	if err != nil {
		return nil, err
	}
	var resp []dto.RespPayment
	for _, payment := range payments {
		resp = append(resp, *helper.ToResponsePayment(&payment))
	}
	return resp, nil
}

func (pu *paymentUseCase) PaymentNotification(req *dto.ReqTransactionNotification) (*dto.RespPayment, error) {
	payment, err := pu.paymentRepository.RetrievePaymentByOrderUID(req.OrderID)
	if err != nil {
		return nil, err
	}

	order, err := pu.orderRepository.RetrieveOrderByUID(req.OrderID)
	if err != nil {
		return nil, err
	}

	if req.TransactionStatus == "settlement" {
		payment.Status = "paid"
		payment.PaidDate = time.Now()
		order.OrderStatus = common.OrderBaruStatus
	} else if req.TransactionStatus == "deny" || req.TransactionStatus == "expire" || req.TransactionStatus == "cancel" {
		payment.Status = "canceled"
		order.OrderStatus = common.TidakBayarStatus
	} else {
		payment.Status = "pending"
	}

	_, err = pu.orderRepository.UpdateOrder(order.ID, &model.Order{OrderStatus: order.OrderStatus})
	if err != nil {
		return nil, err
	}

	_, err = pu.paymentRepository.UpdatePayment(payment.ID, &model.Payment{Status: payment.Status, PaidDate: payment.PaidDate})
	if err != nil {
		return nil, err
	}

	model, err := pu.paymentRepository.RetrievePaymentByID(payment.ID)
	if err != nil {
		return nil, err
	}

	resp := helper.ToResponsePayment(model)

	return resp, nil
}
