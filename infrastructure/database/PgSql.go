package driver

import (
	"github.com/rizghz/clean/infrastructure/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgSqlDriver struct {
	DB *gorm.DB
}

func NewPgSqlDriver() *PgSqlDriver {
	dsn := config.NewDatabaseConfig().PgSqlConnectStr()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		TranslateError:         true,
	})
	if err != nil {
		logrus.Fatal(db.Error)
		return nil
	}
	return &PgSqlDriver{db}
}
