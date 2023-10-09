// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"echo-mangosteen/internal/common/data"
	"echo-mangosteen/internal/controller"
	"echo-mangosteen/internal/repo"
	"echo-mangosteen/internal/router"
	"echo-mangosteen/internal/service"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

// Injectors from wire.go:

func NewApp() (*echo.Echo, func(), error) {
	dataData, cleanup, err := data.NewData()
	if err != nil {
		return nil, nil, err
	}
	pingController := controller.NewPingController(dataData)
	userRepo := repo.NewUserRepo(dataData)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	echoEcho := router.NewRouter(pingController, userController)
	return echoEcho, func() {
		cleanup()
	}, nil
}

// wire.go:

var controllerProvider = wire.NewSet(controller.NewPingController, controller.NewUserController)

var repoProvider = wire.NewSet(repo.NewUserRepo)

var serviceProvider = wire.NewSet(service.NewUserService)
