package game

import (
	"fmt"
	"serverLogic/server/src/csvs"
	"time"
)

// RoleInfo 角色
type RoleInfo struct {
	RoleId     int
	GetTimes   int //添加次数（获得次数），可以通过第几次获得来转化为对应的材料
	RelicsInfo []int
	WeaponInfo int
}

type ModRole struct {
	RoleInfo  map[int]*RoleInfo
	HpPool    int   // 血池
	HpCalTime int64 // 血池恢复时间
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

// AddItem 增加角色
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

func (m *ModRole) HandleSendRoleInfo(player *Player) {
	fmt.Println("当前拥有角色信息如下")
	for _, v := range m.RoleInfo {
		v.SendRoleInfo(player)
	}
}

func (r *RoleInfo) SendRoleInfo(player *Player) {
	fmt.Println(fmt.Sprintf("%s:,累计获得次数:%d", csvs.GetItemName(r.RoleId), r.GetTimes))
	r.ShowInfo(player)
}

func (r *RoleInfo) ShowInfo(player *Player) {
	fmt.Println(fmt.Sprintf("当前角色:%s,角色ID:%d", csvs.GetItemName(r.RoleId), r.RoleId))

	weaponNow := player.ModWeapon.WeaponInfo[r.WeaponInfo]
	if weaponNow == nil {
		fmt.Println(fmt.Sprintf("武器:未穿戴"))
	} else {
		fmt.Println(fmt.Sprintf("武器:%s,key:%d", csvs.GetItemName(weaponNow.WeaponId), r.WeaponInfo))
	}

	suitMap := make(map[int]int)
	for _, v := range r.RelicsInfo {
		relicsNow := player.ModRelics.RelicsInfo[v]
		if relicsNow == nil {
			fmt.Println(fmt.Sprintf("未穿戴"))
			continue
		}
		fmt.Println(fmt.Sprintf("%s,key:%d", csvs.GetItemName(relicsNow.RelicsId), v))
		relicsNowConfig := csvs.GetRelicsConfig(relicsNow.RelicsId)
		if relicsNowConfig != nil {
			suitMap[relicsNowConfig.Type]++
		}
	}

	suitSkill := make([]int, 0)
	for suit, num := range suitMap {
		for _, config := range csvs.ConfigRelicsSuitMap[suit] {
			if num >= config.Num {
				suitSkill = append(suitSkill, config.SuitSkill)
			}
		}
	}
	for _, v := range suitSkill {
		fmt.Println(fmt.Sprintf("激活套装效果:%d", v))
	}

}

func (m *ModRole) GetRoleInfoForPoolCheck() (map[int]int, map[int]int) {
	fiveInfo := make(map[int]int)
	fourInfo := make(map[int]int)

	for _, v := range m.RoleInfo {
		// 获取配置
		roleConfig := csvs.GetRoleConfig(v.RoleId)
		if roleConfig == nil {
			continue
		}
		if roleConfig.Star == 5 {
			fiveInfo[roleConfig.RoleId] = v.GetTimes
		} else if roleConfig.Star == 4 {
			fourInfo[roleConfig.RoleId] = v.GetTimes
		}
	}
	return fiveInfo, fourInfo
}

// CalHpPool 血池回复
func (m *ModRole) CalHpPool() {
	if m.HpCalTime == 0 {
		m.HpCalTime = time.Now().Unix()
	}
	calTime := time.Now().Unix() - m.HpCalTime
	// 当前减上次的
	m.HpPool += int(calTime) * 10
	m.HpCalTime = time.Now().Unix()
	fmt.Println("当前血池回复量：", m.HpPool)
}

// WearRelics 把圣遗物穿在角色身上
func (m *ModRole) WearRelics(roleInfo *RoleInfo, relics *Relics, player *Player) {
	relicsConfig := csvs.GetRelicsConfig(relics.RelicsId)
	if relicsConfig == nil {
		return
	}
	m.CheckRelicsPos(roleInfo, relicsConfig.Pos)
	if relicsConfig.Pos < 0 || relicsConfig.Pos > len(roleInfo.RelicsInfo) {
		return
	}

	oldRelicsKeyId := roleInfo.RelicsInfo[relicsConfig.Pos-1]
	if oldRelicsKeyId > 0 {
		oldRelics := player.ModRelics.RelicsInfo[oldRelicsKeyId]
		if oldRelics != nil {
			oldRelics.RoleId = 0
		}
		roleInfo.RelicsInfo[relicsConfig.Pos-1] = 0
	}
	oldRoleId := relics.RoleId
	if oldRoleId > 0 {
		oldRole := player.ModRole.RoleInfo[oldRoleId]
		if oldRole != nil {
			oldRole.RelicsInfo[relicsConfig.Pos-1] = 0
		}
		relics.RoleId = 0
	}

	roleInfo.RelicsInfo[relicsConfig.Pos-1] = relics.KeyId
	relics.RoleId = roleInfo.RoleId

	if oldRelicsKeyId > 0 && oldRoleId > 0 {
		oldRelics := player.ModRelics.RelicsInfo[oldRelicsKeyId]
		oldRole := player.ModRole.RoleInfo[oldRoleId]
		if oldRelics != nil && oldRole != nil {
			m.WearRelics(oldRole, oldRelics, player)
		}
	}
	roleInfo.ShowInfo(player)
}

// CheckRelicsPos 检查圣遗物的位置
func (m *ModRole) CheckRelicsPos(roleInfo *RoleInfo, pos int) {
	nowSize := len(roleInfo.RelicsInfo)
	needAdd := pos - nowSize

	for i := 0; i < needAdd; i++ {
		roleInfo.RelicsInfo = append(roleInfo.RelicsInfo, 0)
	}
}

// TakeOffRelics 脱下圣遗物
func (m *ModRole) TakeOffRelics(roleInfo *RoleInfo, relics *Relics, player *Player) {
	relicsConfig := csvs.GetRelicsConfig(relics.RelicsId)
	if relicsConfig == nil {
		return
	}
	m.CheckRelicsPos(roleInfo, relicsConfig.Pos)
	if relicsConfig.Pos < 0 || relicsConfig.Pos > len(roleInfo.RelicsInfo) {
		return
	}
	if roleInfo.RelicsInfo[relicsConfig.Pos-1] != relics.KeyId {
		fmt.Println(fmt.Sprintf("当前角色没有穿戴这个物品"))
		return
	}

	roleInfo.RelicsInfo[relicsConfig.Pos-1] = 0
	relics.RoleId = 0
	roleInfo.ShowInfo(player)
}

// WearWeapon 武器穿戴
func (m *ModRole) WearWeapon(roleInfo *RoleInfo, weapon *Weapon, player *Player) {
	weaponConfig := csvs.GetWeaponConfig(weapon.WeaponId)
	if weaponConfig == nil {
		fmt.Println("数据异常，武器配置不存在")
		return
	}

	//先判断武器和角色是否匹配
	roleConfig := csvs.GetRoleConfig(roleInfo.RoleId)
	if roleConfig.Type != weaponConfig.Type {
		fmt.Println("武器和角色不匹配")
		return
	}
	// 判定有没有之前的武器
	oldWeaponKey := 0
	if roleInfo.WeaponInfo > 0 {
		oldWeaponKey = roleInfo.WeaponInfo
		roleInfo.WeaponInfo = 0
		// 解除和这把武器的关系
		oldWeapon := player.ModWeapon.WeaponInfo[oldWeaponKey]
		if oldWeapon != nil {
			oldWeapon.RoleId = 0
		}
	}
	// 穿戴的这把武器是否绑定别的人
	oldRoleId := 0
	if weapon.RoleId > 0 {
		oldRoleId = weapon.RoleId
		weapon.RoleId = 0
		oldRole := player.ModRole.RoleInfo[oldRoleId]
		if oldRole != nil {
			oldRole.WeaponInfo = 0
		}
	}

	roleInfo.WeaponInfo = weapon.KeyId
	weapon.RoleId = roleInfo.RoleId

	if roleInfo.WeaponInfo > 0 && weapon.RoleId > 0 {
		oldWeapon := player.ModWeapon.WeaponInfo[oldWeaponKey]
		oldRole := player.ModRole.RoleInfo[oldRoleId]
		if oldWeapon != nil && oldRole != nil {
			m.WearWeapon(oldRole, oldWeapon, player)
		}
	}
}

// TakeOffWeapon 脱下武器
func (m *ModRole) TakeOffWeapon(roleInfo *RoleInfo, weapon *Weapon, player *Player) {
	weaponConfig := csvs.GetWeaponConfig(weapon.WeaponId)
	if weaponConfig == nil {
		fmt.Println("数据异常，武器配置不存在")
		return
	}
	if roleInfo.WeaponInfo != weapon.KeyId {
		fmt.Println("角色没有装备这把武器")
		return
	}
	//根据位置看是否身上有对应圣遗物
	roleInfo.WeaponInfo = 0
	weapon.RoleId = 0
}
