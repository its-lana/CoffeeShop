package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/its-lana/coffee-shop/apperr"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/usecase"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(ac usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: ac,
	}
}

func (a *AuthHandler) LoginCustomer(c *gin.Context) {
	var loginInfo dto.LoginRequest
	err := c.ShouldBindJSON(&loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(apperr.ErrInvalidBody.Message, apperr.ErrInvalidBody.Code))
		return
	}

	response, err := a.authUsecase.LoginCustomer(loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResponseFailed(err.Error(), http.StatusUnauthorized))
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("login successfully", response))
}

func (a *AuthHandler) LoginMerchant(c *gin.Context) {
	var loginInfo dto.LoginRequest
	err := c.ShouldBindJSON(&loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(apperr.ErrInvalidBody.Message, apperr.ErrInvalidBody.Code))
		return
	}

	response, err := a.authUsecase.LoginMerchant(loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResponseFailed(err.Error(), http.StatusUnauthorized))
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("login successfully", response))
}
