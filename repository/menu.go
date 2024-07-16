package repository

import (
	"github.com/its-lana/coffee-shop/config"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/model"
	"gorm.io/gorm"
)

type MenuRepository interface {
	CreateMenu(*dto.ReqMenu) (*dto.PayloadID, error)
	RetrieveAllMenu() ([]model.Menu, error)
	RetrieveMenuByID(int) (*model.Menu, error)
	UpdateMenu(int, *dto.ReqMenu) (*dto.PayloadID, error)
}

type menuRepository struct {
	DB *gorm.DB
}

func NewMenuRepository(config *config.GormDatabase) MenuRepository {
	return &menuRepository{
		DB: config.DB,
	}
}

func (mr *menuRepository) RetrieveAllMenu() ([]model.Menu, error) {
	var menu []model.Menu
	err := mr.DB.Debug().Find(&menu).Error
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (mr *menuRepository) RetrieveMenuByID(id int) (*model.Menu, error) {
	var menu model.Menu
	err := mr.DB.Debug().First(&menu, id).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (mr *menuRepository) CreateMenu(req *dto.ReqMenu) (*dto.PayloadID, error) {
	model := helper.ToMenuModel(req)
	res := mr.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: model.ID}, nil
}

func (mr *menuRepository) UpdateMenu(id int, req *dto.ReqMenu) (*dto.PayloadID, error) {
	data := helper.ToMenuModel(req)
	res := mr.DB.Model(model.Menu{}).Where("id = ?", id).Updates(&data)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: data.ID}, nil
}
