// +build wireinject

package di

import (
    "github.com/google/wire"

    "github.com/architecture-template/echo-ddd/config/db"
    "github.com/architecture-template/echo-ddd/infra/dao"
    "github.com/architecture-template/echo-ddd/auth/service"	
    "github.com/architecture-template/echo-ddd/auth/presentation/controller"
	"github.com/architecture-template/echo-ddd/api/presentation/middleware"
)

// example
func InitializeUserController() controller.UserController {
	wire.Build(
		db.NewDB,
		dao.NewUserDao,
		dao.NewTransactionDao,
		service.NewUserService,
		controller.NewUserController,
	)

    return nil
}

// user
func InitializeUserMiddleware() middleware.UserMiddleware {
    wire.Build(
		middleware.NewUserMiddleware,
    )
    return nil
}
