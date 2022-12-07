package delivery

import (
	"immersive-dashboard/features/team"
)

type TeamRequest struct {
	Team_Name string `json:"team_name" form:"team_name"`
}

func toCore(data TeamRequest) team.CoreTeam {
	return team.CoreTeam{}
}
