package repository

import (
	"immersive-dashboard/features/auth"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint
	Full_Name  string
	Email      string `gorm:"unique"`
	Password   string
	TeamID     uint
	Role       string
	Status     string
	Permission string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

//DTO

func (data User) ToCore() auth.CoreUser {
	return auth.CoreUser{
		ID:         data.ID,
		Full_Name:  data.Full_Name,
		Email:      data.Email,
		Password:   data.Password,
		TeamID:     data.TeamID,
		Role:       data.Role,
		Status:     data.Status,
		Permission: data.Permission,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		DeletedAt:  data.DeletedAt,
	}
}
