package csvs

import (
	"fmt"
	"math/rand"
)

// CheckLoadCsv 通常建立一个check函数，来表示初始化的结束
// CheckLoadCsv 初始化完成之后，表示结束

var (
	ConfigDropGroupMap        map[int]*DropGroup
	ConfigDropItemGroupMap    map[int]*DropItemGroup
	ConfigStatueMap           map[int]map[int]*ConfigStatue
	ConfigRelicsEntryGroupMap map[int]map[int]*ConfigRelicsEntry
	ConfigRelicsLevelMap      map[int]map[int]*ConfigRelicsLevel
	ConfigRelicsSuitMap       map[int][]*ConfigRelicsSuit
	ConfigWeaponLevelMap      map[int]map[int]*ConfigWeaponLevel
	ConfigWeaponStarMap       map[int]map[int]*ConfigWeaponStar
)

type DropGroup struct {
	DropId      int
	WeightAll   int // 权重总和
	DropConfigs []*ConfigDrop
}

type DropItemGroup struct {
	DropId      int
	DropConfigs []*ConfigDropItem
}

func CheckLoadCsv() {
	//二次处理
	MakeDropGroupMap()
	MakeDropItemGroupMap()
	MakeConfigStatueMap()
	MakeConfigRelicsEntryGroupMap()
	MakeConfigRelicsLevelMap()
	MakeConfigRelicsSuitMap()
	MakeConfigWeaponLevelMap()
	MakeConfigWeaponStarMap()
	fmt.Println("csv配置读取完成---ok")
}

// MakeDropGroupMap 生成掉落组的map
func MakeDropGroupMap() {
	ConfigDropGroupMap = make(map[int]*DropGroup)
	// 遍历这个map，读取响应配置即可
	for _, v := range ConfigDropSlice {
		dropGroup, ok := ConfigDropGroupMap[v.DropId]
		if !ok {
			dropGroup = new(DropGroup)
			dropGroup.DropId = v.DropId
			ConfigDropGroupMap[v.DropId] = dropGroup
		}
		// 权重总和
		dropGroup.WeightAll += v.Weight
		dropGroup.DropConfigs = append(dropGroup.DropConfigs, v)
	}
	//RandDropTest()
	return
}

func MakeDropItemGroupMap() {
	ConfigDropItemGroupMap = make(map[int]*DropItemGroup)
	for _, v := range ConfigDropItemSlice {
		dropGroup, ok := ConfigDropItemGroupMap[v.DropId]
		if !ok {
			dropGroup = new(DropItemGroup)
			dropGroup.DropId = v.DropId
			ConfigDropItemGroupMap[v.DropId] = dropGroup
		}
		dropGroup.DropConfigs = append(dropGroup.DropConfigs, v)
	}
	//configs:=GetDropItemGroupNew(4)
	//println(configs)
	return
}

func MakeConfigStatueMap() {
	ConfigStatueMap = make(map[int]map[int]*ConfigStatue)
	for _, v := range ConfigStatueSlice {
		statueMap, ok := ConfigStatueMap[v.StatueId]
		if !ok {
			statueMap = make(map[int]*ConfigStatue)
			ConfigStatueMap[v.StatueId] = statueMap
		}
		statueMap[v.Level] = v
	}
	return
}

func MakeConfigRelicsEntryGroupMap() {
	ConfigRelicsEntryGroupMap = make(map[int]map[int]*ConfigRelicsEntry)
	for _, v := range ConfigRelicsEntryMap {
		groupMap, ok := ConfigRelicsEntryGroupMap[v.Group]
		if !ok {
			groupMap = make(map[int]*ConfigRelicsEntry)
			ConfigRelicsEntryGroupMap[v.Group] = groupMap
		}
		groupMap[v.Id] = v
	}
	return
}

func MakeConfigRelicsLevelMap() {
	ConfigRelicsLevelMap = make(map[int]map[int]*ConfigRelicsLevel)
	for _, v := range ConfigRelicsLevelSlice {
		levelMap, ok := ConfigRelicsLevelMap[v.EntryId]
		if !ok {
			levelMap = make(map[int]*ConfigRelicsLevel)
			ConfigRelicsLevelMap[v.EntryId] = levelMap
		}
		levelMap[v.Level] = v
	}
	return
}

