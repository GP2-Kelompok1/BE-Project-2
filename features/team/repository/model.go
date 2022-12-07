package repository

import (
	_team "immersive-dashboard/features/team"

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
func fromCore(dataCore _team.CoreTeam) Team {
	userGorm := Team{
		Team_Name: dataCore.Team_Name,
	}
	return userGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *Team) toCore() _team.CoreTeam {
	return _team.CoreTeam{
		Team_Name: dataModel.Team_Name,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Team) []_team.CoreTeam {
	var dataCore []_team.CoreTeam
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
