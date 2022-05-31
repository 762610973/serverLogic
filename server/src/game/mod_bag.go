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
func (m *ModBag) AddItem(itemId int, num int64, player *Player) {
	// 根据id获取物品信息
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "物品不存在")
		return
	}
	switch itemConfig.SortType {
	//根据物品种类进行识别
	/*case csvs.ItemTypeNormal: //普通物品
	//fmt.Println("普通物品：", itemConfig.ItemName)
	m.AddItemToBag(itemId, num)*/
	case csvs.ItemTypeRole: //角色
		//fmt.Println("角色：", itemConfig.ItemName)
		player.ModRole.AddItem(itemId, num, player)
	case csvs.ItemTypeIcon: //头像
		player.ModIcon.AddItem(itemId)
	case csvs.ItemTypeCard: //名片
		player.ModCard.AddItem(itemId, 11)
	case csvs.ItemTypeWeapon: //武器
		player.ModWeapon.AddItem(itemId, num)
	case csvs.ItemTypeRelics: //圣遗物
		player.ModRelics.AddItem(itemId, num)
	case csvs.ItemTypeCook:
		player.ModCook.AddItem(itemId)
	case csvs.ItemTypeHomeItem:
		player.ModHome.AddItem(itemId, num, player)
	default: // 普通
		m.AddItemToBag(itemId, num)
	}
}

// AddItemToBag 增加物品到背包
func (m *ModBag) AddItemToBag(itemId int, num int64) {
	// 判断之前有没有这个东西
	_, ok := m.BagInfo[itemId]
	if ok {
		//有的话直接增加
		m.BagInfo[itemId].ItemNum += num
	} else {
		m.BagInfo[itemId] = &ItemInfo{
			ItemId:  itemId,
			ItemNum: num,
		}
	}
	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("获得物品：", config.ItemName, "------数量：", num)
	}
}

// RemoveItem 往背包里增加物品，通过id来获取这个物品的所有信息，然后根据这个信息进行分类，然后根据不同的类别，调用相应模块的add函数
func (m *ModBag) RemoveItem(itemId int, num int64, player *Player) {
	// 根据id获取物品信息
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "物品不存在")
		return
	}
	switch itemConfig.SortType {
	case csvs.ItemTypeNormal: //普通物品
		m.RemoveItemToBagGM(itemId, num)
	default: //
		//m.RemoveItemToBag(itemId, 1)
	}
}

// RemoveItemToBagGM GM权限移除
func (m *ModBag) RemoveItemToBagGM(itemId int, num int64) {
	// 判断之前有没有这个东西
	_, ok := m.BagInfo[itemId]
	if ok {
		//有的话直接增加
		m.BagInfo[itemId].ItemNum -= num
	} else {
		m.BagInfo[itemId] = &ItemInfo{
			ItemId:  itemId,
			ItemNum: 0 - num,
		}
	}
	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("扣除物品：", config.ItemName, "------数量：", num)
		fmt.Println("当前物品：", config.ItemName, "------数量：", m.BagInfo[itemId].ItemNum)
	}
}

// RemoveItemToBag 移除
func (m *ModBag) RemoveItemToBag(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	switch itemConfig.SortType {
	case csvs.ItemTypeRole:
		fmt.Println("此物品无法扣除")
		return
	case csvs.ItemTypeIcon:
		fmt.Println("此物品无法扣除")
		return
	case csvs.ItemTypeCard:
		fmt.Println("此物品无法扣除")
		return
	default: //同普通
	}

	if !m.HasEnoughItem(itemId, num) {
		config := csvs.GetItemConfig(itemId)
		if config != nil {
			nowNum := int64(0)
			_, ok := m.BagInfo[itemId]
			if ok {
				nowNum = m.BagInfo[itemId].ItemNum
			}
			fmt.Println(config.ItemName, "数量不足", "----当前数量：", nowNum)
		}
		return
	}
	_, ok := m.BagInfo[itemId]
	if ok {
		m.BagInfo[itemId].ItemNum -= num
	} else {
		m.BagInfo[itemId] = &ItemInfo{ItemId: itemId, ItemNum: 0 - num}
	}
	fmt.Println("扣除物品", itemConfig.ItemName, "----数量：", num, "----当前数量：", m.BagInfo[itemId].ItemNum)
}

// HasEnoughItem 判断是否有足够的物品
func (m *ModBag) HasEnoughItem(itemId int, num int64) bool {
	if itemId == 0 {
		return true
	}
	_, ok := m.BagInfo[itemId]
	if !ok {
		return false
	} else if m.BagInfo[itemId].ItemNum < num {
		return false
	}
	return true
}

// UseItem 背包模块的物品使用功能
func (m *ModBag) UseItem(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "物品不存在")
		return
	}
	if !m.HasEnoughItem(itemId, num) {
		config := csvs.GetItemConfig(itemId)
		if config != nil {
			nowNum := int64(0)
			_, ok := m.BagInfo[itemId]
			if ok {
				nowNum = m.BagInfo[itemId].ItemNum
			}
			fmt.Println(config.ItemName, "数量不足", "----当前数量：", nowNum)
		}
		return
	}
	switch itemConfig.SortType {
	case csvs.ItemTypeCook:
		m.UseCookBook(itemId, num, player)
	case csvs.ItemTypeFood:
		//给英雄加属性
	default: //同普通
		fmt.Println(itemId, "此物品无法使用")
		return
	}

}

func (m *ModBag) UseCookBook(itemId int, num int64, player *Player) {
	cookBookConfig := csvs.GetCookBookConfig(itemId)
	if cookBookConfig == nil {
		fmt.Println(itemId, "物品不存在")
		return
	}
	// 使用的话先扣除物品，id、数量
	m.RemoveItem(itemId, num, player)
	m.AddItem(cookBookConfig.Reward, num, player)
}
