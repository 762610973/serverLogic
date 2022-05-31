package csvs

import "serverLogic/server/src/utils"

type ConfigMap struct {
	MapId   int    `json:"MapId"`   //地图id
	MapName string `json:"MapName"` //地图名称
}

// ConfigMapEvent 地图事件配置
type ConfigMapEvent struct {
	EventId     int    `json:"EventId"`
	EventType   int    `json:"EventType"`
	Name        string `json:"Name"`
	RefreshType int    `json:"RefreshType"`
	EventDrop   int    `json:"EventDrop"`
	MapId       int    `json:"MapId"`
}

var (
	ConfigMapMap      map[int]*ConfigMap
	ConfigMapEventMap map[int]*ConfigMapEvent
)

func init() {
	ConfigMapMap = make(map[int]*ConfigMap)
	utils.GetCsvUtilMgr().LoadCsv("Map", &ConfigMapMap)

	ConfigMapEventMap = make(map[int]*ConfigMapEvent)
	utils.GetCsvUtilMgr().LoadCsv("MapEvent", &ConfigMapEventMap)
	return
}

// GetMapName 获取地图名称
func GetMapName(mapId int) string {
	_, ok := ConfigMapMap[mapId]
	if !ok {
		return ""
	}
	return ConfigMapMap[mapId].MapName
}

// GetEventName 获取事件名称
func GetEventName(eventId int) string {
	_, ok := ConfigMapEventMap[eventId]
	if !ok {
		return ""
	}
	return ConfigMapEventMap[eventId].Name
}

// GetEventConfig 获取事件配置
func GetEventConfig(eventId int) *ConfigMapEvent {
	return ConfigMapEventMap[eventId]
}
