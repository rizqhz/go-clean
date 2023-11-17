package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type DatabaseConfig struct {
	host string
	port int
	user string
	pass string
	name string
}

func (c *DatabaseConfig) MySqlConnectStr() string {
	format := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf(format, c.user, c.pass, c.host, c.port, c.name)
}

func (c *DatabaseConfig) PgSqlConnectStr() string {
	format := "host=%s user=%s password=%s dbname=%s port=%d"
	return fmt.Sprintf(format, c.host, c.user, c.pass, c.name, c.port)
}

func NewDatabaseConfig() *DatabaseConfig {
	env, err := NewDatabaseEnv()
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return &DatabaseConfig{
		host: env["DB_HOST"].(string),
		port: env["DB_PORT"].(int),
		user: env["DB_USER"].(string),
		pass: env["DB_PASS"].(string),
		name: env["DB_NAME"].(string),
	}
}
