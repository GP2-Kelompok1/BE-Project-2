package auth

import (
	"time"

	"gorm.io/gorm"
)

type CoreUser struct {
	ID         uint
	Full_Name  string
	Email      string
	Password   string
	TeamID     uint
	Role       string
	Status     string
	Permission string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

type ServiceInterface interface {
	Login(email, password string) (token string, err error)
}

type RepositoryInterface interface {
	FindUser(email, password string) (token string, err error)
}
