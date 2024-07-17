package usecase

import (
	"errors"

	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/repository"
	"gorm.io/gorm"
)

type CustomerUseCase interface {
	RetrieveAllCustomer() ([]dto.RespCustomer, error)
	CreateCustomer(*dto.ReqCustomer) (*dto.RespCustomer, error)
	RetrieveCustomerCart(int) (*dto.RespCart, error)
	AddItemToCart(int, *dto.ReqOrderItem) (*dto.RespCart, error) //param : (customerID, reqOrderItem)
	DeleteAllItemInCart(int) (*dto.RespCart, error)              //param : (customerID)
	DeleteOrderItemFromCart(int, int) (*dto.RespCart, error)     //param : (customerID, menuID )
}

type customerUseCase struct {
	customerRepository  repository.CustomerRepository
	cartRepository      repository.CartRepository
	orderItemRepository repository.OrderItemRepository
	menuRepository      repository.MenuRepository
}

func NewCustomerUseCase(custRepo repository.CustomerRepository, cartRepo repository.CartRepository, ordItem repository.OrderItemRepository, menuRepo repository.MenuRepository) CustomerUseCase {
	return &customerUseCase{
		customerRepository:  custRepo,
		cartRepository:      cartRepo,
		orderItemRepository: ordItem,
		menuRepository:      menuRepo,
	}
}

func (pu *customerUseCase) RetrieveAllCustomer() ([]dto.RespCustomer, error) {
	customers, err := pu.customerRepository.RetrieveAllCustomer()
	if err != nil {
		return nil, err
	}
	var resp []dto.RespCustomer
	for _, customer := range customers {
		resp = append(resp, *helper.ToResponseCustomer(&customer))
	}
	return resp, nil
}

func (pu *customerUseCase) RetrieveCustomerCart(custId int) (*dto.RespCart, error) {
	cart, err := pu.cartRepository.RetrieveCartByCustomerID(custId)
	if err != nil {
		return nil, err
	}
	return helper.ToResponseCart(cart), nil
}

func (cu *customerUseCase) DeleteAllItemInCart(custID int) (*dto.RespCart, error) {
	cart, err := cu.cartRepository.RetrieveCartByCustomerID(custID)
	if err != nil {
		return nil, err
	}
	// change merchant id in cart to 0
	_, err = cu.cartRepository.UpdateCart(cart.ID, &dto.ReqCart{MerchantID: 0})
	if err != nil {
		return nil, err
	}
	// delete order item where owner_type = cart and owner_id = cart.id
	err = cu.orderItemRepository.DeleteAllItemInCart(cart.ID)
	if err != nil {
		return nil, err
	}
	cart, err = cu.cartRepository.RetrieveCartByCustomerID(custID)
	if err != nil {
		return nil, err
	}
	return helper.ToResponseCart(cart), nil
}

func (cu *customerUseCase) DeleteOrderItemFromCart(custID, menuID int) (*dto.RespCart, error) {
	cart, err := cu.cartRepository.RetrieveCartByCustomerID(custID)
	if err != nil {
		return nil, err
	}
	_, err = cu.orderItemRepository.RetrieveOrderItemByMenuIDAndCartID(menuID, cart.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("the menu you want to delete is not in the cart")
		}
		return nil, err
	}
	// delete order item where owner_type = cart and owner_id = cart.id
	err = cu.orderItemRepository.DeleteOrderItemFromCart(menuID, cart.ID)
	if err != nil {
		return nil, err
	}
	cart, err = cu.cartRepository.RetrieveCartByCustomerID(custID)
	if err != nil {
		return nil, err
	}
	return helper.ToResponseCart(cart), nil
}

func (cu *customerUseCase) AddItemToCart(custID int, req *dto.ReqOrderItem) (*dto.RespCart, error) {

	menu, err := cu.menuRepository.RetrieveMenuByID(req.MenuID)
	if err != nil {
		return nil, err
	}

	merchantInMenuID, err := cu.menuRepository.RetrieveMerchantIDByMenuID(req.MenuID)
	if err != nil {
		return nil, err
	}

	if !menu.AvailabilityStatus {
		return nil, errors.New("the menu is currently not available")
	}

	cart, err := cu.cartRepository.RetrieveCartByCustomerID(custID)
	if err != nil {
		return nil, err
	}

	if cart.MerchantID > 0 && cart.MerchantID != merchantInMenuID.Id {
		return nil, errors.New("there are menus from another merchant in your cart! please delete it first")
	}

	isExist, err := cu.orderItemRepository.IsExistOrderItemByMenuIDAndCartID(req.MenuID, cart.ID)
	if err != nil {
		return nil, err
	}

	if isExist {
		orderItem, err := cu.orderItemRepository.RetrieveOrderItemByMenuIDAndCartID(req.MenuID, cart.ID)
		if err != nil {
			return nil, err
		}
		req.Quantity += orderItem.Quantity
		_, err = cu.orderItemRepository.UpdateOrderItem(orderItem.ID, req)
		if err != nil {
			return nil, err
		}
	} else {
		req.OwnerType = "cart"
		req.OwnerID = cart.ID
		_, err = cu.orderItemRepository.CreateOrderItem(req)
		if err != nil {
			return nil, err
		}
		_, err = cu.cartRepository.UpdateCart(cart.ID, &dto.ReqCart{MerchantID: merchantInMenuID.Id})
		if err != nil {
			return nil, err
		}
	}

	cart, err = cu.cartRepository.RetrieveCartByCustomerID(custID)
	if err != nil {
		return nil, err
	}

	return helper.ToResponseCart(cart), nil
}

func (pu *customerUseCase) CreateCustomer(req *dto.ReqCustomer) (*dto.RespCustomer, error) {
	data, err := pu.customerRepository.CreateCustomer(req)
	if err != nil {
		return nil, err
	}

	reqCart := dto.ReqCart{CustomerID: data.Id, MerchantID: 0}
	_, err = pu.cartRepository.CreateCart(&reqCart)
	if err != nil {
		return nil, err
	}

	model, err := pu.customerRepository.RetrieveCustomerByID(data.Id)
	if err != nil {
		return nil, err
	}

	resp := helper.ToResponseCustomer(model)

	return resp, nil
}
