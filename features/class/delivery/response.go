package delivery

import (
	"immersive-dashboard/features/class"
)

type ClassResponse struct {
	ID           uint         `json:"id"`
	Class_Name   string       `json:"class_name"`
	Users        UserResponse `json:"users"`
	Started_Date string       `json:"started_date"`
	End_Date     string       `json:"end_date"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Full_Name string `json:"nama_user"`
}

func fromCore(dataCore class.CoreClass) ClassResponse {
	return ClassResponse{
		ID:         dataCore.ID,
		Class_Name: dataCore.Class_Name,
		Users: UserResponse{
			ID:        dataCore.Users.ID,
			Full_Name: dataCore.Users.Full_Name,
		},
		Started_Date: dataCore.Started_Date,
		End_Date:     dataCore.End_Date,
	}
}

func fromCoreList(dataCore []class.CoreClass) []ClassResponse {
	var dataResponse []ClassResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
