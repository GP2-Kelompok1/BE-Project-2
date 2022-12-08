package class

type CoreClass struct {
	ID           uint
	Class_Name   string
	Users        CoreUser
	Started_Date string
	End_Date     string
}

type CoreUser struct {
	ID         uint
	Full_Name  string
	Email      string
	Password   string
	TeamID     uint
	Role       string
	Status     string
	Permission string
}

type ServiceInterface interface {
	GetAll() (data []CoreClass, err error)
	Create(input CoreClass) (err error)
	GetById(id int) (data CoreClass, err error)
	// Update(data CoreUser, id int) (string, error)
	// Delete(id int) (string, error)
}

type RepositoryInterface interface {
	GetAll() (data []CoreClass, err error)
	Create(input CoreClass) (row int, err error)
	GetById(id int) (data CoreClass, err error)
	// Update(data CoreUser, id int) (string, error)
	// Delete(id int) (string, error)
}
