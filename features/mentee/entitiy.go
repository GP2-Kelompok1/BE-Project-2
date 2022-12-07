package mentee

import "time"

type Core struct {
	ID                    uint
	Name                  string `validate:"required"`
	ID_class              uint
	Status                string `validate:"required"`
	Kategori              string `validate:"required"`
	Gender                string `validate:"required"`
	Current_Address       string `validate:"required"`
	Home_Address          string `validate:"required"`
	Email                 string `validate:"required,email"`
	Telegram              string `validate:"required"`
	Phone                 string `validate:"required"`
	Emergency_Name        string `validate:"required"`
	Emergency_Phone       string `validate:"required"`
	Emergency_Status      string `validate:"required"`
	Education_Type        string `validate:"required"`
	Education_Major       string `validate:"required"`
	Education_Graduate    string `validate:"required"`
	Education_Institution string `validate:"required"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) (err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) (row int, err error)
	GetById(id int) (data Core, err error)
}
