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

func ToResponseOrderItem(ordItem *model.OrderItem) *dto.RespOrderItem {
	return &dto.RespOrderItem{
		ID:           ordItem.ID,
		MenuID:       ordItem.MenuID,
		Quantity:     ordItem.Quantity,
		OwnerID:      ordItem.OwnerID,
		OwnerType:    ordItem.OwnerType,
		MerchantMenu: ToResponseMenu(&ordItem.Menu),
	}
}

func ToResponseCart(cart *model.Cart) *dto.RespCart {
	var orderItems []dto.RespOrderItem
	var totalAmount int
	for _, ordItem := range cart.OrderItem {
		orderItems = append(orderItems, *ToResponseOrderItem(&ordItem))
		totalAmount += (int(ordItem.Menu.Price) * ordItem.Quantity)
	}

	return &dto.RespCart{
		ID:          cart.ID,
		CustomerID:  cart.CustomerID,
		MerchantID:  cart.MerchantID,
		OrderItems:  orderItems,
		TotalAmount: totalAmount,
	}
}

func ToResponsePayment(payment *model.Payment) *dto.RespPayment {
	return &dto.RespPayment{
		ID:            payment.ID,
		CustomerID:    payment.CustomerID,
		OrderUID:      payment.OrderUID,
		PaymentAmount: payment.PaymentAmount,
		Status:        payment.Status,
		PaymentMethod: payment.PaymentMethod,
		PaymentURL:    payment.PaymentURL,
		PaidDate:      payment.PaidDate,
	}
}

func ToResponseOrder(o *model.Order) *dto.RespOrder {
	var orderItems []dto.RespOrderItem
	for _, ordItem := range o.OrderItem {
		orderItems = append(orderItems, *ToResponseOrderItem(&ordItem))
	}
	return &dto.RespOrder{
		ID:          o.ID,
		OrderUID:    o.OrderUID,
		CustomerID:  o.CustomerID,
		MerchantID:  o.MerchantID,
		FinalAmount: o.FinalAmount,
		OrderType:   o.OrderType,
		OrderNotes:  o.OrderNotes,
		OrderStatus: o.OrderStatus,
		NoteStatus:  o.NoteStatus,
		OrderDate:   o.OrderDate,
		Payment:     *ToResponsePayment(&o.Payment),
		OrderCode:   o.OrderCode,
		OrderItems:  orderItems,
	}
}
