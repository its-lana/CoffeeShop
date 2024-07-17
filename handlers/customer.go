package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/its-lana/coffee-shop/apperr"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/usecase"
)

type CustomerHandler struct {
	customerUseCase usecase.CustomerUseCase
}

func NewCustomerHandler(cust usecase.CustomerUseCase) *CustomerHandler {
	return &CustomerHandler{
		customerUseCase: cust,
	}
}

func (ch *CustomerHandler) RetrieveAllCustomer(c *gin.Context) {
	res, err := ch.customerUseCase.RetrieveAllCustomer()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve all customers "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("The customers record has been successfully retrieved", res))
}

func (ch *CustomerHandler) RetrieveCustomerCart(c *gin.Context) {
	custID, err := strconv.Atoi(c.Param("cust-id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed("customer id is invalid", http.StatusBadRequest))
		return
	}

	res, err := ch.customerUseCase.RetrieveCustomerCart(custID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve customer cart "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("The customer cart has been successfully retrieved", res))
}

func (ch *CustomerHandler) DeleteAllItemInCart(c *gin.Context) {
	custID, err := strconv.Atoi(c.Param("cust-id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed("customer id is invalid", http.StatusBadRequest))
		return
	}

	res, err := ch.customerUseCase.DeleteAllItemInCart(custID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to delete all item in cart "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to delete all item in cart ", res))
}

func (ch *CustomerHandler) DeleteOrderItemFromCart(c *gin.Context) {
	custID, err := strconv.Atoi(c.Param("cust-id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed("customer id is invalid", http.StatusBadRequest))
		return
	}

	menuID, err := strconv.Atoi(c.Param("menu-id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed("menu id is invalid", http.StatusBadRequest))
		return
	}

	res, err := ch.customerUseCase.DeleteOrderItemFromCart(custID, menuID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to delete item from cart "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to delete item from cart ", res))
}

func (ch *CustomerHandler) AddItemToCart(c *gin.Context) {
	var req dto.ReqOrderItem
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(apperr.ErrInvalidBody.Message, apperr.ErrInvalidBody.Code))
		return
	}

	custID, err := strconv.Atoi(c.Param("cust-id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed("customer id is invalid", http.StatusBadRequest))
		return
	}

	res, err := ch.customerUseCase.AddItemToCart(custID, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to add item to cart, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("The order item has been successfully added to cart", res))
}

func (ch *CustomerHandler) RegisterCustomer(c *gin.Context) {
	var req dto.ReqCustomer
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(apperr.ErrInvalidBody.Message, apperr.ErrInvalidBody.Code))
		return
	}

	res, err := ch.customerUseCase.CreateCustomer(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to register customers "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to register customers ", res))
}
