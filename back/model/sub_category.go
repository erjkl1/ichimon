package model

type SubCategory struct {
	App  App `gorm:"embedded"`
	Category Category   `json:"category" gorm:"foreignKey:CategoryId; constraint:OnDelete:CASCADE"`
	CategoryId uint   `json:"category_id" gorm:"not null"`
	CreatedUser   User    `json:"created_user" gorm:"foreignKey:CreatedUserId; constraint:OnDelete:CASCADE"`
	CreatedUserId   uint `json:"created_user_id" gorm:"not null"`
	SubCategoryName  string    `json:"sub_category_name"`
}
