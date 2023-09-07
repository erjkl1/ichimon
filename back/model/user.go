package model

type User struct {
	App `gorm:"embedded"`
	Email     string    `json:"email" gorm:"unique;not null size:100"`
	Password  string    `json:"password" gorm:"not null"`
	Grades []Grade `json:"grades" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	CreatedCategories []Category `json:"created_categories" gorm:"foreignKey:CreatedUserId; constraint:OnDelete:CASCADE"`
	CreatedSubCategories []SubCategory `json:"created_sub_categories" gorm:"foreignKey:CreatedUserId; constraint:OnDelete:CASCADE"`
	CreatedQuestions []Question `json:"created_questions" gorm:"foreignKey:CreatedUserId; constraint:OnDelete:CASCADE"`
	UpdatedQuestions []Question `json:"updateed_questions" gorm:"foreignKey:UpdatedUserId; constraint:OnDelete:CASCADE"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
}