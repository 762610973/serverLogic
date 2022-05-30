package game

import (
	"fmt"
	"serverLogic/server/src/csvs"
)

type Cook struct {
	CookId int
}

type ModCook struct {
	CookInfo map[int]*Cook
}

func (m *ModCook) AddItem(itemId int) {
	_, ok := m.CookInfo[itemId]
	if ok {
		fmt.Println("已习得：", csvs.GetItemName(itemId))
		return
	}
	config := csvs.GetCookConfig(itemId)
	if config == nil {
		fmt.Println("没有这个烹饪技能：", csvs.GetItemName(itemId))
		return
	}
	m.CookInfo[itemId] = &Cook{CookId: itemId}
	fmt.Println("学会烹饪：", itemId)
}
