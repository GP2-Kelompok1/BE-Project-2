package auth

import (
	"gorm.io/gorm"
)

type CoreUser struct {
	gorm.Model
	Full_Name  string
	Email      string
	Password   string
	Teams      []CoreTeam
	Role       string
	Status     string
	Permission string
}

type CoreTeam struct {
	gorm.Model
	Team_Name string
}

type ServiceInterface interface {
	Login(email, password string) (token string, err error)
}

type RepositoryInterface interface {
	FindUser(email, password string) (token string, err error)
}