func MakeConfigRelicsSuitMap() {
	ConfigRelicsSuitMap = make(map[int][]*ConfigRelicsSuit)
	for _, v := range ConfigRelicsSuitSlice {
		ConfigRelicsSuitMap[v.Type] = append(ConfigRelicsSuitMap[v.Type], v)
	}
	return
}

func MakeConfigWeaponLevelMap() {
	ConfigWeaponLevelMap = make(map[int]map[int]*ConfigWeaponLevel)
	for _, v := range ConfigWeaponLevelSlice {
		levelMap, ok := ConfigWeaponLevelMap[v.WeaponStar]
		if !ok {
			levelMap = make(map[int]*ConfigWeaponLevel)
			ConfigWeaponLevelMap[v.WeaponStar] = levelMap
		}
		levelMap[v.Level] = v
	}
	return
}

func MakeConfigWeaponStarMap() {
	ConfigWeaponStarMap = make(map[int]map[int]*ConfigWeaponStar)
	for _, v := range ConfigWeaponStarSlice {
		starMap, ok := ConfigWeaponStarMap[v.WeaponStar]
		if !ok {
			starMap = make(map[int]*ConfigWeaponStar)
			ConfigWeaponStarMap[v.WeaponStar] = starMap
		}
		starMap[v.StarLevel] = v
	}
	return
}

// RandDropTest 随机掉落测试
func RandDropTest() {
	dropGroup := ConfigDropGroupMap[1000]
	if dropGroup == nil {
		return
	}
	num := 0
	for {
		config := GetRandDropNew(dropGroup)
		if config.IsEnd == LoginTrue {
			fmt.Println(GetItemName(config.Result))
			num++
			dropGroup = ConfigDropGroupMap[1000]
			if num >= 100 {
				break
			} else {
				continue
			}
		}
		dropGroup = ConfigDropGroupMap[config.Result]
		if dropGroup == nil {
			break
		}
	}
}

// GetRandDrop 获得随机掉落物，返回掉落物配置信息
func GetRandDrop(dropGroup *DropGroup) *ConfigDrop {
	//拿到总的权重，生成0-总权重之间的一个数
	randNum := rand.Intn(dropGroup.WeightAll)
	// 获取随机数
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		// 权重依次相加
		randNow += v.Weight
		if randNow > randNum {
			return v
		}
	}
	return nil
}

// GetRandDropNew 使用递归
func GetRandDropNew(dropGroup *DropGroup) *ConfigDrop {
	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			if v.IsEnd == LoginTrue {
				// 结束的话直接返回
				return v
			}
			dropGroup := ConfigDropGroupMap[v.Result]
			// 增加容错
			if dropGroup == nil {
				return nil
			}
			// 将参数刷新为递归后的参数
			return GetRandDropNew(dropGroup)
		}
	}
	return nil
}

func GetRandDropNew1(dropGroup *DropGroup, fiveInfo map[int]int, fourInfo map[int]int) *ConfigDrop {
	for _, v := range dropGroup.DropConfigs {
		_, ok := fiveInfo[v.Result]
		if ok {
			index := 0
			maxGetTime := 0
			for k, config := range dropGroup.DropConfigs {
				_, nowOK := fiveInfo[config.Result]
				if !nowOK {
					continue
				}
				if maxGetTime < fiveInfo[config.Result] {
					maxGetTime = fiveInfo[config.Result]
					index = k
				}
			}
			return dropGroup.DropConfigs[index]
		}

		_, ok = fourInfo[v.Result]
		if ok {
			index := 0
			maxGetTime := 0
			for k, config := range dropGroup.DropConfigs {
				_, nowOK := fourInfo[config.Result]
				if !nowOK {
					continue
				}
				if maxGetTime < fourInfo[config.Result] {
					maxGetTime = fourInfo[config.Result]
					index = k
				}
			}
			return dropGroup.DropConfigs[index]
		}
	}

	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			if v.IsEnd == LoginTrue {
				return v
			}
			dropGroup := ConfigDropGroupMap[v.Result]
			if dropGroup == nil {
				return nil
			}
			return GetRandDropNew1(dropGroup, fiveInfo, fourInfo)
		}
	}
	return nil
}

