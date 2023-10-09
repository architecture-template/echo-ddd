package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/architecture-template/echo-ddd/domain/model"
)

type UserRepository interface {
	FindByEmail(email string) (*model.User, error)
	Insert(param *model.User, tx *gorm.DB) (*model.User, error)
	Update(param *model.User, tx *gorm.DB) (*model.User, error)
}
