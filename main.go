package main

import (
	"github.com/bayuuat/go-sprint-1/internal/api"
	"github.com/bayuuat/go-sprint-1/internal/config"
	"github.com/bayuuat/go-sprint-1/internal/connection"
	"github.com/bayuuat/go-sprint-1/internal/repository"
	"github.com/bayuuat/go-sprint-1/internal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()

	dbConnection := connection.GetDatabase(cnf.Database)

	userRepository := repository.NewUser(dbConnection)
	authService := service.NewUser(cnf, userRepository)

	app := fiber.New()

	api.NewUser(app, authService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
