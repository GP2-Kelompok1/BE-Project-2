package service

import (
	"errors"
	"immersive-dashboard/features/class"

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
func (service *classService) GetAll() (data []class.CoreClass, err error) {
	data, err = service.classRepository.GetAll()
	return

}

func (service *classService) GetById(id int) (data class.CoreClass, err error) {
	data, errGet := service.classRepository.GetById(id)
	if errGet != nil {
		return data, errors.New("failed get class by id data, error query")
	}
	return data, nil
}

func (service *classService) UpdateClass(dataCore class.CoreClass, id int) (err error) {
	errUpdate := service.classRepository.UpdateClass(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}

func (service *classService) DeleteClass(id int) (err error) {
	_, errDel := service.classRepository.DeleteClass(id)
	if errDel != nil {
		return errors.New("failed delete class, error query")
	}
	return nil
}
