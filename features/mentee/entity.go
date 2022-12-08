package mentee

type CoreMentee struct {
	ID                    uint
	Mentee_Name           string `validate:"required"`
	Class                 CoreClass
	Status                string `validate:"required"`
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
}
type CoreClass struct {
	ID           uint
	Class_Name   string
	Started_Date string
	End_Date     string
}

type ServiceInterface interface {
	GetAll() (data []CoreMentee, err error)
	Create(input CoreMentee) (err error)
	GetById(id int) (data CoreMentee, err error)
	UpdateMentee(input CoreMentee, id int) (err error)
	DeleteMentee(id int) (err error)
}

type RepositoryInterface interface {
	GetAll() (data []CoreMentee, err error)
	Create(input CoreMentee) (row int, err error)
	GetById(id int) (data CoreMentee, err error)
	UpdateMentee(input CoreMentee, id int) (err error)
	DeleteMentee(id int) (row int, err error)
}
