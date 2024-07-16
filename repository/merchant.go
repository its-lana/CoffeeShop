package repository

import (
	"github.com/its-lana/coffee-shop/config"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/model"
	"gorm.io/gorm"
)

type MerchantRepository interface {
	CreateMerchant(*dto.ReqMerchant) (*dto.PayloadID, error)
	RetrieveAllMerchant() ([]model.Merchant, error)
	RetrieveMerchantByID(int) (*model.Merchant, error)
	RetrieveMerchantByEmail(string) (*model.Merchant, error)
}

type merchantRepository struct {
	DB *gorm.DB
}

func NewMerchantRepository(config *config.GormDatabase) MerchantRepository {
	return &merchantRepository{
		DB: config.DB,
	}
}

func (cr *merchantRepository) RetrieveAllMerchant() ([]model.Merchant, error) {
	var merchants []model.Merchant
	err := cr.DB.Debug().Find(&merchants).Error
	if err != nil {
		return nil, err
	}
	return merchants, nil
}

func (cr *merchantRepository) RetrieveMerchantByID(id int) (*model.Merchant, error) {
	var merchant model.Merchant
	err := cr.DB.Debug().First(&merchant, id).Error
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (cr *merchantRepository) CreateMerchant(req *dto.ReqMerchant) (*dto.PayloadID, error) {
	model, err := helper.ToMerchantModel(req)
	if err != nil {
		return nil, err
	}
	res := cr.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: model.ID}, nil
}

func (cr *merchantRepository) RetrieveMerchantByEmail(email string) (*model.Merchant, error) {
	var merchant model.Merchant
	err := cr.DB.Debug().First(&merchant, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}
