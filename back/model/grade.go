package model

import "time"

type Grade struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	User   User    `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Question   Question    `json:"question" gorm:"foreignKey:QuestionId; constraint:OnDelete:CASCADE"`
	Correction  bool    `json:"correction"`
	GoodStatus  bool    `json:"good_status"`
	DelFlg    bool      `json:"del_flg"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
