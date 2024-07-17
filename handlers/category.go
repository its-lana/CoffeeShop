package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/its-lana/coffee-shop/apperr"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/usecase"
)

type CategoryHandler struct {
	categoryUseCase usecase.CategoryUseCase
}

func NewCategoryHandler(cust usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{
		categoryUseCase: cust,
	}
}

func (ch *CategoryHandler) GetAllCategories(c *gin.Context) {
	res, err := ch.categoryUseCase.RetrieveAllCategory()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve all categories, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("The categories record has been successfully retrieved", res))
}

func (ch *CategoryHandler) AddNewCategory(c *gin.Context) {
	var req dto.ReqCategory
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(apperr.ErrInvalidBody.Message, apperr.ErrInvalidBody.Code))
		return
	}

	res, err := ch.categoryUseCase.CreateCategory(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to add category, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to add new category", res))
}
