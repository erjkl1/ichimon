package main

import (

	// "ichimonApi/controller"
	"ichimonApi/controller"
	"ichimonApi/db"
	"ichimonApi/repository"
	"ichimonApi/router"
	"ichimonApi/usecase"
	// "ichimonApi/repository"
	// "ichimonApi/router"
	// "ichimonApi/usecase"
	// "ichimonApi/validator"
)

func main() {
	db := db.NewDB()

	// userValidator := validator.NewUserValidator()
	// taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	// taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	// taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	// taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":3000"))
}