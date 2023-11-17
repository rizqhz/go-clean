package service

import (
	"errors"
	"strings"

	"github.com/rizghz/clean/internal/middleware"
	"github.com/rizghz/clean/module/user/entity"
	"github.com/rizghz/clean/module/user/repository"
	"github.com/rizghz/clean/module/user/transfer"
)

type UserServiceImpl struct {
	repo  repository.UserRepository
	token middleware.JwtTokenInterface
}

func NewUserServiceImpl(r repository.UserRepository, t middleware.JwtTokenInterface) UserService {
	return &UserServiceImpl{
		repo:  r,
		token: t,
	}
}

func (srv *UserServiceImpl) UserLogin(req *transfer.LoginRequestBody) (*transfer.LoginResponseBody, error) {
	if !srv.validateEmail(req.Email) && !srv.validatePassword(req.Password) {
		return nil, errors.New("invalid user credential")
	}
	data := &transfer.UserRequestBody{
		Email:    req.Email,
		Password: req.Password,
	}
	if result := srv.repo.Find(data); result != nil {
		token := *srv.token.GenerateToken(result.ID)
		res := &transfer.LoginResponseBody{Token: token}
		return res, nil
	}
	return nil, errors.New("user not found")
}

func (srv *UserServiceImpl) GetAllUsers() []*transfer.UserResponseBody {
	data := make([]*transfer.UserResponseBody, 0)
	for _, v := range srv.repo.Get() {
		data = append(data, &transfer.UserResponseBody{
			UserId:   v.ID,
			Name:     v.Name,
			Email:    v.Email,
			Password: v.Password,
		})
	}
	return data
}

func (srv *UserServiceImpl) CreateUser(data *transfer.UserRequestBody) (bool, error) {
	if !srv.validateEmail(data.Email) && !srv.validatePassword(data.Password) {
		return false, errors.New("invalid user credential")
	}
	user := &entity.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
	if check, _ := srv.repo.Create(user); !check {
		return false, errors.New("failed create new user")
	}
	return true, nil
}

func (srv *UserServiceImpl) validateEmail(email string) bool {
	return strings.Contains(email, "@")
}

func (srv *UserServiceImpl) validatePassword(password string) bool {
	return len(password) > 8
}
