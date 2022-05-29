package game

import (
	"fmt"
	"serverLogic/server/src/csvs"
)

type ModBag struct {
	BagInfo map[int]*ItemInfo
}
type ItemInfo struct {
	ItemId  int
	ItemNum int64
}

// AddItem 往背包里增加物品，通过id来获取这个物品的所有信息，然后根据这个信息进行分类，然后根据不同的类别，调用相应模块的add函数
func (m *ModBag) AddItem(itemId int, player *Player) {
	// 根据id获取物品信息
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "物品不存在")
		return
	}
	switch itemConfig.SortType {
	//根据物品种类进行识别
	case csvs.ItemTypeNormal:
		fmt.Println("普通物品：", itemConfig.ItemName)
	case csvs.ItemTypeRole:
		fmt.Println("角色：", itemConfig.ItemName)
	case csvs.ItemTypeIcon:
		player.ModIcon.AddItem(itemId)
	case csvs.ItemTypeCard:
		fmt.Println("名片：", itemConfig.ItemName)
	}
}
