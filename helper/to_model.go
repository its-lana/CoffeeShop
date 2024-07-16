package helper

import (
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/model"
)

func ToCustomerModel(req *dto.ReqCustomer) (*model.Customer, error) {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	return &model.Customer{
		FullName:    req.FullName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    hashedPassword,
	}, nil
}

func ToMerchantModel(req *dto.ReqMerchant) (*model.Merchant, error) {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	return &model.Merchant{
		MerchantName: req.MerchantName,
		Address:      req.Address,
		PICName:      req.PICName,
		Email:        req.Email,
		Password:     hashedPassword,
		PhoneNumber:  req.PhoneNumber,
		IsOpen:       req.IsOpen,
	}, nil
}

func ToMenuModel(req *dto.ReqMenu) *model.Menu {
	return &model.Menu{
		ProductName:        req.ProductName,
		Price:              req.Price,
		Description:        req.Description,
		ProductCode:        req.ProductCode,
		ProductImage:       req.ProductImage,
		CategoryID:         req.CategoryID,
		AvailabilityStatus: req.AvailabilityStatus,
	}
}

func ToCategoryModel(req *dto.ReqCategory) *model.Category {
	return &model.Category{
		CategoryName: req.CategoryName,
		MerchantID:   req.MerchantID,
	}
}
