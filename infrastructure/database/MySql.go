package driver

import (
	"github.com/rizghz/clean/infrastructure/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlDriver struct {
	DB *gorm.DB
}

func NewMySqlDriver() *MySqlDriver {
	dsn := config.NewDatabaseConfig().MySqlConnectStr()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		TranslateError:         true,
	})
	if err != nil {
		logrus.Fatal(db.Error)
		return nil
	}
	return &MySqlDriver{db}
}
