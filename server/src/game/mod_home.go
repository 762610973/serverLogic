package game

import (
	"fmt"
	"serverLogic/server/src/csvs"
)

type HomeItemId struct {
	HomeItemId  int
	HomeItemNum int64
	KeyId       int
}

type ModHome struct {
	HomeItemIdInfo map[int]*HomeItemId
	//UseHomeItemIdInfo map[int]*HomeItemId
	//Map
}

func (m *ModHome) AddItem(itemId int, num int64, player *Player) {
	_, ok := m.HomeItemIdInfo[itemId]
	if ok {
		// 存在的话数量直接相加
		m.HomeItemIdInfo[itemId].HomeItemNum += num
	} else {
		// 不存在则生成
		m.HomeItemIdInfo[itemId] = &HomeItemId{HomeItemId: itemId, HomeItemNum: num}
	}
	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("获得家具物品", config.ItemName, "----数量：", num, "----当前数量：", m.HomeItemIdInfo[itemId].HomeItemNum)
	}
}
