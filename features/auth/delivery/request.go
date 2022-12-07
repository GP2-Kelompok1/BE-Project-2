package delivery

import (
	"immersive-dashboard/features/auth"
)

type UserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(userReq UserRequest) auth.CoreUser {
	userCore := auth.CoreUser{
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	return userCore
}
