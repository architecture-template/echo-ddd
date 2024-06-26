// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/architecture-template/echo-ddd/api/presentation/controller"
	"github.com/architecture-template/echo-ddd/api/presentation/middleware"
	"github.com/architecture-template/echo-ddd/api/service"
	"github.com/architecture-template/echo-ddd/config/db"
	"github.com/architecture-template/echo-ddd/infra/dao"
)

// Injectors from wire.go:

// example
func InitializeExampleController() controller.ExampleController {
	sqlHandler := db.NewDB()
	exampleRepository := dao.NewExampleDao(sqlHandler)
	exampleService := service.NewExampleService(exampleRepository)
	exampleController := controller.NewExampleController(exampleService)
	return exampleController
}

// user
func InitializeUserMiddleware() middleware.UserMiddleware {
	userMiddleware := middleware.NewUserMiddleware()
	return userMiddleware
}
