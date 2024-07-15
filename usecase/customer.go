package usecase

import (
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/repository"
)

type CustomerUseCase interface {
	RetrieveAllCustomer() ([]dto.Customer, error)
	CreateCustomer(*dto.ReqCustomer) (*dto.Customer, error)
}

type customerUseCase struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerUseCase(sr repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		customerRepository: sr,
	}
}

func (pu *customerUseCase) RetrieveAllCustomer() ([]dto.Customer, error) {
	customers, err := pu.customerRepository.RetrieveAllCustomer()
	if err != nil {
		return nil, err
	}
	var resp []dto.Customer
	for _, customer := range customers {
		resp = append(resp, *helper.ToResponseCustomer(&customer))
	}
	return resp, nil
}

func (pu *customerUseCase) CreateCustomer(req *dto.ReqCustomer) (*dto.Customer, error) {
	data, err := pu.customerRepository.CreateCustomer(req)
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
