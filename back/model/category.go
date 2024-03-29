package model

type Category struct {
	App  `gorm:"embedded"`
	CreatedUser   User    `json:"created_user" gorm:"foreignKey:CreatedUserId; constraint:OnDelete:CASCADE"`
	CreatedUserId   uint    `json:"created_user_id" gorm:"not null"`
	CategoryName  string    `json:"category_name" gorm:"size:50"`
	ViewOrder  uint    `json:"view_order"`
	SubCategories []SubCategory  `json:"sub_categories" gorm:"foreignKey:CategoryId; constraint:OnDelete:CASCADE"`
}

type CategoryResponse struct {
	App  `gorm:"embedded"`
	CreatedUserId   uint `json:"created_user_id" gorm:"not null"`
	CategoryName  string    `json:"category_name" gorm:"size:50"`
	ViewOrder  uint    `json:"view_order"`
	SubCategories []SubCategory  `json:"sub_categories" gorm:"foreignKey:CategoryId; constraint:OnDelete:CASCADE"`
}

type CreateCategoryRequest struct {
	CategoryName  string    `json:"category_name" gorm:"size:50"`
	ViewOrder  uint    `json:"view_order"`
}
