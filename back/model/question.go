package model

type Question struct {
	App  App `gorm:"embedded"`
	CreatedUser   User    `json:"created_user" gorm:"foreignKey:CreatedUserId; constraint:OnDelete:CASCADE"`
	UpdatedUser   User    `json:"updated_user" gorm:"foreignKey:UpdatedUserId; constraint:OnDelete:CASCADE"`
	CreatedUserId   uint `json:"created_user_id" gorm:"not null"`
	UpdatedUserId   uint `json:"updated_user_id" gorm:"not null"`
	SubCategory  SubCategory `json:"sub_category" gorm:"foreignKey:SubCategoryId; constraint:OnDelete:CASCADE"`
	SubCategoryId  uint `json:"sub_category_id" gorm:"not null"`
	Title  string    `json:"title"`
	Statement  string    `json:"statement"`
	Selection1  string    `json:"selection_1"`
	Selection2  string    `json:"selection_2"`
	Selection3  string    `json:"selection_3"`
	Selection4  string    `json:"selection_4"`
	Answer  bool    `json:"answer"`
	Description  string    `json:"description"`
}
