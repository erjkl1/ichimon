package model

type Category struct {
	App  `gorm:"embedded"`
	CreatedUser   User    `json:"created_user" gorm:"foreignKey:CreatedUserId; constraint:OnDelete:CASCADE"`
	CreatedUserId   uint    `json:"created_user_id" gorm:"not null"`
	CategoryName  string    `json:"category_name"`
}
