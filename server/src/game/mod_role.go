package game

import (
	"fmt"
	"serverLogic/server/src/csvs"
)

type RoleInfo struct {
	RoleId     int
	GetTimes   int //添加次数（获得次数），可以通过第几次获得来转化为对应的材料
	RelicsInfo []int
	WeaponInfo int
}

type ModRole struct {
	RoleInfo  map[int]*RoleInfo
	HpPool    int
	HpCalTime int64
	player    *Player
	path      string
}

func (m *ModRole) IsHasRole(roleId int) bool {
	_, ok := m.RoleInfo[roleId]
	return ok
}
func (m *ModRole) GetRoleLevel(roleId int) int {
	//这个函数应该传入角色id，返回角色等级，暂时无法实现
	return 80
}

func (m *ModRole) AddItem(roleId int, num int64, player *Player) {
	config := csvs.GetRoleConfig(roleId)
	if config == nil {
		fmt.Println("配置不存在roleId:", roleId)
		return
	}
	for i := 0; i < int(num); i++ {
		_, ok := m.RoleInfo[roleId]
		if !ok {
			data := new(RoleInfo)
			data.RoleId = roleId
			data.GetTimes = 1
			m.RoleInfo[roleId] = data
			fmt.Println("获得角色次数：", roleId, "----", data.GetTimes, "次")
		} else {
			//判断实际获得的东西
			m.RoleInfo[roleId].GetTimes++
			temp := m.RoleInfo[roleId].GetTimes
			if temp >= csvs.AddRoleTimeNormalMin && temp <= csvs.AddRoleTimeNormalMax {
				player.ModBag.AddItemToBag(config.Stuff, config.StuffNum) //命座材料
				player.ModBag.AddItemToBag(config.StuffItem, config.StuffItemNum)
			} else {
				player.ModBag.AddItemToBag(config.MaxStuffItem, config.MaxStuffItemNum)
			}
		}
	}
	itemConfig := csvs.GetItemConfig(roleId)
	if itemConfig != nil {
		fmt.Println("获得角色", itemConfig.ItemName, "次数", roleId, "------", m.RoleInfo[roleId].GetTimes, "次")
	}
	player.ModIcon.CheckGetIcon(roleId)
	player.ModCard.CheckGetCard(roleId, 10)
}
