package service

import (
	"errors"
	"immersive-dashboard/features/user"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userRepository user.RepositoryInterface
	validate       *validator.Validate
}

func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *userService) Create(input user.CoreUser) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.userRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetAll implements user.ServiceInterface
func (service *userService) GetAll() (data []user.CoreUser, err error) {
	data, err = service.userRepository.GetAll()
	return

}

func (service *userService) GetById(id int) (data user.CoreUser, err error) {
	data, errGet := service.userRepository.GetById(id)
	if errGet != nil {
		return data, errors.New("failed get user by id data, error query")
	}
	return data, nil
}

func (service *userService) UpdateUser(dataCore user.CoreUser, id int) (err error) {
	errUpdate := service.userRepository.UpdateUser(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}

func (service *userService) DeleteUser(id int) (err error) {
	_, errDel := service.userRepository.DeleteUser(id)
	if errDel != nil {
		return errors.New("failed delete user, error query")
	}
	return nil
}
