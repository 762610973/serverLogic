package game

import (
	"fmt"
	"serverLogic/server/src/csvs"
)

type Card struct {
	CardId int
}
type ModCard struct {
	CardInfo map[int]*Card
}

func (m *ModCard) IsHasCard(cardId int) bool {
	_, ok := m.CardInfo[cardId]
	return ok
}

func (m *ModCard) AddItem(itemId int, friendliness int) {
	_, ok := m.CardInfo[itemId]
	if ok {
		fmt.Println("已存在名片：", itemId)
		return
	}
	config := csvs.GetCardConfig(itemId)
	if config == nil {
		fmt.Println("非法名片：", itemId)
		return
	}
	if friendliness < config.Friendliness {
		fmt.Println("好感度不足：", itemId)
		return
	}
	m.CardInfo[itemId] = &Card{itemId}
	fmt.Println("获得名片：", itemId)
}

func (m *ModCard) CheckGetCard(roleId int, friendliness int) {
	config := csvs.GetCardConfigByRoleId(roleId)
	if config == nil {
		return
		//有些英雄获得可能不送头像
	}
	m.AddItem(config.CardId, friendliness)
}
