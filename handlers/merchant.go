package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/its-lana/coffee-shop/apperr"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/usecase"
)

type MerchantHandler struct {
	merchantUseCase usecase.MerchantUseCase
}

func NewMerchantHandler(cust usecase.MerchantUseCase) *MerchantHandler {
	return &MerchantHandler{
		merchantUseCase: cust,
	}
}

func (ch *MerchantHandler) GetAllMerchants(c *gin.Context) {
	res, err := ch.merchantUseCase.RetrieveAllMerchant()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve all merchants "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("The merchants record has been successfully retrieved", res))
}

func (ch *MerchantHandler) RegisterMerchant(c *gin.Context) {
	var req dto.ReqMerchant
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(apperr.ErrInvalidBody.Message, apperr.ErrInvalidBody.Code))
		return
	}

	res, err := ch.merchantUseCase.CreateMerchant(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to register merchants "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to register merchants ", res))
}
