package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/its-lana/coffee-shop/apperr"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/usecase"
)

type MenuHandler struct {
	menuUseCase usecase.MenuUseCase
}

func NewMenuHandler(cust usecase.MenuUseCase) *MenuHandler {
	return &MenuHandler{
		menuUseCase: cust,
	}
}

func (ch *MenuHandler) GetAllMenus(c *gin.Context) {
	res, err := ch.menuUseCase.RetrieveAllMenu()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve all menus, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("The menus record has been successfully retrieved", res))
}

func (ch *MenuHandler) AddNewMenu(c *gin.Context) {
	var req dto.ReqMenu
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(apperr.ErrInvalidBody.Message, apperr.ErrInvalidBody.Code))
		return
	}

	res, err := ch.menuUseCase.CreateMenu(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to register menu, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to register menu", res))
}
