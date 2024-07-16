package usecase

import (
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/repository"
)

type CategoryUseCase interface {
	RetrieveAllCategory() ([]dto.RespCategory, error)
	CreateCategory(*dto.ReqCategory) (*dto.RespCategory, error)
}

type categoryUseCase struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUseCase(repo repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{
		categoryRepository: repo,
	}
}

func (pu *categoryUseCase) RetrieveAllCategory() ([]dto.RespCategory, error) {
	categorys, err := pu.categoryRepository.RetrieveAllCategory()
	if err != nil {
		return nil, err
	}
	var resp []dto.RespCategory
	for _, category := range categorys {
		resp = append(resp, *helper.ToResponseCategory(&category))
	}
	return resp, nil
}

func (pu *categoryUseCase) CreateCategory(req *dto.ReqCategory) (*dto.RespCategory, error) {
	data, err := pu.categoryRepository.CreateCategory(req)
	if err != nil {
		return nil, err
	}

	model, err := pu.categoryRepository.RetrieveCategoryByID(data.Id)
	if err != nil {
		return nil, err
	}

	resp := helper.ToResponseCategory(model)

	return resp, nil
}
