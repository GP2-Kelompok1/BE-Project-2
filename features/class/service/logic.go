package service

import (
	"errors"
	"immersive-dashboard/features/class"
	"immersive-dashboard/features/user"

	"github.com/go-playground/validator/v10"
)

type classService struct {
	classRepository class.RepositoryInterface
	validate        *validator.Validate
}

func New(repo class.RepositoryInterface) class.ServiceInterface {
	return &classService{
		classRepository: repo,
		validate:        validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *classService) Create(input class.CoreClass) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.classRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetAll implements user.ServiceInterface
func (service *classService) GetAll() (data []user.CoreUser, err error) {
	data, err = service.classRepository.GetAll()
	return

}
