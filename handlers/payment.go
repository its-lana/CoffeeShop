package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/its-lana/coffee-shop/apperr"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/usecase"
)

type PaymentHandler struct {
	paymentUseCase usecase.PaymentUseCase
}

func NewPaymentHandler(payment usecase.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: payment,
	}
}

func (ch *PaymentHandler) GetAllPayments(c *gin.Context) {
	res, err := ch.paymentUseCase.RetrieveAllPayment()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve all payments, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("The payments record has been successfully retrieved", res))
}

func (ch *PaymentHandler) PaymentNotification(c *gin.Context) {
	var req dto.ReqTransactionNotification
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(apperr.ErrInvalidBody.Message, apperr.ErrInvalidBody.Code))
		return
	}

	res, err := ch.paymentUseCase.PaymentNotification(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to process payment, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to process payment", res))
}
