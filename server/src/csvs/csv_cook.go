package csvs

import "serverLogic/server/src/utils"

type ConfigCook struct {
	CookId int `json:"CookId"`
}

var (
	ConfigCookMap map[int]*ConfigCook
)

func init() {
	ConfigCookMap = make(map[int]*ConfigCook)
	utils.GetCsvUtilMgr().LoadCsv("Cook", &ConfigCookMap)
	return
}

// GetCookConfig 获得烹饪配置map
func GetCookConfig(cookId int) *ConfigCook {
	return ConfigCookMap[cookId]
}
