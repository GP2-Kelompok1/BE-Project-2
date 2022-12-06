package repository

import (
	"immersive-dashboard/features/auth"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Full_Name  string
	Email      string `gorm:"unique"`
	Password   string
	Teams      []Team
	Role       string
	Status     string
	Permission string
}

type Team struct {
	gorm.Model
	Team_Name string
}

//DTO

func (data User) toCore() auth.CoreUser {
	return auth.CoreUser{
		Model:      gorm.Model{},
		Full_Name:  data.Full_Name,
		Email:      data.Email,
		Password:   data.Password,
		Teams:      []auth.CoreTeam{},
		Role:       data.Role,
		Status:     data.Status,
		Permission: data.Permission,
	}
}
