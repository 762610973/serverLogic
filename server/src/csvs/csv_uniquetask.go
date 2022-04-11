package csvs

import "serverLogic/server/src/utils"

//突破任务表的加载

// ConfigUniqueTask 定义一个结构体，对应表中的各个字段
type ConfigUniqueTask struct {
	TaskId    int `json:"TaskId"`    //突破任务的id
	SortType  int `json:"SortType"`  //突破任务的分类
	OpenLevel int `json:"OpenLevel"` //开启等级
	TaskType  int `json:"TaskType"`  //这个任务需要什么条件
	Condition int `json:"Condition"` //
}

var (
	ConfigUniqueTaskMap map[int]*ConfigUniqueTask
)

func init() {
	ConfigUniqueTaskMap = make(map[int]*ConfigUniqueTask)
	utils.GetCsvUtilMgr().LoadCsv("UniqueTask", &ConfigUniqueTaskMap)
	return
}
