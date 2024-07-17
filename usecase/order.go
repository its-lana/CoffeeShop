package usecase

import (
	"errors"
	"time"

	"github.com/its-lana/coffee-shop/common"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/model"
	"github.com/its-lana/coffee-shop/repository"
)

type OrderUseCase interface {
	RetrieveAllOrder() ([]dto.RespOrder, error)
	RetrieveCustomerOrder(int) ([]dto.RespOrder, error)
	PlaceOrder(*dto.ReqOrder) (*dto.RespOrder, error)
	UpdateOrderStatus(uid, orderCode string) (*dto.RespOrder, error)
}

type orderUseCase struct {
	orderRepository     repository.OrderRepository
	cartRepository      repository.CartRepository
	paymentRepository   repository.PaymentRepository
	orderItemRepository repository.OrderItemRepository
	customerRepository  repository.CustomerRepository
}

func NewOrderUseCase(oRepo repository.OrderRepository, cRepo repository.CartRepository, pRepo repository.PaymentRepository, oiRepo repository.OrderItemRepository, custRepo repository.CustomerRepository) OrderUseCase {
	return &orderUseCase{
		orderRepository:     oRepo,
		cartRepository:      cRepo,
		paymentRepository:   pRepo,
		orderItemRepository: oiRepo,
		customerRepository:  custRepo,
	}
}

func (pu *orderUseCase) RetrieveAllOrder() ([]dto.RespOrder, error) {
	orders, err := pu.orderRepository.RetrieveAllOrder()
	if err != nil {
		return nil, err
	}
	var resp []dto.RespOrder
	for _, order := range orders {
		resp = append(resp, *helper.ToResponseOrder(&order))
	}
	return resp, nil
}

func (pu *orderUseCase) RetrieveCustomerOrder(custID int) ([]dto.RespOrder, error) {
	orders, err := pu.orderRepository.RetrieveOrderByCustomerID(custID)
	if err != nil {
		return nil, err
	}
	var resp []dto.RespOrder
	for _, order := range orders {
		resp = append(resp, *helper.ToResponseOrder(&order))
	}
	return resp, nil
}

func (pu *orderUseCase) UpdateOrderStatus(uid, orderCode string) (*dto.RespOrder, error) {
	currentOrder, err := pu.orderRepository.RetrieveOrderByUID(uid)
	if err != nil {
		return nil, err
	}

	var updatingOrder model.Order
	switch currentOrder.OrderStatus {
	case common.OrderBaruStatus:
		updatingOrder.OrderStatus = common.DisiapkanStatus
	case common.DisiapkanStatus:
		updatingOrder.OrderStatus = common.SiapDiambilStatus
	case common.SiapDiambilStatus:
		if orderCode != currentOrder.OrderCode {
			return nil, errors.New("incorrect order code")
		}
		updatingOrder.OrderStatus = common.SelesaiStatus
	}

	_, err = pu.orderRepository.UpdateOrder(currentOrder.ID, &updatingOrder)
	if err != nil {
		return nil, err
	}

	currentOrder, err = pu.orderRepository.RetrieveOrderByID(currentOrder.ID)
	if err != nil {
		return nil, err
	}
	return helper.ToResponseOrder(currentOrder), nil
}

func (pu *orderUseCase) PlaceOrder(req *dto.ReqOrder) (*dto.RespOrder, error) {
	cart, err := pu.cartRepository.RetrieveCartByCustomerID(req.CustomerID)
	if err != nil {
		return nil, err
	}
	var finalAmount int
	for _, oItem := range cart.OrderItem {
		finalAmount += (oItem.Menu.Price * oItem.Quantity)
	}

	orderUID, err := pu.orderRepository.CreateOrderNumber(req.OrderType, cart.MerchantID)
	if err != nil {
		return nil, err
	}

	orderCode, err := helper.GenerateSecretOrderCode()
	if err != nil {
		return nil, err
	}

	req.OrderUID = orderUID
	req.OrderCode = orderCode
	req.OrderDate = time.Now()
	req.OrderStatus = common.BelumBayarStatus
	req.FinalAmount = finalAmount
	req.MerchantID = cart.MerchantID

	createdOrderID, err := pu.orderRepository.CreateOrder(req)
	if err != nil {
		return nil, err
	}

	customer, err := pu.customerRepository.RetrieveCustomerByID(req.CustomerID)

	transDetail := dto.ReqTransactionDetail{
		OrderUID:      orderUID,
		GrossAmount:   finalAmount,
		FullName:      customer.FullName,
		Email:         customer.Email,
		PaymentMethod: req.PaymentMethod,
	}

	snapResp, err := pu.paymentRepository.GetPaymentTokenSnap(&transDetail)

	reqPayment := dto.ReqPayment{
		PaymentMethod: req.PaymentMethod,
		OrderUID:      orderUID,
		CustomerID:    req.CustomerID,
		FinalAmount:   finalAmount,
		PaymentURL:    snapResp.RedirectURL,
	}
	_, err = pu.paymentRepository.CreatePayment(&reqPayment)
	if err != nil {
		return nil, err
	}

	err = pu.orderItemRepository.UpdateOwnerOrderItem(&dto.Owner{OwnerID: cart.ID, OwnerType: "cart"}, &dto.Owner{OwnerID: createdOrderID.Id, OwnerType: "order"})
	if err != nil {
		return nil, err
	}

	_, err = pu.cartRepository.UpdateCart(cart.ID, &dto.ReqCart{MerchantID: 0})
	if err != nil {
		return nil, err
	}

	newOrder, err := pu.orderRepository.RetrieveOrderByID(createdOrderID.Id)
	if err != nil {
		return nil, err
	}

	return helper.ToResponseOrder(newOrder), nil
}
