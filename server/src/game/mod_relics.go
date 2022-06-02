package game

import (
	"fmt"
	"math/rand"
	"serverLogic/server/src/csvs"
)

type Relics struct {
	RelicsId   int
	KeyId      int
	MainEntry  int // 主词条
	Level      int
	Exp        int
	OtherEntry []int
	RoleId     int
}

type ModRelics struct {
	RelicsInfo map[int]*Relics
	MaxKey     int
}

func (m *ModRelics) AddItem(itemId int, num int64) {
	config := csvs.GetRelicsConfig(itemId)
	if config == nil {
		fmt.Println("配置不存在")
		return
	}
	if len(m.RelicsInfo)+int(num) > csvs.RelicsMaxCount {
		fmt.Println("超过最大值")
		return
	}
	for i := int64(0); i < num; i++ {
		relics := m.NewRelics(itemId)
		m.RelicsInfo[relics.KeyId] = relics
		fmt.Println("获得圣遗物:")
		relics.ShowInfo()
	}
}

func (m *ModRelics) NewRelics(itemId int) *Relics {
	relicRel := new(Relics)
	relicRel.RelicsId = itemId
	m.MaxKey++
	relicRel.KeyId = m.MaxKey
	// 拿配置
	config := csvs.ConfigRelicsMap[itemId]
	if config == nil {
		return nil
	}

	relicRel.MainEntry = m.MakeMainEntry(config.MainGroup)
	for i := 0; i < config.OtherGroupNum; i++ {
		if i == config.OtherGroupNum-1 {
			// 最后一次走概率
			randNum := rand.Intn(csvs.PercentAll)
			if randNum < csvs.AllEntryRate {
				relicRel.OtherEntry = append(relicRel.OtherEntry, m.MakeOtherEntry(relicRel, config.OtherGroup))
			}
		} else {
			relicRel.OtherEntry = append(relicRel.OtherEntry, m.MakeOtherEntry(relicRel, config.OtherGroup))
		}
	}
	return relicRel
}

// MakeMainEntry 生成主词条
func (m *ModRelics) MakeMainEntry(mainGroup int) int {
	configs, ok := csvs.ConfigRelicsEntryGroupMap[mainGroup]
	if !ok {
		return 0
	}
	allRate := 0
	// 获取总权值
	for _, v := range configs {
		allRate += v.Weight
	}
	randNum := rand.Intn(allRate)
	nowNum := 0
	for _, v := range configs {
		nowNum += v.Weight
		if nowNum > randNum {
			return v.Id
		}
	}
	return 0
}

// MakeOtherEntry 生成副词条
func (m *ModRelics) MakeOtherEntry(relics *Relics, otherGroup int) int {
	configs, ok := csvs.ConfigRelicsEntryGroupMap[otherGroup]
	if !ok {
		return 0
	}
	configNow := csvs.GetRelicsConfig(relics.RelicsId)
	if configNow == nil {
		return 0
	}
	if len(relics.OtherEntry) >= configNow.OtherGroupNum {
		allEntry := make(map[int]int)
		for _, id := range relics.OtherEntry {
			otherConfig, _ := csvs.ConfigRelicsEntryMap[id]
			if otherConfig != nil {
				allEntry[otherConfig.AttrType] = csvs.LoginTrue
			}
		}

		allRate := 0
		for _, v := range configs {
			_, ok := allEntry[v.AttrType]
			if !ok {
				continue
			}
			allRate += v.Weight
		}
		randNum := rand.Intn(allRate)
		nowNum := 0
		for _, v := range configs {
			_, ok := allEntry[v.AttrType]
			if !ok {
				continue
			}
			nowNum += v.Weight
			if nowNum > randNum {
				return v.Id
			}
		}
	} else {
		// 做排重处理
		allEntry := make(map[int]int)
		mainConfig, _ := csvs.ConfigRelicsEntryMap[relics.MainEntry]
		if mainConfig != nil {
			allEntry[mainConfig.AttrType] = csvs.LoginTrue
		}
		for _, id := range relics.OtherEntry {
			otherConfig, _ := csvs.ConfigRelicsEntryMap[id]
			if otherConfig != nil {
				allEntry[otherConfig.AttrType] = csvs.LoginTrue
			}
		}

		allRate := 0
		for _, v := range configs {
			_, ok := allEntry[v.AttrType]
			if ok {
				continue
			}
			allRate += v.Weight
		}
		randNum := rand.Intn(allRate)
		nowNum := 0
		for _, v := range configs {
			_, ok := allEntry[v.AttrType]
			if ok {
				continue
			}
			nowNum += v.Weight
			if nowNum > randNum {
				return v.Id
			}
		}
	}
	return 0
}

