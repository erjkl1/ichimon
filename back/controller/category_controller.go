package controller

import (
	"ichimonApi/model"
	"ichimonApi/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ICategoryController interface {
	FindById(c echo.Context) error
	FindAll(c echo.Context) error
	CreateCategory(c echo.Context) error
	CreateSubCategory(c echo.Context) error
}

type categoryController struct {
	cu usecase.ICategoryUsecase
}

func NewCategoryController(cu usecase.ICategoryUsecase) ICategoryController {
	return &categoryController{cu}
}

func (cc *categoryController) FindById(c echo.Context) error {

	categoryId, _ := strconv.Atoi(c.Param("categoryId"))
	
	resCategory, err := cc.cu.FindById(uint(categoryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, resCategory)
}

//TODO sub_categoryの関連がおかしい？取れない。
func (cc *categoryController) FindAll(c echo.Context) error {	
	categories, err := cc.cu.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, categories)
}

//TODO view_order指定できない問題
func (cc *categoryController) CreateCategory(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	createdUserId := claims["user_id"]

	category := model.Category{}
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// category.UpdatedUserId = uint(createdUserId.(float64))
	category.CreatedUserId = uint(createdUserId.(float64))
	categoryResponse, err := cc.cu.CreateCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, categoryResponse)
}

func (cc *categoryController) CreateSubCategory(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	createdUserId := claims["user_id"]

	subCategory := model.SubCategory{}
	if err := c.Bind(&subCategory); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// subCategory.UpdatedUserId = uint(createdUserId.(float64))
	subCategory.CreatedUserId = uint(createdUserId.(float64))
	categoryResponse, err := cc.cu.CreateSubCategory(subCategory)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, categoryResponse)
}