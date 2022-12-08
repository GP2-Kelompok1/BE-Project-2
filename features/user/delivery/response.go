package delivery

import "immersive-dashboard/features/user"

type UserResponse struct {
	ID         uint         `json:"id"`
	Full_Name  string       `json:"full_name"`
	Email      string       `json:"email"`
	Teams      TeamResponse `json:"teams"`
	Role       string       `json:"role"`
	Status     string       `json:"status"`
	Permission string       `json:"permission"`
}

type TeamResponse struct {
	ID        uint   `json:"id"`
	Team_Name string `json:"team_name"`
}

func fromCore(dataCore user.CoreUser) UserResponse {
	return UserResponse{
		ID:        dataCore.ID,
		Full_Name: dataCore.Full_Name,
		Email:     dataCore.Email,
		Teams: TeamResponse{
			ID:        dataCore.Teams.ID,
			Team_Name: dataCore.Teams.Team_Name,
		},
		Role:       dataCore.Role,
		Status:     dataCore.Status,
		Permission: dataCore.Permission,
	}
}

// func toCoreTeam(dataCore user.CoreTeam) TeamResponse {
// 	return TeamResponse{
// 		ID:        dataCore.ID,
// 		Team_Name: dataCore.Team_Name,
// 	}
// }

func fromCoreList(dataCore []user.CoreUser) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
