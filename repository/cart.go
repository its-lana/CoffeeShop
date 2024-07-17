package repository

import (
	"github.com/its-lana/coffee-shop/config"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CartRepository interface {
	CreateCart(*dto.ReqCart) (*dto.PayloadID, error)
	RetrieveCartByCustomerID(int) (*model.Cart, error)
	RetrieveCartIDByCustomerID(int) (*dto.PayloadID, error)
	UpdateCart(int, *dto.ReqCart) (*dto.PayloadID, error)
}

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(config *config.GormDatabase) CartRepository {
	return &cartRepository{
		DB: config.DB,
	}
}

func (mr *cartRepository) RetrieveCartIDByCustomerID(custID int) (*dto.PayloadID, error) {
	var cartID int
	err := mr.DB.Debug().Model(&model.Cart{}).Where("customer_id = ?", custID).Pluck("id", &cartID).Error
	if err != nil {
		return nil, err
	}
	return &dto.PayloadID{Id: cartID}, nil
}

func (mr *cartRepository) RetrieveCartByCustomerID(custID int) (*model.Cart, error) {
	var cart model.Cart
	err := mr.DB.Debug().Preload(clause.Associations).Preload("OrderItem.Menu").First(&cart, "customer_id = ?", custID).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (mr *cartRepository) CreateCart(req *dto.ReqCart) (*dto.PayloadID, error) {
	model := helper.ToCartModel(req)
	res := mr.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: model.ID}, nil
}

func (mr *cartRepository) UpdateCart(cartID int, req *dto.ReqCart) (*dto.PayloadID, error) {
	res := mr.DB.Model(model.Cart{}).Where("id = ?", cartID).Updates(&req)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: cartID}, nil
}
