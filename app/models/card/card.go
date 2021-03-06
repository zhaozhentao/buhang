package card

import (
	"buhang/bootstrap"
	"buhang/config"
	"gorm.io/gorm"
	"time"
)

type Card struct {
	ID           uint64    `json:"id"`
	CategoryId   uint64    `json:"category_id"`
	Img          string    `json:"img"`
	GoodsNumber  string    `json:"goodsNumber"`
	GoodsName    string    `json:"goodsName"`
	Weight       string    `json:"weight"`
	MeterPerKilo float64   `json:"meterPerKilo"`
	Width        float64   `json:"width"`
	Sort         int       `json:"sort"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (c *Card) AfterFind(tx *gorm.DB) (err error)  {
	c.Img = config.Viper.GetString("IMAGE_PATH") + c.Img
	return
}

func (Card) TableName() string {
	return "card"
}

func (c *Card) Create() (err error) {
	if err = bootstrap.DB.Create(&c).Error; err != nil {
		return err
	}
	return nil
}

func Index() []Card {
	var cards []Card

	if err := bootstrap.DB.Find(&cards).Error; err != nil {
		return cards
	}

	return cards
}

func Show(id string) Card {
	var card Card
	bootstrap.DB.First(&card, id)
	return card
}
