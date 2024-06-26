//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/haerul-umam/capstone-project-mikti/app"
	"github.com/haerul-umam/capstone-project-mikti/controller"
	"github.com/haerul-umam/capstone-project-mikti/helper"
	"github.com/haerul-umam/capstone-project-mikti/repository"
	"github.com/haerul-umam/capstone-project-mikti/service"
	"github.com/labstack/echo/v4"
)

var authSet = wire.NewSet(
	repository.NewUserRepository,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	helper.NewTokenUseCase,
	wire.Bind(new(helper.TokenUseCase), new(*helper.TokenUseCaseImpl)),
	service.NewUserService,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
	controller.NewAuthController,
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
)

func StartServer() *echo.Echo {
	wire.Build(
		app.InitConnetion,
		authSet,
		app.Router,
	)
	return nil
}