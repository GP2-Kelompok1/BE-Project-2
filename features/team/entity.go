package team

type CoreTeam struct {
	ID        uint
	Team_Name string
}

type ServiceInterface interface {
	GetAll() (data []CoreTeam, err error)
	Create(input CoreTeam) (err error)
	GetById(id int) (data CoreTeam, err error)
	UpdateTeam(input CoreTeam, id int) (err error)
	DeleteTeam(id int) (err error)
}

type RepositoryInterface interface {
	GetAll() (data []CoreTeam, err error)
	Create(input CoreTeam) (row int, err error)
	GetById(id int) (data CoreTeam, err error)
	UpdateTeam(input CoreTeam, id int) (err error)
	DeleteTeam(id int) (row int, err error)
}
