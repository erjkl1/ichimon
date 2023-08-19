package model

import "time"

type Question struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	SubCategory  SubCategory `json:"sub_category" gorm:"foreignKey:SubCategoryId; constraint:OnDelete:CASCADE"`
	Title  string    `json:"title"`
	Statement  string    `json:"statement"`
	Selection1  string    `json:"selection_1"`
	Selection2  string    `json:"selection_2"`
	Selection3  string    `json:"selection_3"`
	Selection4  string    `json:"selection_4"`
	Answer  bool    `json:"answer"`
	Description  string    `json:"description"`
	CreatedUser   User    `json:"created_user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	ModifiedUser   User    `json:"modified_user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	DelFlg    bool      `json:"del_flg"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
