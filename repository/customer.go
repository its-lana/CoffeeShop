package repository

import (
	"github.com/its-lana/coffee-shop/config"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(*dto.ReqCustomer) (*dto.PayloadID, error)
	RetrieveAllCustomer() ([]model.Customer, error)
	RetrieveCustomerByID(int) (*model.Customer, error)
	RetrieveCustomerByEmail(string) (*model.Customer, error)
}

type customerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(config *config.GormDatabase) CustomerRepository {
	return &customerRepository{
		DB: config.DB,
	}
}

func (cr *customerRepository) RetrieveAllCustomer() ([]model.Customer, error) {
	var customers []model.Customer
	err := cr.DB.Debug().Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (cr *customerRepository) RetrieveCustomerByID(id int) (*model.Customer, error) {
	var customer model.Customer
	err := cr.DB.Debug().First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (cr *customerRepository) CreateCustomer(req *dto.ReqCustomer) (*dto.PayloadID, error) {
	model, err := helper.ToCustomerModel(req)
	if err != nil {
		return nil, err
	}
	res := cr.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: model.ID}, nil
}

func (cr *customerRepository) RetrieveCustomerByEmail(email string) (*model.Customer, error) {
	var customer model.Customer
	err := cr.DB.Debug().First(&customer, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
