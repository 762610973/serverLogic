package csvs

import (
	"serverLogic/server/src/utils"
)

type ConfigIcon struct {
	IconId int `json:"IconId"`
	Check  int `json:"Check"`
}

var (
	ConfigIconMap map[int]*ConfigIcon
)

func init() {
	ConfigIconMap = make(map[int]*ConfigIcon)
	utils.GetCsvUtilMgr().LoadCsv("Icon", &ConfigIconMap)
	return
}
