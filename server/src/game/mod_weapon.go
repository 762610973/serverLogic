package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	MaxKey     int //确保不会获得重复的武器
	player     *Player
	path       string
}

// AddItem 增加武器
func (m *ModWeapon) AddItem(itemId int, num int64) {
	//防止出现武器能加到你身上，但是配置读不出来的情况
	//判断配置表有没有这个东西
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
		nextLevelConfig := csvs.GetWeaponLevelConfig(weaponConfig.Star, weapon.Level+1)
		if nextLevelConfig == nil {
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

func (self *Weapon) ShowInfo() {
	fmt.Println(fmt.Sprintf("key:%d,Id:%d", self.KeyId, self.WeaponId))
	fmt.Println(fmt.Sprintf("当前等级:%d,当前经验:%d,当前突破等级:%d,当前精炼等级:%d",
		self.Level, self.Exp, self.StarLevel, self.RefineLevel))
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
	if weapon.RefineLevel >= csvs.WeaponMaxRefine {
		fmt.Println("超过了最大精炼等级")
		return
	}
	weapon.RefineLevel++
	delete(m.WeaponInfo, targetKeyId)
	weapon.ShowInfo()
}

func (m *ModWeapon) SaveData() {
	content, err := json.Marshal(m)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(m.path, content, os.ModePerm)
	if err != nil {
		return
	}
}

func (m *ModWeapon) LoadData(player *Player) {

	m.player = player
	m.path = m.player.localPath + "/weapon.json"

	configFile, err := ioutil.ReadFile(m.path)
	if err != nil {
		m.InitData()
		return
	}
	err = json.Unmarshal(configFile, &m)
	if err != nil {
		m.InitData()
		return
	}

	if m.WeaponInfo == nil {
		m.WeaponInfo = make(map[int]*Weapon)
	}
	return
}

func (m *ModWeapon) InitData() {
	if m.WeaponInfo == nil {
		m.WeaponInfo = make(map[int]*Weapon)
	}
}
