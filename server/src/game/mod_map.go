package game

import (
	"fmt"
	"math/rand"
	"serverLogic/server/src/csvs"
	"time"
)

type Map struct {
	MapId     int            //地图id
	EventInfo map[int]*Event //事件信息
}

// Event 事件
type Event struct {
	EventId       int   //事件id
	State         int   //转态
	NextResetTime int64 //下一次重置时间
}

// StatueInfo 神像信息
type StatueInfo struct {
	StatueId int
	Level    int
	ItemInfo map[int]*ItemInfo
}

// ModMap 地图模块
type ModMap struct {
	MapInfo map[int]*Map
	Statue  map[int]*StatueInfo // 神像信息
}

func (m *ModMap) InitData() {
	m.MapInfo = make(map[int]*Map)
	m.Statue = make(map[int]*StatueInfo)

	// 目前只有一张地图
	for _, v := range csvs.ConfigMapMap {
		// 识别哪些地图可用
		_, ok := m.MapInfo[v.MapId]
		if !ok {
			//不存在的话先把蒙得地图设置出来
			m.MapInfo[v.MapId] = m.NewMapInfo(v.MapId)
		}
	}

	for _, v := range csvs.ConfigMapEventMap {
		_, ok := m.MapInfo[v.MapId]
		if !ok {
			continue
		}
		_, ok = m.MapInfo[v.MapId].EventInfo[v.EventId]
		if !ok {
			m.MapInfo[v.MapId].EventInfo[v.EventId] = new(Event)
			m.MapInfo[v.MapId].EventInfo[v.EventId].EventId = v.EventId
			m.MapInfo[v.MapId].EventInfo[v.EventId].State = csvs.EventStart
		}
	}
}

func (m *ModMap) NewMapInfo(mapId int) *Map {
	mapInfo := new(Map)
	mapInfo.MapId = mapId
	mapInfo.EventInfo = make(map[int]*Event)
	return mapInfo
}

// GetEventList 获取事件列表
func (m *ModMap) GetEventList(config *csvs.ConfigMap) {
	_, ok := m.MapInfo[config.MapId]
	if !ok {
		return
	}
	for _, v := range m.MapInfo[config.MapId].EventInfo {
		// 检查一下刷新
		m.CheckRefresh(v)
		lastTime := v.NextResetTime - time.Now().Unix()
		noticeTime := ""
		if lastTime <= 0 {
			noticeTime = "已刷新"
		} else {
			noticeTime = fmt.Sprintf("%d秒后刷新", lastTime)
		}
		fmt.Println(fmt.Sprintf("事件Id:%d,名字:%s,状态:%d,%s", v.EventId, csvs.GetEventName(v.EventId), v.State, noticeTime))
	}
}

// SetEventState 设置状态
func (m *ModMap) SetEventState(mapId int, eventId int, state int, player *Player) {
	_, ok := m.MapInfo[mapId]
	if !ok {
		fmt.Println("地图不存在")
		return
	}
	_, ok = m.MapInfo[mapId].EventInfo[eventId]
	if !ok {
		fmt.Println("事件不存在")
		return
	}
	if m.MapInfo[mapId].EventInfo[eventId].State >= state {
		fmt.Println("状态异常")
		return
	}
	eventConfig := csvs.GetEventConfig(m.MapInfo[mapId].EventInfo[eventId].EventId)
	if eventConfig == nil {
		return
	}
	configMap := csvs.ConfigMapMap[mapId]
	if configMap == nil {
		return
	}
	if !player.ModBag.HasEnoughItem(eventConfig.CostItem, eventConfig.CostNum) {
		fmt.Println(fmt.Sprintf("%s不足!", csvs.GetItemName(eventConfig.CostItem)))
		return
	}
	if configMap.MapType == csvs.RefreshPlayer && eventConfig.EventType == csvs.EventTypeReward {
		for _, v := range m.MapInfo[mapId].EventInfo {
			eventConfigNow := csvs.GetEventConfig(v.EventId)
			if eventConfigNow == nil {
				continue
			}
			if eventConfigNow.EventType != csvs.EventTypeNormal {
				continue
			}
			if v.EventId == eventId {
				continue
			}
			// 遍历事件，发现任何一个没完成就返回就返回
			if v.State != csvs.EventEnd {
				fmt.Println("有事件尚未完成:", v.EventId)
				return
			}
		}
	}

	m.MapInfo[mapId].EventInfo[eventId].State = state
	if state == csvs.EventFinish {
		fmt.Println("事件完成")
	}
	if state == csvs.EventEnd {
		for i := 0; i < eventConfig.EventDropTimes; i++ {
			config := csvs.GetDropItemGroupNew(eventConfig.EventDrop)
			for _, v := range config {
				randNum := rand.Intn(csvs.PercentAll)
				if randNum < v.Weight {
					randAll := v.ItemNumMax - v.ItemNumMin + 1
					itemNum := rand.Intn(randAll) + v.ItemNumMin
					worldLevel := player.ModPlayer.GetWorldLevelNow()
					if worldLevel > 0 {
						itemNum = itemNum * (csvs.PercentAll + worldLevel*v.WorldAdd) / csvs.PercentAll
					}
					player.ModBag.AddItem(v.ItemId, int64(itemNum), player)
				}
			}
		}
		fmt.Println("事件领取")
	}
	if state > 0 {
		switch eventConfig.RefreshType {
		case csvs.MapRefreshSelf:
			m.MapInfo[mapId].EventInfo[eventId].NextResetTime = time.Now().Unix() + csvs.MapRefreshSelfTime
		}
	}
}

