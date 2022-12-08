package delivery

import "immersive-dashboard/features/team"

type TeamResponse struct {
	ID        uint   `json:"id"`
	Team_Name string `json:"team_name"`
}

//ngerubah core ke response
func fromCore(dataCore team.CoreTeam) TeamResponse {
	return TeamResponse{
		ID:        dataCore.ID,
		Team_Name: dataCore.Team_Name,
	}
}

func fromCoreList(dataCore []team.CoreTeam) []TeamResponse {
	var dataResponse []TeamResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
