package main

import (
	"fmt"
	"math/rand"
	"serverLogic/server/src/csvs"
	"serverLogic/server/src/game"
	"time"
)

func main() {

	//公共管理类，每个类都调用一个线程 x 1
	//每个玩家都是一个线程 x 1
	csvs.CheckLoadCsv()
	fmt.Println("数据测试----start")
	rand.Seed(time.Now().UnixNano())
	//需要进行服务器的配置
	// 启动一个违禁词goroutine

	//go game.GetManageBanWord().Run()
	//这里先获得一个基本的违禁词表，然后再调用run函数，run函数会调用读取配置表的函数，读取到完整的配置表

	playerGM := game.NewTestPlayer()
	go playerGM.Run()
	for {

	}

	//设置一个定时器
	/*	ticker := time.NewTicker(time.Second * 10)

		for {
			select {
			case <-ticker.C:
				playerTest := game.NewTestPlayer()
				go playerTest.Run()
			}
		}*/

}
