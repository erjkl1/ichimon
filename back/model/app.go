package model

import "gorm.io/gorm"

type App struct {
	gorm.Model
	DelFlg    bool     `json:"del_flg"`
}