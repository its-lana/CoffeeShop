package usecase

import (
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/repository"
)

type CustomerUseCase interface {
	RetrieveAllCustomer() ([]dto.RespCustomer, error)
	CreateCustomer(*dto.ReqCustomer) (*dto.RespCustomer, error)
}

type customerUseCase struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		customerRepository: repo,
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

func (pu *customerUseCase) CreateCustomer(req *dto.ReqCustomer) (*dto.RespCustomer, error) {
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
