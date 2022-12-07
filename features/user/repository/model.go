package repository

import (
	_user "immersive-dashboard/features/user"

	"gorm.io/gorm"
)

// struct gorm model
type Team struct {
	gorm.Model
	Team_Name string
	Users     []User
}
type User struct {
	gorm.Model
	Full_Name  string
	Email      string
	Password   string
	TeamID     uint // <<tidak ada underscore (default setting)
	Role       string
	Status     string
	Permission string
	Teams      Team
}

// mengubah struct core ke struct model gorm
func fromCore(dataCore _user.CoreUser) User {
	userGorm := User{
		Full_Name: dataCore.Full_Name,
		Email:     dataCore.Email,
		Password:  dataCore.Password,
		TeamID:    dataCore.Teams.ID, // <<< lihat ini
		Role:      dataCore.Role,
	}
	return userGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *User) toCore() _user.CoreUser {
	return _user.CoreUser{
		ID:        dataModel.ID,
		Full_Name: dataModel.Full_Name,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Role:      dataModel.Role,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []User) []_user.CoreUser {
	var dataCore []_user.CoreUser
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
