package card

import (
	"buhang/bootstrap"
	"time"
)

type Card struct {
	ID           uint64    `json:"id"`
	CategoryId   uint64    `json:"category_id"`
	Img          string    `json:"img"`
	GoodsNumber  string    `json:"goods_number"`
	GoodsName    string    `json:"goods_name"`
	Weight       string    `json:"weight"`
	MeterPerKilo float32   `json:"meter_per_kilo"`
	Width        float32   `json:"width"`
	Sort         int       `json:"sort"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
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
