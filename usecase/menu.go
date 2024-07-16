package usecase

import (
	"github.com/its-lana/coffee-shop/dto"
	"github.com/its-lana/coffee-shop/helper"
	"github.com/its-lana/coffee-shop/repository"
)

type MenuUseCase interface {
	RetrieveAllMenu() ([]dto.RespMenu, error)
	CreateMenu(*dto.ReqMenu) (*dto.RespMenu, error)
}

type menuUseCase struct {
	menuRepository repository.MenuRepository
}

func NewMenuUseCase(repo repository.MenuRepository) MenuUseCase {
	return &menuUseCase{
		menuRepository: repo,
	}
}

func (pu *menuUseCase) RetrieveAllMenu() ([]dto.RespMenu, error) {
	menus, err := pu.menuRepository.RetrieveAllMenu()
	if err != nil {
		return nil, err
	}
	var resp []dto.RespMenu
	for _, menu := range menus {
		resp = append(resp, *helper.ToResponseMenu(&menu))
	}
	return resp, nil
}

func (pu *menuUseCase) CreateMenu(req *dto.ReqMenu) (*dto.RespMenu, error) {
	data, err := pu.menuRepository.CreateMenu(req)
	if err != nil {
		return nil, err
	}

	model, err := pu.menuRepository.RetrieveMenuByID(data.Id)
	if err != nil {
		return nil, err
	}

	resp := helper.ToResponseMenu(model)

	return resp, nil
}
