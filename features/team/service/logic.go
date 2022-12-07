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
