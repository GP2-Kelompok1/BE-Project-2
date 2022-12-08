package delivery

import (
	"immersive-dashboard/features/feedback"
)

type FeedbackRequest struct {
	MenteeID       uint   `json:"mentee_id" form:"mentee_id"`
	UserID         uint   `json:"user_id" form:"user_id"`
	Description    string `json:"description" form:"description"`
	Mentee_Status  string `json:"mentee_status" form:"mentee_status"`
	Changed_Status string `json:"changed_status" form:"changed_status"`
}

func toCore(data FeedbackRequest) feedback.CoreFeedback {
	return feedback.CoreFeedback{
		Mentees: feedback.CoreMentee{
			ID: data.MenteeID,
		},
		Users: feedback.CoreUser{
			ID: data.UserID,
		},
		Description:    data.Description,
		Mentee_Status:  data.Mentee_Status,
		Changed_Status: data.Changed_Status,
	}
}
