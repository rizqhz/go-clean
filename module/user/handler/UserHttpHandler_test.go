package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rizghz/clean/internal/utils"
	"github.com/rizghz/clean/mocks"
	"github.com/rizghz/clean/module/user/transfer"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	s := mocks.NewUserService(t)
	h := NewUserHttpHandler(s)
	e := echo.New()

	t.Run("Invalid User Data Payload", func(t *testing.T) {
		data := []byte(`falskdjf1kj21lkf`)

		req := httptest.NewRequest(http.MethodGet, "/users/login", bytes.NewReader(data))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		res := &transfer.LoginResponseBody{}

		e.GET("/users/login", h.Login())
		e.ServeHTTP(rec, req)

		json.Unmarshal(rec.Body.Bytes(), &res)

		if assert.NoError(t, nil, h.Login()) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Empty(t, res.Token)
		}
	})

	t.Run("Invalid User Data Payload", func(t *testing.T) {
		data := []byte(`{"email":"tidakada@mail.com", "password":"tidakada12345"}`)

		s.On("UserLogin", &transfer.LoginRequestBody{
			Email:    "tidakada@mail.com",
			Password: "tidakada12345",
		}).Return(nil, errors.New("user not found")).Once()

		req := httptest.NewRequest(http.MethodGet, "/users/login", bytes.NewReader(data))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		res := &transfer.LoginResponseBody{}

		e.GET("/users/login", h.Login())
		e.ServeHTTP(rec, req)

		json.Unmarshal(rec.Body.Bytes(), &res)

		if assert.NoError(t, nil, h.Login()) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("Valid User Login", func(t *testing.T) {
		data := []byte(`{"email":"rizuki@mail.com", "password":"rizuki12345"}`)

		s.On("UserLogin", &transfer.LoginRequestBody{
			Email:    "rizuki@mail.com",
			Password: "rizuki12345",
		}).Return(&transfer.LoginResponseBody{
			Token: "tokenjwt.12j43k2j3423j432",
		}, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/users/login", bytes.NewReader(data))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		res := &transfer.LoginResponseBody{}

		e.GET("/users/login", h.Login())
		e.ServeHTTP(rec, req)

		json.Unmarshal(rec.Body.Bytes(), &res)

		if assert.NoError(t, nil, h.Login()) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotEmpty(t, res.Token)
		}
	})
}

func TestIndex(t *testing.T) {
	s := mocks.NewUserService(t)
	h := NewUserHttpHandler(s)
	e := echo.New()

	t.Run("Valid User Index", func(t *testing.T) {
		s.On("GetAllUsers").Return([]*transfer.UserResponseBody{
			{Name: "User A", Email: "a@mail.com", Password: "a12345678"},
			{Name: "User B", Email: "b@mail.com", Password: "b12345678"},
			{Name: "User C", Email: "c@mail.com", Password: "c12345678"},
		}).Once()

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		rec := httptest.NewRecorder()

		e.GET("/users", h.Index())
		e.ServeHTTP(rec, req)

		res := []*transfer.UserResponseBody{}
		json.Unmarshal(rec.Body.Bytes(), &res)

		if assert.NoError(t, nil, h.Index()) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotNil(t, res)
		}
	})

	t.Run("Invalid User Index", func(t *testing.T) {
		s.On("GetAllUsers").Return(nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		rec := httptest.NewRecorder()

		e.GET("/users", h.Index())
		e.ServeHTTP(rec, req)

		res := []*transfer.UserResponseBody{}
		json.Unmarshal(rec.Body.Bytes(), &res)

		if assert.NoError(t, nil, h.Index()) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Nil(t, res)
		}
	})
}

func TestStore(t *testing.T) {
	s := mocks.NewUserService(t)
	h := NewUserHttpHandler(s)
	e := echo.New()

	t.Run("Invalid User Data Payload", func(t *testing.T) {
		data := []byte(`{"name":"contoh", "email":"contohmail.com", "password":"contoh"}`)

		s.On("CreateUser", &transfer.UserRequestBody{
			Name:     "contoh",
			Email:    "contohmail.com",
			Password: "contoh",
		}).Return(false, errors.New("invalid user credential")).Once()

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(data))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		res := &utils.Response[any]{}

		e.POST("/users", h.Store())
		e.ServeHTTP(rec, req)

		json.Unmarshal(rec.Body.Bytes(), &res)

		if assert.NoError(t, nil, h.Login()) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "invalid user credential", res.Message)
		}
	})

	t.Run("Invalid User Data Payload", func(t *testing.T) {
		data := []byte(`{"email":"contoh@mail.com", "password":"contoh12345"}`)

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(data))
		// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		res := &utils.Response[any]{}

		e.POST("/users", h.Store())
		e.ServeHTTP(rec, req)

		json.Unmarshal(rec.Body.Bytes(), &res)

		if assert.NoError(t, nil, h.Login()) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "invalid user data payload", res.Message)
		}
	})

	t.Run("Valid User Store", func(t *testing.T) {
		data := []byte(`{"name":"Rizqi", "email":"rizuki@mail.com", "password":"rizuki12345"}`)

		s.On("CreateUser", &transfer.UserRequestBody{
			Name:     "Rizqi",
			Email:    "rizuki@mail.com",
			Password: "rizuki12345",
		}).Return(true, nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(data))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		res := &utils.Response[any]{}

		e.POST("/users", h.Store())
		e.ServeHTTP(rec, req)

		json.Unmarshal(rec.Body.Bytes(), &res)

		if assert.NoError(t, nil, h.Login()) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, "success", res.Message)
		}
	})
}
