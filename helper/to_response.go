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

func ToResponseMenu(menu *model.Menu) *dto.RespMenu {
	return &dto.RespMenu{
		ID:                 menu.ID,
		ProductName:        menu.ProductName,
		Price:              menu.Price,
		Description:        menu.Description,
		ProductCode:        menu.ProductCode,
		ProductImage:       menu.ProductImage,
		CategoryID:         menu.CategoryID,
		AvailabilityStatus: menu.AvailabilityStatus,
	}
}

func ToResponseCategory(category *model.Category) *dto.RespCategory {
	var menus []dto.RespMenu
	for _, menu := range category.Menus {
		menus = append(menus, *ToResponseMenu(&menu))
	}
	return &dto.RespCategory{
		ID:           category.ID,
		CategoryName: category.CategoryName,
		MerchantID:   category.MerchantID,
		Menus:        menus,
	}
}
