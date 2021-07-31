package card_item

import (
	"buhang/bootstrap"
	"buhang/pkg/types"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CardItem struct {
	ID     uint64 `json:"id"`
	CardId uint64 `json:"card_id"`
	Image  string `json:"image"`
	Sort   int    `json:"sort"`
}

func (CardItem) TableName() string {
	return "card_item"
}

func Items(cardId string) []CardItem {
	var items []CardItem

	bootstrap.DB.Find(&items).Where(gin.H{"card_id": cardId})

	return items
}

func Create(images []string, cardId string) {
	_cardId := types.StringToInt(cardId)

	var items = make([]CardItem, len(images))

	for i := range images {
		items[i] = CardItem{
			CardId: _cardId,
			Image:  images[i],
		}
	}

	if err := bootstrap.DB.CreateInBatches(&items, len(items)).Error; err != nil {
		fmt.Println(err)
	}
}
