package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizghz/clean/internal/utils"
	"github.com/rizghz/clean/module/user/service"
	"github.com/rizghz/clean/module/user/transfer"
)

type UserHttpHandler struct {
	srv service.UserService
}

func NewUserHttpHandler(srv service.UserService) UserHandler {
	return &UserHttpHandler{srv}
}

func (h *UserHttpHandler) Login() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := transfer.LoginRequestBody{}
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, &utils.Response[any]{
				Message: "invalid user data payload",
			})
		}
		res, err := h.srv.UserLogin(&req)
		if err != nil {
			return ctx.JSON(http.StatusOK, utils.Response[any]{
				Message: err.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, res)
	}
}

func (h *UserHttpHandler) Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		data := h.srv.GetAllUsers()
		return ctx.JSON(http.StatusOK, data)
	}
}

func (h *UserHttpHandler) Store() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := &transfer.UserRequestBody{}
		if err := ctx.Bind(req); err != nil {
			return ctx.JSON(http.StatusBadRequest, utils.Response[any]{
				Message: "invalid user data payload",
			})
		}
		if valid, err := h.srv.CreateUser(req); !valid {
			return ctx.JSON(http.StatusInternalServerError, utils.Response[any]{
				Message: err.Error(),
			})
		}
		return ctx.JSON(http.StatusCreated, utils.Response[any]{
			Message: "success",
		})
	}
}
