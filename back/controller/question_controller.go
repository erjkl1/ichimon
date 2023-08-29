package controller

import (
	"ichimonApi/model"
	"ichimonApi/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IQuestionController interface {
	FindById(c echo.Context) error
	FindAll(c echo.Context) error
	FindAllByCategoryId(c echo.Context) error
	FindAllBySubCategoryId(c echo.Context) error
	FindAllByCreatedUserId(c echo.Context) error
	CreateQuestion(c echo.Context) error
	// CsrfToken(c echo.Context) error
}

type questionController struct {
	qu usecase.IQuestionUsecase
}

func NewQuestionController(qu usecase.IQuestionUsecase) IQuestionController {
	return &questionController{qu}
}

func (qc *questionController) FindById(c echo.Context) error {

	questionId, _ := strconv.Atoi(c.Param("questionId"))
	
	questionResponse, err := qc.qu.FindById(uint(questionId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, questionResponse)
}

func (qc *questionController) FindAll(c echo.Context) error {

	
	questionResponses, err := qc.qu.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, questionResponses)
}

func (qc *questionController) FindAllByCategoryId(c echo.Context) error {

	categoryId, _ := strconv.Atoi(c.Param("categoryId"))
	
	questionResponses, err := qc.qu.FindAllByCategoryId(uint(categoryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, questionResponses)
}

func (qc *questionController) FindAllBySubCategoryId(c echo.Context) error {

	categoryId, _ := strconv.Atoi(c.Param("subCategoryId"))
	
	questionResponses, err := qc.qu.FindAllBySubCategoryId(uint(categoryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, questionResponses)
}

func (qc *questionController) FindAllByCreatedUserId(c echo.Context) error {

	categoryId, _ := strconv.Atoi(c.Param("createdUserId"))
	
	questionResponses, err := qc.qu.FindAllByCreatedUserId(uint(categoryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, questionResponses)
}

func (qc *questionController) CreateQuestion(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	createdUserId := claims["user_id"]

	question := model.Question{}
	if err := c.Bind(&question); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	question.UpdatedUserId = uint(createdUserId.(float64))
	question.CreatedUserId = uint(createdUserId.(float64))
	questionResponse, err := qc.qu.CreateQuestion(question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, questionResponse)
}

func (qc *questionController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}