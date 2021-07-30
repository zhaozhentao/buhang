package card

import (
	"buhang/bootstrap"
)

type Card struct {
	ID         uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`
	CategoryId uint64 `gorm:"column:category_id"`
}

func (Card) TableName() string {
	return "card"
}

func Index() []Card {
	var cards []Card
	if err := bootstrap.DB.Find(&cards).Error; err != nil {
		return cards
	}
	return cards
}
