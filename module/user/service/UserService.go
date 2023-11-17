package service

import "github.com/rizghz/clean/module/user/transfer"

type UserService interface {
	UserLogin(req *transfer.LoginRequestBody) (*transfer.LoginResponseBody, error)
	GetAllUsers() []*transfer.UserResponseBody
	CreateUser(data *transfer.UserRequestBody) (bool, error)
}
