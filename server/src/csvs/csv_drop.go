package csvs

import "serverLogic/server/src/utils"

type ConfigDrop struct {
	DropId int `json:"DropId"`
	Weight int `json:"Weight"` //权重
	Result int `json:"Result"`
	IsEnd  int `json:"IsEnd"`
}

var (
	ConfigDropSlice []*ConfigDrop
)

func init() {
	ConfigDropSlice = make([]*ConfigDrop, 0)
	utils.GetCsvUtilMgr().LoadCsv("Drop", &ConfigDropSlice)
	return
}
