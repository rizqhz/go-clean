package route

import (
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/rizghz/clean/infrastructure/config"
	"github.com/rizghz/clean/module/user/handler"
)

var (
	key, _ = config.NewJwtEnv()
)

func UserRoute(e *echo.Echo, h handler.UserHandler) {
	users, jwt := e.Group("/users"), echojwt.JWT([]byte(key["SECRET_KEY"].(string)))
	users.GET("/login", h.Login())
	users.GET("", h.Index(), jwt)
	users.POST("", h.Store(), jwt)
}
