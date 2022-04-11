package main

import (
	"fmt"
	"serverLogic/server/src/csvs"
	"serverLogic/server/src/game"
	"time"
)

func main() {
	//当前模块：基础信息
	/*	1.UID
		2.头像、名片
		3.签名
		4.名字
		5.冒险等级 冒险阅历
		6.世界等级 冷却时间
		7.生日
		8.展示阵容 展示名片*/

	//公共管理类，每个类都调用一个线程
	//每个玩家都是一个线程
	csvs.CheckLoadCsv()
	fmt.Println("数据测试----start")
	//需要进行服务器的配置
	go game.GetManageBanWord().Run()
	//这里先获得一个基本的违禁词表，然后再调用run函数，run函数会调用读取配置表的函数，读取到完整的配置表
	playerGM := game.NewTestPlayer()
	//设置一个定时器
	ticker := time.NewTicker(time.Second * 1)

	for {
		select {
		case <-ticker.C:
			playerGM.ModPlayer.AddExp(5000, playerGM)
		}
	}
}
