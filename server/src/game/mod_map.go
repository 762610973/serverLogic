package game

import (
	"fmt"
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

// ModMap 地图模块
type ModMap struct {
	MapInfo map[int]*Map
}

func (m *ModMap) InitData() {
	m.MapInfo = make(map[int]*Map)
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

// SetEventState 设置地图状态
func (m *ModMap) SetEventState(mapId int, eventId int, state int) {
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
	m.MapInfo[mapId].EventInfo[eventId].State = state
	if state == csvs.EventFinish {
		fmt.Println("事件完成")
	}
	if state == csvs.EventEnd {
		fmt.Println("事件领取")
	}
	if state > 0 {
		eventConfig := csvs.GetEventConfig(m.MapInfo[mapId].EventInfo[eventId].EventId)
		if eventConfig == nil {
			return
		}
		switch eventConfig.RefreshType {
		case csvs.MapRefreshSelf:
			m.MapInfo[mapId].EventInfo[eventId].NextResetTime = time.Now().Unix() + csvs.MapRefreshSelfTime
		}
	}
}

func (m *ModMap) RefreshDay() {
	for _, v := range m.MapInfo {
		for _, v := range m.MapInfo[v.MapId].EventInfo {
			config := csvs.ConfigMapEventMap[v.EventId]
			if config == nil {
				continue
			}
			if config.RefreshType != csvs.MapRefreshDay {
				continue
			}
			v.State = csvs.EventStart
		}
	}
}

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

func (m *ModMap) CheckRefresh(event *Event) {
	if event.NextResetTime > time.Now().Unix() {
		return
	}
	eventConfig := csvs.GetEventConfig(event.EventId)
	if eventConfig == nil {
		return
	}
	event.State = csvs.EventStart
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
	}
}
