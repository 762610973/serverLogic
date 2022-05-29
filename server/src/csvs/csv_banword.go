package csvs

import "fmt"

// ConfigBanWord 读取配置表的信息
type ConfigBanWord struct {
	Id  int
	Txt string
}

// 声明一个全局变量

var ConfigBanWordSlice []*ConfigBanWord

//第一次调用时会进行初始化
func init() {
	//这里将数据写死
	ConfigBanWordSlice = append(ConfigBanWordSlice,
		&ConfigBanWord{Id: 1, Txt: "外挂"},
		&ConfigBanWord{Id: 2, Txt: "辅助"},
		&ConfigBanWord{Id: 3, Txt: "微信"},
		&ConfigBanWord{Id: 4, Txt: "代练"},
		&ConfigBanWord{Id: 5, Txt: "赚钱"},
		&ConfigBanWord{Id: 6, Txt: "脚本"},
	)
	fmt.Println("csv_banWord 初始化")
}

// GetBanWordBase 读取配置表
func GetBanWordBase() []string {
	resString := make([]string, 0)
	for _, v := range ConfigBanWordSlice {
		resString = append(resString, v.Txt)
	}
	return resString
}
