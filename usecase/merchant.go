package usecase

import (
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/repository"
)

type MerchantUseCase interface {
	RetrieveAllMerchant() ([]dto.RespMerchant, error)
	CreateMerchant(*dto.ReqMerchant) (*dto.RespMerchant, error)
}

type merchantUseCase struct {
	merchantRepository repository.MerchantRepository
}

func NewMerchantUseCase(repo repository.MerchantRepository) MerchantUseCase {
	return &merchantUseCase{
		merchantRepository: repo,
	}
}

func (pu *merchantUseCase) RetrieveAllMerchant() ([]dto.RespMerchant, error) {
	merchants, err := pu.merchantRepository.RetrieveAllMerchant()
	if err != nil {
		return nil, err
	}
	var resp []dto.RespMerchant
	for _, merchant := range merchants {
		resp = append(resp, *helper.ToResponseMerchant(&merchant))
	}
	return resp, nil
}

func (pu *merchantUseCase) CreateMerchant(req *dto.ReqMerchant) (*dto.RespMerchant, error) {
	data, err := pu.merchantRepository.CreateMerchant(req)
	if err != nil {
		return nil, err
	}

	model, err := pu.merchantRepository.RetrieveMerchantByID(data.Id)
	if err != nil {
		return nil, err
	}

	resp := helper.ToResponseMerchant(model)

	return resp, nil
}
