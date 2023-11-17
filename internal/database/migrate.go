package database

import (
	"github.com/rizghz/clean/module/user/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}