func (r *Relics) ShowInfo() {
	fmt.Println(fmt.Sprintf("key:%d,Id:%d", r.KeyId, r.RelicsId))
	fmt.Println(fmt.Sprintf("当前等级:%d,当前经验:%d", r.Level, r.Exp))
	mainEntryConfig := csvs.GetRelicsLevelConfig(r.MainEntry, r.Level)
	if mainEntryConfig != nil {
		fmt.Println(fmt.Sprintf("主词条属性:%s,值:%d", mainEntryConfig.AttrName, mainEntryConfig.AttrValue))
	}
	for _, v := range r.OtherEntry {
		otherEntryConfig := csvs.ConfigRelicsEntryMap[v]
		if otherEntryConfig != nil {
			fmt.Println(fmt.Sprintf("副词条属性:%s,值:%d", otherEntryConfig.AttrName, otherEntryConfig.AttrValue))
		}
	}
}

func (m *ModRelics) RelicsUp(player *Player) {
	relics := m.RelicsInfo[1]
	if relics == nil {
		fmt.Println("找不到对应圣遗物")
		return
	}
	relics.Exp += 100000
	for {
		nextLevelConfig := csvs.GetRelicsLevelConfig(relics.MainEntry, relics.Level+1)
		if nextLevelConfig == nil {
			break
		}
		// 小于下一级的经验，break
		if relics.Exp < nextLevelConfig.NeedExp {
			break
		}
		// 经验加一
		relics.Level += 1
		relics.Exp -= nextLevelConfig.NeedExp
		if relics.Level%4 == 0 {
			relicsConfig := csvs.ConfigRelicsMap[relics.RelicsId]
			if relicsConfig != nil {
				relics.OtherEntry = append(relics.OtherEntry, m.MakeOtherEntry(relics, relicsConfig.OtherGroup))
			}
		}
	}

	relics.ShowInfo()
}

// RelicsTop 模拟一个满级圣遗物
func (m *ModRelics) RelicsTop(player *Player) {
	relics := m.NewRelics(7000005)
	relics.Level = 20
	config := csvs.GetRelicsConfig(relics.RelicsId)
	if config == nil {
		return
	}
	for i := 0; i < 5; i++ {
		relics.OtherEntry = append(relics.OtherEntry, m.MakeOtherEntry(relics, config.OtherGroup))
	}
	relics.ShowInfo()
}

// RelicsTestBest 多少次可以出一个极品头
func (m *ModRelics) RelicsTestBest(player *Player) {
	config := csvs.GetRelicsConfig(7000005)
	if config == nil {
		return
	}
	allTimes := 500000
	relicsBestInfo := make([]*Relics, 0)
	for i := 0; i < allTimes; i++ {
		relics := m.NewRelics(7000005)
		relics.Level = 20
		config := csvs.GetRelicsConfig(relics.RelicsId)
		if config == nil {
			return
		}
		for i := 0; i < 5; i++ {
			relics.OtherEntry = append(relics.OtherEntry, m.MakeOtherEntry(relics, config.OtherGroup))
		}

		configMain := csvs.ConfigRelicsEntryMap[relics.MainEntry]
		if configMain == nil {
			continue
		}

		if configMain.AttrType != 4 && configMain.AttrType != 5 {
			continue
		}
		bestEntryCount := 0
		for _, v := range relics.OtherEntry {
			configOther := csvs.ConfigRelicsEntryMap[v]
			if configOther == nil {
				continue
			}
			if configOther.AttrType != 4 && configOther.AttrType != 5 {
				continue
			}
			bestEntryCount++
		}

		if bestEntryCount < 6 {
			continue
		}
		relicsBestInfo = append(relicsBestInfo, relics)
	}
	fmt.Println(fmt.Sprintf("生成了圣遗物头部位%d个，极品数量%d", allTimes, len(relicsBestInfo)))

	for _, v := range relicsBestInfo {
		v.ShowInfo()
	}
}
