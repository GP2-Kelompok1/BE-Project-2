package service

import (
	"errors"
	"immersive-dashboard/features/mentee"

	"github.com/go-playground/validator/v10"
)

type menteeService struct {
	menteeRepository mentee.RepositoryInterface
	validate         *validator.Validate
}

func New(repo mentee.RepositoryInterface) mentee.ServiceInterface {
	return &menteeService{
		menteeRepository: repo,
		validate:         validator.New(),
	}
}

func (service *menteeService) GetAll() (data []mentee.CoreMentee, err error) {
	data, err = service.menteeRepository.GetAll()
	return

}

// Create implements user.ServiceInterface
func (service *menteeService) Create(input mentee.CoreMentee) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.menteeRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

func (service *menteeService) GetById(id int) (data mentee.CoreMentee, err error) {
	data, errGet := service.menteeRepository.GetById(id)
	if errGet != nil {
		return data, errors.New("failed get user by id data, error query")
	}
	return data, nil
}

func (service *menteeService) UpdateMentee(dataCore mentee.CoreMentee, id int) (err error) {
	errUpdate := service.menteeRepository.UpdateMentee(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}

func (service *menteeService) DeleteMentee(id int) (err error) {
	_, errDel := service.menteeRepository.DeleteMentee(id)
	if errDel != nil {
		return errors.New("failed delete mentee, error query")
	}
	return nil
}
