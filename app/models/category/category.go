package category

import (
	"buhang/bootstrap"
	"fmt"
	"time"
)

func (Category) TableName() string {
	return "category"
}

type Category struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Create(name string) {
	var category = Category{Name: name}

	if err := bootstrap.DB.Create(&category).Error; err != nil {
		fmt.Println(err)
	}
}

func List() []Category {
	var categories []Category
	bootstrap.DB.Find(&categories)
	return categories
}
