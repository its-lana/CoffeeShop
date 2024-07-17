package repository

import (
	"github.com/its-lana/coffee-shop/config"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/model"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	CreateOrderItem(*dto.ReqOrderItem) (*dto.PayloadID, error)
	RetrieveAllOrderItem() ([]model.OrderItem, error)
	RetrieveOrderItemByID(int) (*model.OrderItem, error)
	RetrieveOrderItemByMenuIDAndCartID(int, int) (*model.OrderItem, error) //param (menuID, cartID)
	IsExistOrderItemByMenuIDAndCartID(int, int) (bool, error)              //param (menuID, cartID)
	UpdateOrderItem(int, *dto.ReqOrderItem) (*dto.PayloadID, error)
	DeleteOrderItemFromCart(int, int) error //param (menuID, cartID)
	DeleteAllItemInCart(int) error
}

type orderItemRepository struct {
	DB *gorm.DB
}

func NewOrderItemRepository(config *config.GormDatabase) OrderItemRepository {
	return &orderItemRepository{
		DB: config.DB,
	}
}

func (mr *orderItemRepository) RetrieveAllOrderItem() ([]model.OrderItem, error) {
	var orderItem []model.OrderItem
	err := mr.DB.Debug().Find(&orderItem).Error
	if err != nil {
		return nil, err
	}
	return orderItem, nil
}

func (mr *orderItemRepository) RetrieveOrderItemByID(id int) (*model.OrderItem, error) {
	var orderItem model.OrderItem
	err := mr.DB.Debug().First(&orderItem, id).Error
	if err != nil {
		return nil, err
	}
	return &orderItem, nil
}

func (mr *orderItemRepository) RetrieveOrderItemByMenuIDAndCartID(menuID, cartID int) (*model.OrderItem, error) {
	var orderItem model.OrderItem
	err := mr.DB.Debug().First(&orderItem, "menu_id = ? AND owner_type = ? AND owner_id = ?", menuID, "cart", cartID).Error
	if err != nil {
		return nil, err
	}
	return &orderItem, nil
}

func (mr *orderItemRepository) IsExistOrderItemByMenuIDAndCartID(menuID, cartID int) (bool, error) {
	var count int64
	err := mr.DB.Debug().Model(&model.OrderItem{}).Where("menu_id = ? AND owner_type = ? AND owner_id = ?", menuID, "cart", cartID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (mr *orderItemRepository) CreateOrderItem(req *dto.ReqOrderItem) (*dto.PayloadID, error) {
	model := helper.ToOrderItemModel(req)
	res := mr.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}

	return &dto.PayloadID{Id: model.ID}, nil
}

func (mr *orderItemRepository) UpdateOrderItem(id int, req *dto.ReqOrderItem) (*dto.PayloadID, error) {
	data := helper.ToOrderItemModel(req)
	res := mr.DB.Model(model.OrderItem{}).Where("id = ?", id).Updates(&data)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: data.ID}, nil
}

func (mr *orderItemRepository) DeleteAllItemInCart(cartID int) error {
	err := mr.DB.Delete(&model.OrderItem{}, "owner_id = ? AND owner_type = ?", cartID, "cart").Error
	if err != nil {
		return err
	}
	return nil
}

func (mr *orderItemRepository) DeleteOrderItemFromCart(menuID, cartID int) error {
	err := mr.DB.Delete(&model.OrderItem{}, "menu_id = ? AND owner_id = ? AND owner_type = ?", menuID, cartID, "cart").Error
	if err != nil {
		return err
	}
	return nil
}
