package output

import (
	"github.com/architecture-template/echo-ddd/domain/model"
)

type User struct {
	UserKey  string `json:"user_key"`
	UserName string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Message  string `json:"message"`
}

func ToUser(u *model.User) *User {
	return &User{
		UserKey:  u.UserKey,
		UserName: u.UserName,
		Email:    u.Email,
		Token:    u.Token,
		Message:  "user completed",
	}
}
