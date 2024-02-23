package main

import (
	"ichimonApi/controller"
	"ichimonApi/db"
	"ichimonApi/repository"
	"ichimonApi/router"
	"ichimonApi/usecase"
	"ichimonApi/validator"
)

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