package repository

import (
	"github.com/its-lana/coffee-shop/config"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepository interface {
	CreateCategory(*dto.ReqCategory) (*dto.PayloadID, error)
	RetrieveAllCategory() ([]model.Category, error)
	RetrieveCategoryByID(int) (*model.Category, error)
	UpdateCategory(int, *dto.ReqCategory) (*dto.PayloadID, error)
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(config *config.GormDatabase) CategoryRepository {
	return &categoryRepository{
		DB: config.DB,
	}
}

func (cr *categoryRepository) RetrieveAllCategory() ([]model.Category, error) {
	var category []model.Category
	err := cr.DB.Debug().Preload(clause.Associations).Find(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (cr *categoryRepository) RetrieveCategoryByID(id int) (*model.Category, error) {
	var category model.Category
	err := cr.DB.Debug().Preload(clause.Associations).First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (cr *categoryRepository) CreateCategory(req *dto.ReqCategory) (*dto.PayloadID, error) {
	model := helper.ToCategoryModel(req)
	res := cr.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}

	return &dto.PayloadID{Id: model.ID}, nil
}

func (cr *categoryRepository) UpdateCategory(id int, req *dto.ReqCategory) (*dto.PayloadID, error) {
	data := helper.ToCategoryModel(req)
	res := cr.DB.Model(model.Category{}).Where("id = ?", id).Updates(&data)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: data.ID}, nil
}
