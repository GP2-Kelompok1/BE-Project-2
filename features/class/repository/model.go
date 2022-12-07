package repository

import (
	_class "immersive-dashboard/features/class"

	"gorm.io/gorm"
)

// struct gorm model
type Class struct {
	gorm.Model
	Class_Name   string
	UserID       uint
	Started_Date string
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
}

// mengubah struct core ke struct model gorm
func fromCore(dataCore _class.CoreClass) Class {
	classGorm := Class{
		Class_Name:   dataCore.Class_Name,
		UserID:       dataCore.Users.ID,
		Started_Date: dataCore.Started_Date,
	}
	return classGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *Class) toCore() _class.CoreClass {
	return _class.CoreClass{
		ID:           dataModel.ID,
		Class_Name:   dataModel.Class_Name,
		Users:        dataModel.UserID.ID,
		Started_Date: dataModel.Started_Date,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Class) []_class.CoreClass {
	var dataCore []_class.CoreClass
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
