package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/its-lana/coffee-shop/apperr"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/usecase"
)

type AuthHandler struct {
	authCustomerUsecase usecase.AuthCustomerUsecase
}

func NewAuthHandler(ac usecase.AuthCustomerUsecase) *AuthHandler {
	return &AuthHandler{
		authCustomerUsecase: ac,
	}
}

func (a *AuthHandler) LoginCustomer(c *gin.Context) {
	var loginInfo dto.LoginRequest
	err := c.ShouldBindJSON(&loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(apperr.ErrInvalidBody.Message, apperr.ErrInvalidBody.Code))
		return
	}

	response, err := a.authCustomerUsecase.Login(loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResponseFailed(err.Error(), http.StatusUnauthorized))
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("login successfully", response))
}
