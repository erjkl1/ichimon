package model

import "time"

type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedUser   User    `json:"created_user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	CategoryName  string    `json:"category_name"`
	DelFlg    bool      `json:"del_flg"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
