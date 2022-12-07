package delivery

import "immersive-dashboard/features/user"

type UserResponse struct {
	ID        uint   `json:"id" form:"id"`
	Full_Name string `json:"full_name" form:"full_name"`
	Email     string `json:"email" form:"email"`
}

func FromCore(data user.CoreUser) UserResponse {
	return UserResponse{
		ID:        data.ID,
		Full_Name: data.Full_Name,
		Email:     data.Email,
	}
}
