package service

import (
	"errors"
	"immersive-dashboard/features/team"

	"github.com/go-playground/validator/v10"
)

type teamService struct {
	teamRepository team.RepositoryInterface
	validate       *validator.Validate
}

func New(repo team.RepositoryInterface) team.ServiceInterface {
	return &teamService{
		teamRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *teamService) Create(input team.CoreTeam) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.teamRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetAll implements user.ServiceInterface
func (service *teamService) GetAll() (data []team.CoreTeam, err error) {
	data, err = service.teamRepository.GetAll()
	return

}

func (service *teamService) GetById(id int) (data team.CoreTeam, err error) {
	data, errGet := service.teamRepository.GetById(id)
	if errGet != nil {
		return data, errors.New("failed get team by id data, error query")
	}
	return data, nil
}

func (service *teamService) UpdateTeam(dataCore team.CoreTeam, id int) (err error) {
	errUpdate := service.teamRepository.UpdateTeam(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}

func (service *teamService) DeleteTeam(id int) (err error) {
	_, errDel := service.teamRepository.DeleteTeam(id)
	if errDel != nil {
		return errors.New("failed delete team, error query")
	}
	return nil
}
