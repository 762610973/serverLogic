package game

import (
	"fmt"
	"regexp"
	"time"
)

var manageBanWord *ManageBanWord

type ManageBanWord struct {
	BanWordBase  []string //从配置表中读取违禁词语，是一个服务器线程
	BanWordExtra []string //更新

}

//单例模式，任何时候都用的是同一个

func GetManageBanWord() *ManageBanWord {
	if manageBanWord == nil {
		manageBanWord = new(ManageBanWord)
		manageBanWord.BanWordBase = []string{"外挂", "工具"}
		manageBanWord.BanWordExtra = []string{"原神"}
	}
	return manageBanWord
}

// IsBanWord 判断是否是违禁词
func (self *ManageBanWord) IsBanWord(txt string) bool {
	for _, v := range self.BanWordBase {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			//包含返回ture
			return match
		}
	}
	for _, v := range self.BanWordExtra {
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

func (self *ManageBanWord) Run() {
	//基础词库的更新
	//服务器启动的时候就会调用
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			if time.Now().Unix()%10 == 0 {
				fmt.Println("更新词库")
			} else {
				fmt.Println("待机")
			}
		}
	}
}
