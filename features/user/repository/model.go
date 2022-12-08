package repository

import (
	_user "immersive-dashboard/features/user"

	"gorm.io/gorm"
)

// struct gorm model

type User struct {
	gorm.Model
	Full_Name  string
	Email      string
	Password   string
	TeamID     uint // <<tidak ada underscore (default setting)
	Role       string
	Status     string
	Permission string
	Classes    []Class
	Feedbacks  []Feedback
	Teams      Team
}
type Team struct {
	gorm.Model
	Team_Name string
}
type Class struct {
	gorm.Model
	Class_Name   string
	UserID       uint
	Started_Date string
	End_Date     string
}

type Feedback struct {
	gorm.Model
	MenteeID       uint
	UserID         uint
	Description    string
	Mentee_Status  string
	Changed_Status string
}

// mengubah struct core ke struct model gorm
func fromCore(dataCore _user.CoreUser) User {
	userGorm := User{
		Full_Name:  dataCore.Full_Name,
		Email:      dataCore.Email,
		Password:   dataCore.Password,
		TeamID:     dataCore.Teams.ID, // <<< lihat ini
		Role:       dataCore.Role,
		Status:     dataCore.Status,
		Permission: dataCore.Permission,
	}
	return userGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *User) toCore() _user.CoreUser {
	return _user.CoreUser{
		ID:        dataModel.ID,
		Full_Name: dataModel.Full_Name,
		Email:     dataModel.Email,
		Teams: _user.CoreTeam{ // <<< lihat ini
			ID:        dataModel.TeamID,
			Team_Name: dataModel.Teams.Team_Name,
		},
		Role:       dataModel.Role,
		Status:     dataModel.Status,
		Permission: dataModel.Permission,
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
