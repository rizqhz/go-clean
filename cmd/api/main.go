package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rizghz/clean/infrastructure/config"
	driver "github.com/rizghz/clean/infrastructure/database"
	"github.com/rizghz/clean/internal/database"
	"github.com/rizghz/clean/internal/middleware"
	"github.com/rizghz/clean/internal/route"
	"github.com/rizghz/clean/module/user/handler"
	"github.com/rizghz/clean/module/user/repository"
	"github.com/rizghz/clean/module/user/service"
)

func main() {
	key, _ := config.NewJwtEnv()
	signkey := key["SECRET_KEY"].(string)

	d := driver.NewMySqlDriver()
	r := repository.NewUserSqlRepository(d)
	s := service.NewUserServiceImpl(r, middleware.NewJwt(signkey, ""))
	h := handler.NewUserHttpHandler(s)

	database.Migrate(d.DB)

	e := echo.New()
	route.UserRoute(e, h)
	e.Logger.Fatal(e.Start(":8008").Error())
}
