package repository

import (
	"github.com/rizghz/clean/module/user/entity"
	"github.com/rizghz/clean/module/user/transfer"
)

type UserRepository interface {
	Get() []*entity.User
	Create(data *entity.User) (bool, *entity.User)
	Find(data *transfer.UserRequestBody) *entity.User
}
