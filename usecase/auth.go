package usecase

import (
	"errors"

	"github.com/its-lana/coffee-shop/apperr"
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/repository"
)

type AuthCustomerUsecase interface {
	Login(dto.LoginRequest) (*dto.CustomerLogin, error)
}

type authCustomerUsecase struct {
	customerRepository repository.CustomerRepository
}

func NewAuthCustomerUsecase(sr repository.CustomerRepository) AuthCustomerUsecase {
	return &authCustomerUsecase{
		customerRepository: sr,
	}
}

func (au *authCustomerUsecase) Login(req dto.LoginRequest) (*dto.CustomerLogin, error) {
	// var response dto.CustomerLogin
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
