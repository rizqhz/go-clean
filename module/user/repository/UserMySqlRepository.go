package repository

import (
	driver "github.com/rizghz/clean/infrastructure/database"
	"github.com/rizghz/clean/module/user/entity"
	"github.com/rizghz/clean/module/user/transfer"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserMySqlRepository struct {
	Driver *driver.MySqlDriver
}

func NewUserSqlRepository(drv *driver.MySqlDriver) UserRepository {
	return &UserMySqlRepository{
		Driver: drv,
	}
}

func (repo *UserMySqlRepository) Get() []*entity.User {
	data := make([]*entity.User, 0)
	if err := repo.Driver.DB.Find(&data).Error; err != nil {
		logrus.Error(err.Error())
		return nil
	}
	return data
}

func (repo *UserMySqlRepository) Create(data *entity.User) (bool, *entity.User) {
	err := repo.Driver.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(data).Error; err != nil {
			logrus.Error(err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		logrus.Error(err.Error())
		return false, nil
	}
	return true, data
}

func (repo *UserMySqlRepository) Find(req *transfer.UserRequestBody) *entity.User {
	data := new(entity.User)
	condition := "email = ? AND password = ?"
	result := repo.Driver.DB.Where(condition, req.Email, req.Password).First(data)
	if err := result.Error; err != nil {
		logrus.Error(err.Error())
		return nil
	}
	return data
}
