package repository

import (
	"github.com/its-lana/coffee-shop/config"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CustomerRepository interface {
	CreateCustomer(*dto.ReqCustomer) (*dto.PayloadID, error)
	RetrieveAllCustomer() ([]model.Customer, error)
	RetrieveCustomerByID(id int) (*model.Customer, error)
}

type customerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(config *config.GormDatabase) CustomerRepository {
	return &customerRepository{
		DB: config.DB,
	}
}

func (sr *customerRepository) RetrieveAllCustomer() ([]model.Customer, error) {
	var customers []model.Customer
	err := sr.DB.Debug().Preload(clause.Associations).Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (sr *customerRepository) RetrieveCustomerByID(id int) (*model.Customer, error) {
	var customer model.Customer
	err := sr.DB.Debug().Preload(clause.Associations).First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (sr *customerRepository) CreateCustomer(req *dto.ReqCustomer) (*dto.PayloadID, error) {
	model := helper.ToCustomerModel(req)
	res := sr.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: model.ID}, nil
}
