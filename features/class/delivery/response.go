package delivery

import (
	"immersive-dashboard/features/class"
)

type ClassResponse struct {
	ID           uint           `json:"id"`
	Class_Name   string         `json:"class_name"`
	Users        class.CoreUser `json:"users"`
	Started_Date string         `json:"started_date"`
}

func fromCore(dataCore class.CoreClass) ClassResponse {
	return ClassResponse{
		ID:           dataCore.ID,
		Class_Name:   dataCore.Class_Name,
		Users:        dataCore.Users,
		Started_Date: dataCore.Started_Date,
	}
}

func fromCoreList(dataCore []class.CoreClass) []ClassResponse {
	var dataResponse []ClassResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
