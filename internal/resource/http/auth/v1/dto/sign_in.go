package dto

import dbusermod "github.com/web-rabis/db/user"

type User struct {
	Id         int64                `json:"id"`
	Username   string               `json:"username"`
	Email      string               `json:"email"`
	Name       string               `json:"name"`
	Role       string               `json:"role"`
	Department dbusermod.Department `json:"department"`
}
type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type SignInResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

func NewSignInResponse(token string, user *dbusermod.User) SignInResponse {
	return SignInResponse{
		Token: token,
		User: User{
			Id:         user.Id,
			Username:   user.Username,
			Email:      user.Email,
			Name:       user.Name,
			Department: user.Department,
			Role:       "admin",
		},
	}
}
