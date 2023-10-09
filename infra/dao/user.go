package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/architecture-template/echo-ddd/config/db"
	"github.com/architecture-template/echo-ddd/domain/model"
	"github.com/architecture-template/echo-ddd/domain/repository"
)

type userDao struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func NewUserDao(conn *db.SqlHandler) repository.UserRepository {
	return &userDao{
		Read:  conn.ReadConn,
		Write: conn.WriteConn,
	}
}

// FindByEmail emailで取得する
func (u *userDao) FindByEmail(email string) (*model.User, error) {
	entity := &model.User{}
	res := u.Read.
		Where("email = ?", email).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}


func (u *userDao) Insert(param *model.User, tx *gorm.DB) (*model.User, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = u.Write
	}

	entity := &model.User{
		UserKey:  param.UserKey,
		UserName: param.UserName,
		Email:    param.Email,
		Password: param.Password,
		Token:    param.Token,
	}
	
	res := conn.Model(&model.User{}).Create(entity)
	if err := res.Error; err != nil {
		return entity, err
	}

	return entity, nil
}

func (u *userDao) Update(param *model.User, tx *gorm.DB) (*model.User, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = u.Write
	}
	
	entity := &model.User{
		UserKey:     param.UserKey,
		UserName:    param.UserName,
		Email:       param.Email,
		Password:    param.Password,
		Token:       param.Token,
	}

	res := conn.Model(&model.User{}).
		Where("user_key = ?", entity.UserKey).
		Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
