package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/architecture-template/echo-ddd/config/db"
	"github.com/architecture-template/echo-ddd/domain/model"
	"github.com/architecture-template/echo-ddd/domain/repository"
)

type exampleDao struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func NewExampleDao(conn *db.SqlHandler) repository.ExampleRepository {
	return &exampleDao{
		Read:  conn.ReadConn,
		Write: conn.WriteConn,
	}
}

// List 一覧を取得する
func (e *exampleDao) List(limit int64) (*model.Examples, error) {
	entity := &model.Examples{}
	res := e.Read.
		Limit(limit).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

// FindByKey example_keyで取得する
func (e *exampleDao) FindByExampleKey(email string) (*model.Example, error) {
	entity := &model.Example{}
	res := e.Read.
		Where("example_key = ?", email).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