func GetRandDropNew2(dropGroup *DropGroup, fiveInfo map[int]int, fourInfo map[int]int) *ConfigDrop {
	for _, v := range dropGroup.DropConfigs {
		_, ok := fiveInfo[v.Result]
		if ok {
			index := 0
			minGetTime := 0
			for k, config := range dropGroup.DropConfigs {
				_, nowOK := fiveInfo[config.Result]
				if !nowOK {
					index = k
					break
				}
				if minGetTime == 0 || minGetTime > fiveInfo[config.Result] {
					minGetTime = fiveInfo[config.Result]
					index = k
				}
			}
			return dropGroup.DropConfigs[index]
		}

		_, ok = fourInfo[v.Result]
		if ok {
			index := 0
			minGetTime := 0
			for k, config := range dropGroup.DropConfigs {
				_, nowOK := fourInfo[config.Result]
				if !nowOK {
					index = k
					break
				}
				if minGetTime == 0 || minGetTime > fourInfo[config.Result] {
					minGetTime = fourInfo[config.Result]
					index = k
				}
			}
			return dropGroup.DropConfigs[index]
		}
	}

	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			if v.IsEnd == LoginTrue {
				return v
			}
			dropGroup := ConfigDropGroupMap[v.Result]
			if dropGroup == nil {
				return nil
			}
			return GetRandDropNew2(dropGroup, fiveInfo, fourInfo)
		}
	}
	return nil
}

func GetDropItemGroup(dropId int) *DropItemGroup {
	return ConfigDropItemGroupMap[dropId]
}

func GetDropItemGroupNew(dropId int) []*ConfigDropItem {
	rel := make([]*ConfigDropItem, 0)
	if dropId == 0 {
		return rel
	}
	config := GetDropItemGroup(dropId)
	configsAll := make([]*ConfigDropItem, 0)
	for _, v := range config.DropConfigs {
		if v.DropType == DropItemTypeItem {
			rel = append(rel, v)
		} else if v.DropType == DropItemTypeGroup {
			randNum := rand.Intn(PercentAll)
			if randNum < v.Weight {
				configs := GetDropItemGroupNew(v.ItemId)
				rel = append(rel, configs...)
			}
		} else if v.DropType == DropItemTypeWeight {
			configsAll = append(configsAll, v)
		}
	}
	if len(configsAll) > 0 {
		allRate := 0
		for _, v := range configsAll {
			allRate += v.Weight
		}
		randNum := rand.Intn(allRate)
		nowRate := 0
		for _, v := range configsAll {
			nowRate += v.Weight
			if nowRate > randNum {
				newConfig := new(ConfigDropItem)
				newConfig.Weight = PercentAll
				newConfig.DropId = v.DropId
				newConfig.DropType = v.DropType
				newConfig.ItemId = v.ItemId
				newConfig.ItemNumMin = v.ItemNumMin
				newConfig.ItemNumMax = v.ItemNumMax
				newConfig.WorldAdd = v.WorldAdd
				rel = append(rel, newConfig)
				break
			}
		}
	}
	return rel
}

// GetStatueConfig 获取神像配置
func GetStatueConfig(statueId int, level int) *ConfigStatue {
	_, ok := ConfigStatueMap[statueId]
	if !ok {
		return nil
	}

	_, ok = ConfigStatueMap[statueId][level]
	if !ok {
		return nil
	}
	return ConfigStatueMap[statueId][level]
}

func GetRelicsLevelConfig(mainEntry int, level int) *ConfigRelicsLevel {
	_, ok := ConfigRelicsLevelMap[mainEntry]
	if !ok {
		return nil
	}

	_, ok = ConfigRelicsLevelMap[mainEntry][level]
	if !ok {
		return nil
	}
	return ConfigRelicsLevelMap[mainEntry][level]
}

func GetWeaponLevelConfig(weaponStar int, level int) *ConfigWeaponLevel {
	_, ok := ConfigWeaponLevelMap[weaponStar]
	if !ok {
		return nil
	}

	_, ok = ConfigWeaponLevelMap[weaponStar][level]
	if !ok {
		return nil
	}
	return ConfigWeaponLevelMap[weaponStar][level]
}

// GetWeaponStarConfig 获取武器起始配置
func GetWeaponStarConfig(weaponStar int, starLevel int) *ConfigWeaponStar {
	_, ok := ConfigWeaponStarMap[weaponStar]
	if !ok {
		return nil
	}

	_, ok = ConfigWeaponStarMap[weaponStar][starLevel]
	if !ok {
		return nil
	}
	return ConfigWeaponStarMap[weaponStar][starLevel]
}
