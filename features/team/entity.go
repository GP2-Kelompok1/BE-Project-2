package team

type CoreTeam struct {
	ID        uint
	Team_Name string
}

type ServiceInterface interface {
	GetAll() (data []CoreTeam, err error)
	Create(input CoreTeam) (err error)
	GetById(id int) (data CoreTeam, err error)
	// Update(data CoreUser, id int) (string, error)
	// Delete(id int) (string, error)
}

type RepositoryInterface interface {
	GetAll() (data []CoreTeam, err error)
	Create(input CoreTeam) (row int, err error)
	GetById(id int) (data CoreTeam, err error)
	// Update(data CoreUser, id int) (string, error)
	// Delete(id int) (string, error)
}
