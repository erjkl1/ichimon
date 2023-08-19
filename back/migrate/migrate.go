package main

import (
	"fmt"
	"ichimonApi/db"
	"ichimonApi/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Category{}, &model.SubCategory{}, &model.Grade{}, &model.Question{})
}