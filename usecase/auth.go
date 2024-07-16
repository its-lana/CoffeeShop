package usecase

import (
	"errors"

	"github.com/its-lana/coffee-shop/apperr"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/repository"
)

type AuthUsecase interface {
	LoginCustomer(dto.LoginRequest) (*dto.RespCustomerLogin, error)
	LoginMerchant(dto.LoginRequest) (*dto.RespMerchantLogin, error)
}

type authUsecase struct {
	customerRepository repository.CustomerRepository
	merchantRepository repository.MerchantRepository
}

func NewAuthUsecase(custRepo repository.CustomerRepository, merchRepo repository.MerchantRepository) AuthUsecase {
	return &authUsecase{
		customerRepository: custRepo,
		merchantRepository: merchRepo,
	}
}

func (au *authUsecase) LoginCustomer(req dto.LoginRequest) (*dto.RespCustomerLogin, error) {
	customer, err := au.customerRepository.RetrieveCustomerByEmail(req.Email)
	if err != nil {
		return nil, apperr.ErrWrongCredentialsLogin
	}

	if !helper.ComparePassword(customer.Password, req.Password) {
		return nil, apperr.ErrWrongCredentialsLogin
	}
	token, err := dto.GenerateAccessToken(dto.JWTClaims{
		ID:   customer.ID,
		Role: "customer",
	})
	if err != nil {
		return nil, errors.New("error generate access token")
	}

	response := helper.ToResponseCustomerLogin(customer)
	response.Token = token

	return response, nil
}

func (au *authUsecase) LoginMerchant(req dto.LoginRequest) (*dto.RespMerchantLogin, error) {
	merchant, err := au.merchantRepository.RetrieveMerchantByEmail(req.Email)
	if err != nil {
		return nil, apperr.ErrWrongCredentialsLogin
	}

	if !helper.ComparePassword(merchant.Password, req.Password) {
		return nil, apperr.ErrWrongCredentialsLogin
	}
	token, err := dto.GenerateAccessToken(dto.JWTClaims{
		ID:   merchant.ID,
		Role: "merchant",
	})
	if err != nil {
		return nil, errors.New("error generate access token")
	}

	response := helper.ToResponseMerchantLogin(merchant)
	response.Token = token

	return response, nil
}
