// +build wireinject

package di

import (
    "github.com/google/wire"

    "github.com/architecture-template/echo-ddd/config/db"
    "github.com/architecture-template/echo-ddd/infra/dao"
    "github.com/architecture-template/echo-ddd/api/service"	
    "github.com/architecture-template/echo-ddd/api/presentation/controller"
)

// example
func InitializeExampleController() controller.ExampleController {
	wire.Build(
		db.NewDB,
		dao.NewExampleDao,
		service.NewExampleService,
		controller.NewExampleController,
	)

    return nil
}
