package repository

import (
	_mentee "immersive-dashboard/features/mentee"

	"gorm.io/gorm"
)

// struct gorm model
type Mentee struct {
	gorm.Model
	Mentee_Name           string
	ClassID               uint
	Status                string
	Category              string
	Gender                string
	Current_Address       string
	Home_Address          string
	Email                 string
	Telegram              string
	Phone                 string
	Emergency_Name        string
	Emergency_Phone       string
	Emergency_Status      string
	Education_Type        string
	Education_Major       string
	Education_Graduate    string
	Education_Institution string
	Classes               Class
	Feedbacks             []Feedback
}
type Class struct {
	gorm.Model
	Class_Name   string
	Started_Date string
	Mentees      []Mentee
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
func fromCore(dataCore _mentee.CoreMentee) Mentee {
	userGorm := Mentee{

		Mentee_Name:           dataCore.Mentee_Name,
		ClassID:               dataCore.Classes.ID,
		Status:                dataCore.Status,
		Category:              dataCore.Category,
		Gender:                dataCore.Gender,
		Current_Address:       dataCore.Current_Address,
		Home_Address:          dataCore.Home_Address,
		Email:                 dataCore.Email,
		Telegram:              dataCore.Telegram,
		Phone:                 dataCore.Phone,
		Emergency_Name:        dataCore.Emergency_Name,
		Emergency_Phone:       dataCore.Emergency_Phone,
		Emergency_Status:      dataCore.Emergency_Status,
		Education_Type:        dataCore.Education_Type,
		Education_Major:       dataCore.Education_Major,
		Education_Graduate:    dataCore.Education_Graduate,
		Education_Institution: dataCore.Education_Institution,
	}
	return userGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *Mentee) toCore() _mentee.CoreMentee {
	return _mentee.CoreMentee{
		ID:          dataModel.ID,
		Mentee_Name: dataModel.Mentee_Name,
		Classes: _mentee.CoreClass{
			ID:         dataModel.ClassID,
			Class_Name: dataModel.Classes.Class_Name,
		},
		Status:                dataModel.Status,
		Category:              dataModel.Category,
		Gender:                dataModel.Gender,
		Current_Address:       dataModel.Current_Address,
		Home_Address:          dataModel.Home_Address,
		Email:                 dataModel.Email,
		Telegram:              dataModel.Telegram,
		Phone:                 dataModel.Phone,
		Emergency_Name:        dataModel.Emergency_Name,
		Emergency_Phone:       dataModel.Emergency_Phone,
		Emergency_Status:      dataModel.Emergency_Status,
		Education_Type:        dataModel.Education_Type,
		Education_Major:       dataModel.Education_Major,
		Education_Graduate:    dataModel.Education_Graduate,
		Education_Institution: dataModel.Education_Institution,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Mentee) []_mentee.CoreMentee {
	var dataCore []_mentee.CoreMentee
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

// type Pekerja struct {
// 	gorm.Model
// 	Name      string
// 	CompanyID uint
// 	Company   Company
//   }

//   type Company struct {
// 	ID   uint
// 	Name string
//   }

//   pekerja1, 123
//   pekerja2, 123
//   pekerja3, 234

//   type User struct {
// 	gorm.Model
// 	CreditCard CreditCard
//   }

//   type CreditCard struct {
// 	gorm.Model
// 	Number string
// 	UserID uint
//   }

//   user1, cc1
//   user2, cc2

//   type User struct {
// 	gorm.Model
// 	Name		string
// 	Laptop		Laptop

//   }

//   type Laptop struct {
// 	gorm.Model
// 	NameLaptop 	string
// 	UserID		uint

//   }
// user1, 123, a
// user2, 234, a

// manggil user
// id
// NameLaptop
// id Laptop
// nama Laptop
// user ID

// 1
// deva
// 17
// lenovo

// 17
// lenovo
// 1
// deva

// 2
// budi
// 27
// acer
