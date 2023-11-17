package service

import (
	"testing"

	"github.com/rizghz/clean/mocks"
	"github.com/rizghz/clean/module/user/entity"
	"github.com/rizghz/clean/module/user/transfer"
	"github.com/stretchr/testify/assert"
)

func TestUserLogin(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	token := mocks.NewJwtTokenInterface(t)
	serv := NewUserServiceImpl(repo, token)

	t.Run("Valid User Login", func(t *testing.T) {
		login := &transfer.LoginRequestBody{
			Email:    "a@mail.com",
			Password: "A12345678",
		}
		req := &transfer.UserRequestBody{
			Email:    login.Email,
			Password: login.Password,
		}
		ent := &entity.User{
			Name:     "User A",
			Email:    req.Email,
			Password: req.Password,
		}
		repo.On("Find", req).Return(ent).Once()
		result := "jwttoken.f4i1k6"
		token.On("GenerateToken", ent.ID).Return(&result).Once()
		res, err := serv.UserLogin(login)
		assert.Nil(t, err)
		assert.Equal(t, res.Token, result)
	})

	t.Run("Invalid User Credential", func(t *testing.T) {
		login := &transfer.LoginRequestBody{
			Email:    "amail.com",
			Password: "A123",
		}
		res, err := serv.UserLogin(login)
		assert.EqualError(t, err, "invalid user credential")
		assert.Nil(t, res)
	})

	t.Run("User Not Found", func(t *testing.T) {
		login := &transfer.LoginRequestBody{
			Email:    "tidakada@mail.com",
			Password: "tidakada4312",
		}
		req := &transfer.UserRequestBody{
			Email:    login.Email,
			Password: login.Password,
		}
		repo.On("Find", req).Return(nil).Once()
		res, err := serv.UserLogin(login)
		assert.EqualError(t, err, "user not found")
		assert.Nil(t, res)
	})
}

func TestGetAllUsers(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	token := mocks.NewJwtTokenInterface(t)
	serv := NewUserServiceImpl(repo, token)

	t.Run("Valid Get All Users", func(t *testing.T) {
		data := []*entity.User{
			{Name: "", Email: "", Password: ""},
			{Name: "", Email: "", Password: ""},
			{Name: "", Email: "", Password: ""},
		}
		repo.On("Get").Return(data).Once()
		res := serv.GetAllUsers()
		assert.NotEmpty(t, res)
	})

	t.Run("Invalid Get All Users", func(t *testing.T) {
		repo.On("Get").Return(nil).Once()
		res := serv.GetAllUsers()
		assert.Empty(t, res)
	})
}

func TestCreateUser(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	token := mocks.NewJwtTokenInterface(t)
	serv := NewUserServiceImpl(repo, token)

	t.Run("Valid Create User", func(t *testing.T) {
		req := &transfer.UserRequestBody{
			Name:     "User A",
			Email:    "a@mail.com",
			Password: "A12345678",
		}
		ent := &entity.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		}
		repo.On("Create", ent).Return(true, ent).Once()
		check, err := serv.CreateUser(req)
		assert.True(t, check)
		assert.Nil(t, err)
	})

	t.Run("Invalid User Credential", func(t *testing.T) {
		req := &transfer.UserRequestBody{
			Name:     "user a",
			Email:    "amail.com",
			Password: "A123",
		}
		check, err := serv.CreateUser(req)
		assert.False(t, check)
		assert.EqualError(t, err, "invalid user credential")
	})

	t.Run("Failed Create New User", func(t *testing.T) {
		req := &transfer.UserRequestBody{
			Name:     "user a",
			Email:    "a@mail.com",
			Password: "A12345678",
		}
		ent := &entity.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		}
		repo.On("Create", ent).Return(false, nil).Once()
		check, err := serv.CreateUser(req)
		assert.False(t, check)
		assert.EqualError(t, err, "failed create new user")
	})
}
