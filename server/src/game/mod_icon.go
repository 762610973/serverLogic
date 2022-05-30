package game

import (
	"fmt"
	"serverLogic/server/src/csvs"
)

type Icon struct {
	IconId int
}

type ModIcon struct {
	IconInfo map[int]*Icon
}

// IsHasIcon 判断Icon是否存在，把非法请求拦截在外
func (m *ModIcon) IsHasIcon(iconId int) bool {
	_, ok := m.IconInfo[iconId]
	return ok
}

// AddItem 增加头像
func (m *ModIcon) AddItem(itemId int) {
	_, ok := m.IconInfo[itemId]
	if ok {
		fmt.Println("已存在头像：", itemId)
		return
	}
	config := csvs.GetItemConfig(itemId)
	if config == nil {
		fmt.Println("非法头像：", itemId)
		return
	}
	m.IconInfo[itemId] = &Icon{itemId}
	fmt.Println("获得头像：", itemId)
}

func (m *ModIcon) CheckGetIcon(roleId int) {
	config := csvs.GetIconConfigByRoleId(roleId)
	if config == nil {
		return
		//有些英雄获得可能不送头像
	}
	m.AddItem(config.IconId)
}
