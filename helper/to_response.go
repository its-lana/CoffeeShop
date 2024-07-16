package helper

import (
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/model"
)

func ToResponseCustomer(cust *model.Customer) *dto.RespCustomer {
	return &dto.RespCustomer{
		ID:          cust.ID,
		FullName:    cust.FullName,
		Email:       cust.Email,
		PhoneNumber: cust.PhoneNumber,
	}
}

func ToResponseCustomerLogin(cust *model.Customer) *dto.RespCustomerLogin {
	return &dto.RespCustomerLogin{
		ID:          cust.ID,
		FullName:    cust.FullName,
		Email:       cust.Email,
		PhoneNumber: cust.PhoneNumber,
	}
}

func ToResponseMerchant(merchant *model.Merchant) *dto.RespMerchant {
	return &dto.RespMerchant{
		ID:           merchant.ID,
		MerchantName: merchant.MerchantName,
		Address:      merchant.Address,
		PICName:      merchant.PICName,
		Email:        merchant.Email,
		PhoneNumber:  merchant.PhoneNumber,
		IsOpen:       merchant.IsOpen,
	}
}

func ToResponseMerchantLogin(merchant *model.Merchant) *dto.RespMerchantLogin {
	return &dto.RespMerchantLogin{
		ID:           merchant.ID,
		MerchantName: merchant.MerchantName,
		Address:      merchant.Address,
		PICName:      merchant.PICName,
		Email:        merchant.Email,
		PhoneNumber:  merchant.PhoneNumber,
		IsOpen:       merchant.IsOpen,
	}
}
