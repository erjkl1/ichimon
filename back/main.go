package main

import (
	"net/http"

	// "ichimonApi/controller"
	"ichimonApi/db"
	// "ichimonApi/repository"
	// "ichimonApi/router"
	// "ichimonApi/usecase"
	// "ichimonApi/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {
	db := db.NewDB()
	e.GET("/login")

	// userValidator := validator.NewUserValidator()
	// taskValidator := validator.NewTaskValidator()
	// userRepository := repository.NewUserRepository(db)
	// taskRepository := repository.NewTaskRepository(db)
	// userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	// taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	// userController := controller.NewUserController(userUsecase)
	// taskController := controller.NewTaskController(taskUsecase)
	// e := router.NewRouter(userController, taskController)
	// e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func articleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}