package game

import (
	"fmt"
	"regexp"
	"serverLogic/server/src/csvs"
	"time"
)

var manageBanWord *ManageBanWord

// ManageBanWord 管理模块
type ManageBanWord struct {
	BanWordBase  []string //从配置表中读取违禁词语，是一个服务器线程
	BanWordExtra []string //更新
}

// GetManageBanWord 应用单例模式，任何时候都用的是同一个，节省空间，方便管理
func GetManageBanWord() *ManageBanWord {
	// 不存在就创建，存在就复用
	if manageBanWord == nil {
		manageBanWord = new(ManageBanWord)
		manageBanWord.BanWordBase = []string{"外挂", "工具", "脚本"}
		manageBanWord.BanWordExtra = []string{"原神", "股票", "刷单"}
	}
	return manageBanWord
}

// IsBanWord 判断是否是违禁词
func (m *ManageBanWord) IsBanWord(txt string) bool {
	//
	for _, v := range m.BanWordBase {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			//包含返回ture
			return match
		}
	}
	for _, v := range m.BanWordExtra {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			//包含返回ture
			return match
		}
	}
	return false
}

// 定时器

func (m *ManageBanWord) Run() {
	//这里获取违禁词汇
	m.BanWordBase = csvs.GetBanWordBase()
	//fmt.Println(m.BanWordBase)
	//基础词库的更新
	//服务器启动的时候就会调用

	// 定时器，根据时间做出行为
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			// 每十秒更新一次
			if time.Now().Unix()%10 == 0 {
				fmt.Println("\n更新词库")
			} else {
				//fmt.Println("待机")
			}
		}
	}
}
