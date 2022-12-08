package delivery

import (
	"immersive-dashboard/features/mentee"
)

type MenteeRequest struct {
	Mentee_Name           string `json:"mentee_name"`
	ClassID               uint   `json:"class_id"`
	Status                string `json:"status"`
	Category              string `json:"category"`
	Gender                string `json:"gender"`
	Current_Address       string `json:"current_address"`
	Home_Address          string `json:"home_address"`
	Email                 string `json:"email"`
	Telegram              string `json:"telegram"`
	Phone                 string `json:"phone"`
	Emergency_Name        string `json:"emergency_name"`
	Emergency_Phone       string `json:"emergency-phone"`
	Emergency_Status      string `json:"emergency_status"`
	Education_Type        string `json:"education_type"`
	Education_Major       string `json:"education_major"`
	Education_Graduate    string `json:"education_graduate"`
	Education_Institution string `json:"education_institution"`
}

func toCore(data MenteeRequest) mentee.CoreMentee {
	return mentee.CoreMentee{
		Mentee_Name: data.Mentee_Name,
		Classes: mentee.CoreClass{
			ID: data.ClassID,
		},
		Status:                data.Status,
		Category:              data.Category,
		Gender:                data.Gender,
		Current_Address:       data.Current_Address,
		Home_Address:          data.Home_Address,
		Email:                 data.Email,
		Telegram:              data.Telegram,
		Phone:                 data.Phone,
		Emergency_Name:        data.Emergency_Name,
		Emergency_Phone:       data.Emergency_Phone,
		Emergency_Status:      data.Emergency_Status,
		Education_Type:        data.Education_Type,
		Education_Major:       data.Education_Major,
		Education_Graduate:    data.Education_Graduate,
		Education_Institution: data.Education_Institution,
	}
}
