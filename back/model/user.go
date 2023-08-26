package model

type User struct {
	App `gorm:"embedded"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Grades []Grade `json:"grades" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	CreatedCategories []Category `json:"created_categories" gorm:"foreignKey:CreatedUserId; constraint:OnDelete:CASCADE"`
	CreatedSubCategories []Grade `json:"created_sub_categories" gorm:"foreignKey:CreatedUserId; constraint:OnDelete:CASCADE"`
	CreatedQuestions []Question `json:"created_questions" gorm:"foreignKey:CreatedUserId; constraint:OnDelete:CASCADE"`
	ModifiedQuestions []Question `json:"modified_questions" gorm:"foreignKey:ModifiedUserId; constraint:OnDelete:CASCADE"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
}