package model

import "time"

type SubCategory struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Category Category   `json:"category" gorm:"foreignKey:CategoryId; constraint:OnDelete:CASCADE"`
	CreatedUser   User    `json:"created_user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	SubCategoryName  string    `json:"sub_category_name"`
	DelFlg    bool      `json:"del_flg"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
