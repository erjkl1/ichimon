package model

type Grade struct {
	App  App `gorm:"embedded"`
	User   User    `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId uint   `json:"user_id" gorm:"not null"`
	Question   Question    `json:"question" gorm:"foreignKey:QuestionId; constraint:OnDelete:CASCADE"`
	QuestionId   uint    `json:"question_id" gorm:"not null"`
	Correction  bool    `json:"correction"`
	GoodStatus  bool    `json:"good_status"`
}
