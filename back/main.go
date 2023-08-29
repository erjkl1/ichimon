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
	// questionValidator := validator.NewQuestionValidator()
	userRepository := repository.NewUserRepository(db)
	questionRepository := repository.NewQuestionRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	questionUsecase := usecase.NewQuestionUsecase(questionRepository)
	userController := controller.NewUserController(userUsecase)
	questionController := controller.NewQuestionController(questionUsecase)
	e := router.NewRouter(userController,questionController)
	e.Logger.Fatal(e.Start(":8080"))
}