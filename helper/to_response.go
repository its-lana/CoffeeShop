package helper

import (
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/model"
)

func ToResponseCustomer(cust *model.Customer) *dto.Customer {
	return &dto.Customer{
		ID:          cust.ID,
		FullName:    cust.FullName,
		Email:       cust.Email,
		PhoneNumber: cust.PhoneNumber,
	}
}

func ToResponseCustomerLogin(cust *model.Customer) *dto.CustomerLogin {
	return &dto.CustomerLogin{
		ID:          cust.ID,
		FullName:    cust.FullName,
		Email:       cust.Email,
		PhoneNumber: cust.PhoneNumber,
	}
}
