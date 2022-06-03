package main

import "serverLogic/server/src/game"

func main() {

	//公共管理类，每个类都调用一个线程 x 1
	//每个玩家都是一个线程 x 1
	game.GetServer().Start()
	return
}
