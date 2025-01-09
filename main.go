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
	app := fiber.New()
	dbConnection := connection.GetDatabase(cnf.Database)

	userRepository := repository.NewUser(dbConnection)
	authService := service.NewUser(cnf, userRepository)
	api.NewUser(app, authService)

	departmentRepository := repository.NewDepartment(dbConnection)
	departmentService := service.NewDepartment(cnf, departmentRepository)
	api.NewDepartment(app, departmentService)

	employeeRepository := repository.NewEmployee(dbConnection)
	employeeService := service.NewEmployee(cnf, employeeRepository, departmentRepository)
	api.NewEmployee(app, employeeService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
