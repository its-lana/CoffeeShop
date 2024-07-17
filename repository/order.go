package repository

import (
	"fmt"
	"strconv"
	"time"

	"github.com/its-lana/coffee-shop/common"
	"github.com/its-lana/coffee-shop/config"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository interface {
	CreateOrder(*dto.ReqOrder) (*dto.PayloadID, error)
	RetrieveAllOrder() ([]model.Order, error)
	RetrieveOrderByID(int) (*model.Order, error)
	RetrieveOrderByUID(string) (*model.Order, error)
	UpdateOrder(int, *model.Order) (*dto.PayloadID, error) //param : (orderID, *model.Order)
	CreateOrderNumber(string, int) (string, error)         // param : (orderType, merchantID)
}

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(config *config.GormDatabase) OrderRepository {
	return &orderRepository{
		DB: config.DB,
	}
}

func (cr *orderRepository) CreateOrderNumber(orderType string, merchantID int) (string, error) {
	var orderTypeCode string
	switch orderType {
	case common.TakeAwayOrder:
		orderTypeCode = "TA"
	default:
		orderTypeCode = "DI"
	}
	jakartaTimeZone, err := time.LoadLocation(common.JakartaLocation)
	if err != nil {
		return "", err
	}
	currentTime := time.Now().In(jakartaTimeZone)
	var count int64
	startOfDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, jakartaTimeZone)
	endOfDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 999999999, jakartaTimeZone)
	if err := cr.DB.Debug().Model(&model.Order{}).
		Where("merchant_id=?", merchantID).
		Where("order_date BETWEEN ? AND ?", startOfDay, endOfDay).
		Count(&count).Error; err != nil {
		return "", err
	}

	nextSequence := count + 1
	idLengkap := orderTypeCode + fmt.Sprintf("-0%d", merchantID) + strconv.Itoa(currentTime.Day()) + fmt.Sprintf("%02d", int(currentTime.Month())) + strconv.Itoa(currentTime.Year()) + strconv.Itoa(currentTime.Hour()) + strconv.Itoa(currentTime.Minute()) + strconv.Itoa(currentTime.Second()) + fmt.Sprintf("%03d", nextSequence)

	return idLengkap, nil
}

func (cr *orderRepository) RetrieveAllOrder() ([]model.Order, error) {
	var order []model.Order
	err := cr.DB.Debug().Preload(clause.Associations).Preload("OrderItem.Menu").Find(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (cr *orderRepository) RetrieveOrderByID(id int) (*model.Order, error) {
	var order model.Order
	err := cr.DB.Debug().Preload(clause.Associations).Preload("OrderItem.Menu").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (cr *orderRepository) RetrieveOrderByUID(uid string) (*model.Order, error) {
	var order model.Order
	err := cr.DB.Debug().Preload(clause.Associations).Preload("OrderItem.Menu").First(&order, "order_uid = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (cr *orderRepository) CreateOrder(req *dto.ReqOrder) (*dto.PayloadID, error) {
	model := helper.ToOrderModel(req)
	res := cr.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}

	return &dto.PayloadID{Id: model.ID}, nil
}

func (cr *orderRepository) UpdateOrder(id int, req *model.Order) (*dto.PayloadID, error) {
	res := cr.DB.Model(model.Order{}).Where("id = ?", id).Updates(&req)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: id}, nil
}
