package card_item

import (
	"buhang/bootstrap"
	"github.com/gin-gonic/gin"
)

type CardItem struct {
}

func Items(cardId string) []CardItem {
	var items []CardItem

	bootstrap.DB.Find(&items).Where(gin.H{"card_id": cardId})

	return items
}
