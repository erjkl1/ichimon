package model

type Grade struct {
	User   User    `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Question   Question    `json:"question" gorm:"foreignKey:QuestionId; constraint:OnDelete:CASCADE"`
	QuestionId   uint    `json:"question_id" gorm:"not null"`
	Correction  bool    `json:"correction"`
	GoodStatus  bool    `json:"good_status"`
	App  App `gorm:"embedded"`
}
