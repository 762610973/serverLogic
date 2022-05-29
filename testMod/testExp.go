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
func playerSet(player *game.Player) {
	startTime := time.Now().Nanosecond()
	for i := 0; i < 100000; i++ {
		player.ModUniqueTask.Locker.Lock()
		player.ModUniqueTask.MyTaskInfo[10001] = new(game.TaskInfo)
		player.ModUniqueTask.Locker.Unlock()
	}
	endTime := time.Now().Nanosecond() - startTime
	fmt.Println(endTime / 1000000)
}

func playerGet(player *game.Player) {
	for i := 0; i < 1000000; i++ {
		player.ModUniqueTask.Locker.RLock()
		_, ok := player.ModUniqueTask.MyTaskInfo[10001]
		if ok {

		}
		player.ModUniqueTask.Locker.RUnlock()
	}
}
