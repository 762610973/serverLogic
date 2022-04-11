package csvs

import (
	"serverLogic/server/src/utils"
)

type ConfigPlayerLevel struct {
	PlayerLevel int `json:"PlayerLevel"`
	PlayerExp   int `json:"PlayerExp"`
	WorldLevel  int `json:"WorldLevel"`
	ChapterId   int `json:"ChapterId"`
}

var (
	// ConfigPlayerLevelSlice 配置表中的每一行对应切片中的一个元素
	ConfigPlayerLevelSlice []*ConfigPlayerLevel
)

func init() {
	//加载csv文件
	//这里只需要传入文件名就行，后缀会在utils配置包中保存
	utils.GetCsvUtilMgr().LoadCsv("PlayerLevel", &ConfigPlayerLevelSlice)
	return
}

// GetNowLevelConfig 获取当前等级配置,传入的是等级
func GetNowLevelConfig(level int) *ConfigPlayerLevel {
	//如果传入进来的是非法值，
	if level <= 0 || level > len(ConfigPlayerLevelSlice) {
		return nil
	}
	return ConfigPlayerLevelSlice[level-1]
}