// RefreshDay 日刷新
func (m *ModMap) RefreshDay() {
	for _, v := range m.MapInfo {
		for _, v := range m.MapInfo[v.MapId].EventInfo {
			config := csvs.ConfigMapEventMap[v.EventId]
			if config == nil {
				continue
			}
			// 不满足刷新条件
			if config.RefreshType != csvs.MapRefreshDay {
				continue
			}
			v.State = csvs.EventStart
		}
	}
}

// RefreshWeek 周刷新
func (m *ModMap) RefreshWeek() {
	for _, v := range m.MapInfo {
		for _, v := range m.MapInfo[v.MapId].EventInfo {
			config := csvs.ConfigMapEventMap[v.EventId]
			if config == nil {
				continue
			}
			if config.RefreshType != csvs.MapRefreshWeek {
				continue
			}
			v.State = csvs.EventStart
		}
	}
}

// RefreshSelf 自刷新
func (m *ModMap) RefreshSelf() {
	for _, v := range m.MapInfo {
		for _, v := range m.MapInfo[v.MapId].EventInfo {
			config := csvs.ConfigMapEventMap[v.EventId]
			if config == nil {
				continue
			}
			if config.RefreshType != csvs.MapRefreshSelf {
				continue
			}
			if time.Now().Unix() <= v.NextResetTime {
				v.State = csvs.EventStart
			}
		}
	}
}

// CheckRefresh 检查是否刷新
func (m *ModMap) CheckRefresh(event *Event) {
	if event.NextResetTime > time.Now().Unix() {
		return
	}
	eventConfig := csvs.GetEventConfig(event.EventId)
	if eventConfig == nil {
		return
	}
	// 置为初始状态
	event.State = csvs.EventStart
	// 判断刷新类型
	switch eventConfig.RefreshType {
	case csvs.MapRefreshDay:
		count := time.Now().Unix() / csvs.MapRefreshDayTime
		count++
		event.NextResetTime = count * csvs.MapRefreshDayTime
	case csvs.MapRefreshWeek:
		count := time.Now().Unix() / csvs.MapRefreshWeekTime
		count++
		event.NextResetTime = count * csvs.MapRefreshWeekTime
	case csvs.MapRefreshSelf:
	case csvs.MapRefreshCant: // 识别类型4，类型4是不可以刷新的
		return
	}
	event.State = csvs.EventStart
}

// NewStatue 生成神像
func (m *ModMap) NewStatue(statueId int) *StatueInfo {
	data := new(StatueInfo)
	data.StatueId = statueId
	data.Level = 0
	data.ItemInfo = make(map[int]*ItemInfo)
	return data
}

// UpStatue 升级神像
func (m *ModMap) UpStatue(stateId int, player *Player) {
	_, ok := m.Statue[stateId]
	if !ok {
		// 不存在的话生成神像
		m.Statue[stateId] = m.NewStatue(stateId)
	}
	info, ok := m.Statue[stateId]
	if !ok {
		return
	}
	nextLevel := info.Level + 1 //下一级
	// 获取下一级的配置
	nextConfig := csvs.GetStatueConfig(stateId, nextLevel)
	if nextConfig == nil {
		return
	}
	// 调用背包模块，判断消耗的东西背包中是否充足
	_, okNow := info.ItemInfo[nextConfig.CostItem]
	nowNum := int64(0)
	if okNow {
		nowNum = info.ItemInfo[nextConfig.CostItem].ItemNum
	}
	needNum := nextConfig.CostNum - nowNum
	if !player.ModBag.HasEnoughItem(nextConfig.CostItem, needNum) {
		num := player.ModBag.GetItemNum(nextConfig.CostItem, player)
		// 判断是否充足
		if num <= 0 {
			fmt.Println(fmt.Sprintf("神像升级物品不足"))
			return
		}

		_, okItem := info.ItemInfo[nextConfig.CostItem]
		if !okItem {
			info.ItemInfo[nextConfig.CostItem] = new(ItemInfo)
			info.ItemInfo[nextConfig.CostItem].ItemId = nextConfig.CostItem
			info.ItemInfo[nextConfig.CostItem].ItemNum = 0
		}
		_, okItem = info.ItemInfo[nextConfig.CostItem]
		if !okItem {
			return
		}
		info.ItemInfo[nextConfig.CostItem].ItemNum += num
		player.ModBag.RemoveItemToBag(nextConfig.CostItem, num, player)
		fmt.Println(fmt.Sprintf("神像升级,提交物品%d，数量%d，当前数量%d", nextConfig.CostItem, num, info.ItemInfo[nextConfig.CostItem].ItemNum))

	} else {
		player.ModBag.RemoveItemToBag(nextConfig.CostItem, needNum, player)
		info.Level++
		info.ItemInfo = make(map[int]*ItemInfo)
		fmt.Println(fmt.Sprintf("神像升级成功,神像:%d，当前等级:%d", info.StatueId, info.Level))
	}
}

func (m *ModMap) RefreshByPlayer(mapId int) {
	config := csvs.ConfigMapMap[mapId]
	if config == nil {
		return
	}
	if config.MapType != csvs.RefreshPlayer {
		return
	}
	_, ok := m.MapInfo[config.MapId]
	if !ok {
		return
	}
	for _, v := range m.MapInfo[config.MapId].EventInfo {
		v.State = csvs.EventStart
		//置位起始状态
	}
}
