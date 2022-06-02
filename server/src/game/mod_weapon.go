package game

import (
	"fmt"
	"serverLogic/server/src/csvs"
)

type Weapon struct {
	WeaponId    int
	KeyId       int
	Level       int
	Exp         int
	StarLevel   int
	RefineLevel int
	RoleId      int
}

type ModWeapon struct {
	WeaponInfo map[int]*Weapon
	MaxKey     int
}

func (m *ModWeapon) AddItem(itemId int, num int64) {

	config := csvs.GetWeaponConfig(itemId)
	if config == nil {
		fmt.Println("配置不存在")
		return
	}

	if len(m.WeaponInfo)+int(num) > csvs.WeaponMaxCount {
		fmt.Println("超过最大值")
		return
	}

	for i := int64(0); i < num; i++ {
		weapon := new(Weapon)
		weapon.WeaponId = itemId
		m.MaxKey++
		weapon.KeyId = m.MaxKey
		m.WeaponInfo[weapon.KeyId] = weapon
		fmt.Println("获得武器:", csvs.GetItemName(itemId), "------武器编号:", weapon.KeyId)
	}
}

func (m *ModWeapon) WeaponUp(keyId int, player *Player) {
	weapon := m.WeaponInfo[keyId]
	if weapon == nil {
		return
	}
	weaponConfig := csvs.GetWeaponConfig(weapon.WeaponId)
	if weaponConfig == nil {
		return
	}
	weapon.Exp += 5000
	for {
		// 去除下一级配置
		nextLevelConfig := csvs.GetWeaponLevelConfig(weaponConfig.Star, weapon.Level+1)
		if nextLevelConfig == nil {
			// 武器不存在，说明已经到了上限
			fmt.Println("返还武器经验:", weapon.Exp)
			weapon.Exp = 0
			break
		}
		if weapon.StarLevel < nextLevelConfig.NeedStarLevel {
			fmt.Println("返还武器经验:", weapon.Exp)
			weapon.Exp = 0
			break
		}
		if weapon.Exp < nextLevelConfig.NeedExp {
			break
		}
		weapon.Level++
		weapon.Exp -= nextLevelConfig.NeedExp
	}
	weapon.ShowInfo()
}

func (w *Weapon) ShowInfo() {
	fmt.Println(fmt.Sprintf("key:%d,Id:%d", w.KeyId, w.WeaponId))
	fmt.Println(fmt.Sprintf("当前等级:%d,当前经验:%d,当前突破等级:%d,当前精炼等级:%d",
		w.Level, w.Exp, w.StarLevel, w.RefineLevel))
}

func (m *ModWeapon) WeaponUpStar(keyId int, player *Player) {
	weapon := m.WeaponInfo[keyId]
	if weapon == nil {
		return
	}
	weaponConfig := csvs.GetWeaponConfig(weapon.WeaponId)
	if weaponConfig == nil {
		return
	}
	nextStarConfig := csvs.GetWeaponStarConfig(weaponConfig.Star, weapon.StarLevel+1)
	if nextStarConfig == nil {
		return
	}
	//验证物品充足并扣除
	//........
	if weapon.Level < nextStarConfig.Level {
		fmt.Println("武器等级不够，无法突破")
		return
	}
	weapon.StarLevel++
	weapon.ShowInfo()
}

// WeaponUpRefine 武器精炼
func (m *ModWeapon) WeaponUpRefine(keyId int, targetKeyId int, player *Player) {
	if keyId == targetKeyId {
		fmt.Println("错误的材料")
		return
	}
	weapon := m.WeaponInfo[keyId]
	if weapon == nil {
		return
	}
	weaponTarget := m.WeaponInfo[targetKeyId]
	if weaponTarget == nil {
		return
	}
	if weapon.WeaponId != weaponTarget.WeaponId {
		fmt.Println("错误的材料")
		return
	}
	if weapon.RefineLevel >= csvs.WeaponMaxCount {
		fmt.Println("超过了最大精炼等级")
		return
	}
	weapon.RefineLevel++
	delete(m.WeaponInfo, targetKeyId)
	weapon.ShowInfo()
}
