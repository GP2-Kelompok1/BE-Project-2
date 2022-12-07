package delivery

import (
	"immersive-dashboard/features/class"
)

type ClassRequest struct {
	Class_Name   string `json:"class_name" form:"class_name"`
	UserID       uint   `json:"user_id" form:"user_id"`
	Started_Date string `json:"started_date" form:"started_date"`
}

func toCore(data ClassRequest) class.CoreClass {
	return class.CoreClass{
		Class_Name: data.Class_Name,
		Users: class.CoreUser{ // <<< lihat ini
			ID: data.UserID,
		},
		Started_Date: data.Started_Date,
	}
}
