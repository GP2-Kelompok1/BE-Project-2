package service

import (
	"errors"
	"immersive-dashboard/features/feedback"

	"github.com/go-playground/validator/v10"
)

type feedbackService struct {
	feedbackRepository feedback.RepositoryInterface
	validate           *validator.Validate
}

func New(repo feedback.RepositoryInterface) feedback.ServiceInterface {
	return &feedbackService{
		feedbackRepository: repo,
		validate:           validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *feedbackService) Create(input feedback.CoreFeedback) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.feedbackRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetAll implements user.ServiceInterface
func (service *feedbackService) GetAll() (data []feedback.CoreFeedback, err error) {
	data, err = service.feedbackRepository.GetAll()
	return

}
