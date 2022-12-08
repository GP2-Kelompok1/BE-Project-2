package delivery

import (
	"immersive-dashboard/features/feedback"
)

type FeedbackResponse struct {
	ID             uint           `json:"id"`
	Mentees        MenteeResponse `json:"mentees"`
	Users          UserResponse   `json:"users"`
	Description    string         `json:"description"`
	Mentee_Status  string         `json:"mentee_status"`
	Changed_Status string         `json:"changed_status"`
}
type MenteeResponse struct {
	ID          uint   `json:"id"`
	Mentee_Name string `json:"mentee_name"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Full_Name string `json:"full_name"`
}

func fromCore(dataCore feedback.CoreFeedback) FeedbackResponse {
	return FeedbackResponse{
		ID: dataCore.ID,
		Mentees: MenteeResponse{
			ID:          dataCore.Mentees.ID,
			Mentee_Name: dataCore.Mentees.Mentee_Name,
		},
		Users: UserResponse{
			ID:        dataCore.Users.ID,
			Full_Name: dataCore.Users.Full_Name,
		},
		Description:    dataCore.Description,
		Mentee_Status:  dataCore.Mentee_Status,
		Changed_Status: dataCore.Changed_Status,
	}
}

func fromCoreList(dataCore []feedback.CoreFeedback) []FeedbackResponse {
	var dataResponse []FeedbackResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
