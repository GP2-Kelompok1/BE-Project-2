package delivery

import "immersive-dashboard/features/mentee"

type MenteeResponse struct {
	ID                    uint          `json:"id"`
	Mentee_Name           string        `json:"mentee_name"`
	Classes               ClassResponse `json:"classes"`
	Status                string        `json:"status"`
	Category              string        `json:"category"`
	Gender                string        `json:"gender"`
	Current_Address       string        `json:"current_address"`
	Home_Address          string        `json:"home_address"`
	Email                 string        `json:"email"`
	Telegram              string        `json:"telegram"`
	Phone                 string        `json:"phone"`
	Emergency_Name        string        `json:"emergency_name"`
	Emergency_Phone       string        `json:"emergency-phone"`
	Emergency_Status      string        `json:"emergency_status"`
	Education_Type        string        `json:"education_type"`
	Education_Major       string        `json:"education_major"`
	Education_Graduate    string        `json:"education_graduate"`
	Education_Institution string        `json:"education_institution"`
}
type ClassResponse struct {
	ID         uint   `json:"id"`
	Class_Name string `json:"class_name"`
}

func fromCore(dataCore mentee.CoreMentee) MenteeResponse {
	return MenteeResponse{
		ID:          dataCore.ID,
		Mentee_Name: dataCore.Mentee_Name,
		Classes: ClassResponse{
			ID:         dataCore.Class.ID,
			Class_Name: dataCore.Class.Class_Name,
		},
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
}
func fromCoreList(dataCore []mentee.CoreMentee) []MenteeResponse {
	var dataResponse []MenteeResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
