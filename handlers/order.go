package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/its-lana/coffee-shop/apperr"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/usecase"
)

type OrderHandler struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderHandler(order usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: order,
	}
}

func (ch *OrderHandler) GetAllOrders(c *gin.Context) {
	res, err := ch.orderUseCase.RetrieveAllOrder()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve all orders, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("The orders record has been successfully retrieved", res))
}

func (ch *OrderHandler) PlaceNewOrder(c *gin.Context) {
	var req dto.ReqOrder
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(apperr.ErrInvalidBody.Message, apperr.ErrInvalidBody.Code))
		return
	}

	res, err := ch.orderUseCase.PlaceOrder(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to place order, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to place new order", res))
}
