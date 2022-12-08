package user

type CoreUser struct {
	ID         uint
	Full_Name  string
	Email      string `validate:"required,email"`
	Password   string `validate:"required"`
	Teams      CoreTeam
	Role       string
	Status     string
	Permission string
}

type CoreTeam struct {
	ID        uint
	Team_Name string
}

type ServiceInterface interface {
	GetAll() (data []CoreUser, err error)
	Create(input CoreUser) (err error)
	GetById(id int) (data CoreUser, err error)
	UpdateUser(input CoreUser, id int) (err error)
	// Delete(id int) (string, error)
}

type RepositoryInterface interface {
	GetAll() (data []CoreUser, err error)
	Create(input CoreUser) (row int, err error)
	GetById(id int) (data CoreUser, err error)
	UpdateUser(input CoreUser, id int) (err error)
	// Delete(id int) (string, error)
}
