package testMod

import (
	"fmt"
	"serverLogic/server/src/csvs"
	"serverLogic/server/src/game"
	"time"
)

func testExp() {
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
