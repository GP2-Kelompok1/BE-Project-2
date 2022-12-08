package delivery

import (
	"immersive-dashboard/features/user"
)

type UserRequest struct {
	Full_Name  string `json:"full_name" form:"full_name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	TeamID     uint   `json:"team_id" form:"team_id"`
	Role       string `json:"role" form:"role"`
	Status     string `json:"status" form:"status"`
	Permission string `json:"permission" form:"permission"`
}

func toCore(data UserRequest) user.CoreUser {
	return user.CoreUser{
		Full_Name: data.Full_Name,
		Email:     data.Email,
		Password:  data.Password,
		Team: user.CoreTeam{
			ID: data.TeamID,
		},
		Role:       data.Role,
		Status:     data.Status,
		Permission: data.Permission,
	}
}
