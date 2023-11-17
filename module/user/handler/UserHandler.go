package handler

import "github.com/labstack/echo/v4"

type UserHandler interface {
	Login() echo.HandlerFunc
	Index() echo.HandlerFunc
	Store() echo.HandlerFunc
}
