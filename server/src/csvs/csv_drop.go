package csvs

import "serverLogic/server/src/utils"

type ConfigDrop struct {
	DropId int `json:"DropId"`
	Weight int `json:"Weight"`
	Result int `json:"Result"`
	IsEnd  int `json:"IsEnd"`
}

type ConfigDropItem struct {
	DropId     int `json:"DropId"`
	DropType   int `json:"DropType"` //类型
	Weight     int `json:"Weight"`   //权重
	ItemId     int `json:"ItemId"`
	ItemNumMin int `json:"ItemNumMin"`
	ItemNumMax int `json:"ItemNumMax"`
	WorldAdd   int `json:"WorldAdd"` //大世界等级加成
}

var (
	ConfigDropSlice     []*ConfigDrop
	ConfigDropItemSlice []*ConfigDropItem
)

func init() {
	ConfigDropSlice = make([]*ConfigDrop, 0)
	utils.GetCsvUtilMgr().LoadCsv("Drop", &ConfigDropSlice)

	ConfigDropItemSlice = make([]*ConfigDropItem, 0)
	utils.GetCsvUtilMgr().LoadCsv("DropItem", &ConfigDropItemSlice)
	return
}
