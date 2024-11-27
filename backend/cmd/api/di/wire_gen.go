// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"backend/internal/controller"
	"backend/internal/repository"
	"backend/internal/routes"
	"backend/internal/service"
)

// Injectors from wire.go:

func InitializedApp() (IPillar, error) {
	iHashService := service.ProvideHashService()
	iUserRepository := repository.ProvideUserRepository()
	iAuthService := service.ProvideAuthService(iHashService, iUserRepository)
	iTokenService := service.ProvideTokenService()
	iAuthController := controller.ProvideAuthController(iAuthService, iTokenService)
	authRouter := routes.ProvideAuthRouter(iAuthController)
	iRouter := routes.ProvideRouter(authRouter)
	iPillar := ProvideGinEngine(iRouter)
	return iPillar, nil
}