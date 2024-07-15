package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/usecase"
)

type CustomerHandler struct {
	customerUseCase usecase.CustomerUseCase
}

func NewCustomerHandler(stud usecase.CustomerUseCase) *CustomerHandler {
	return &CustomerHandler{
		customerUseCase: stud,
	}
}

func (sh *CustomerHandler) GetAllCustomers(c *gin.Context) {
	res, err := sh.customerUseCase.RetrieveAllCustomer()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve all customers "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("The customers record has been successfully retrieved", res))
}
