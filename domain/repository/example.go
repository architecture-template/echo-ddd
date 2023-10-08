package repository

import (
	"github.com/architecture-template/echo-ddd/domain/model"
)

type ExampleRepository interface {
	List(limit int64) (*model.Examples, error) 
	FindByKey(testKey string) (*model.Example, error)
}
