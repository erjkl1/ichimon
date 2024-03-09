package main

import (
	"ichimonApi/controller"
	"ichimonApi/db"
	"ichimonApi/repository"
	"ichimonApi/router"
	"ichimonApi/usecase"
	"ichimonApi/validator"

	_ "ichimonApi/docs"
)

// @title IchimonItto API
// @version 1.0
// @description Api for IchimonItto App

// @contact.name API Support
// @contact.url https://twitter.com/Chuke_yamaha

// @host localhost:8080
// @BasePath /
// @schemes https

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]

func main() {
	db := db.NewDB()

	userValidator := validator.NewUserValidator()
	questionValidator := validator.NewQuestionValidator()
	categoryValidator := validator.NewCategoryValidator()
	subCategoryValidator := validator.NewSubCategoryValidator()

	userRepository := repository.NewUserRepository(db)
	questionRepository := repository.NewQuestionRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	// subCategoryRepository := repository.NewCategoryRepository(db)
	
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	questionUsecase := usecase.NewQuestionUsecase(questionRepository, questionValidator)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository, categoryValidator, subCategoryValidator)

	userController := controller.NewUserController(userUsecase)
	questionController := controller.NewQuestionController(questionUsecase)
	categoryController := controller.NewCategoryController(categoryUsecase)

	e := router.NewRouter(userController,questionController,categoryController) 
	e.Logger.Fatal(e.Start(":8080"))
}