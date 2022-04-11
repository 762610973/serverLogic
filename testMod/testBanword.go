package testMod

import (
	"serverLogic/server/src/game"
	"time"
)

func testM() {
	go game.GetManageBanWord().Run()
	//这里先获得一个基本的违禁词表，然后再调用run函数，run函数会调用读取配置表的函数，读取到完整的配置表
	player := game.NewTestPlayer()
	//设置一个定时器
	ticker := time.NewTicker(time.Second * 1)

	for {
		select {
		case <-ticker.C:
			if time.Now().Unix()%3 == 0 {
				player.ReceiveSetName("专业代练") //预期目标是拦在外面

			} else if time.Now().Unix()%5 == 0 {
				player.ReceiveSetName("正常玩家")
			}
		}
	}
}
