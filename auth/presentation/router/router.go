package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/architecture-template/echo-ddd/log"
	"github.com/architecture-template/echo-ddd/auth/di"
	_ "github.com/architecture-template/echo-ddd/docs/swagger/auth"
)

func Init() {
	// di: wire ./api/di/wire.go
	userController := di.InitializeUserController()

	e := echo.New()

	// Swagger
	e.GET("/swagger/*any", echoSwagger.WrapHandler)

	// Log
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ Output: log.GenerateApiLog() }))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	auth := e.Group("/user")
	auth.POST("/register_user", userController.RegisterUser())
	auth.POST("/login_user", userController.LoginUser())

	e.Logger.Fatal(e.Start(":8000"))
}
