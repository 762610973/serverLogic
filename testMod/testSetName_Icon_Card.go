package testMod

import "serverLogic/server/src/game"

func testS() {
	//头像、卡片、id、name测试
	player := game.NewTestPlayer()
	player.ReceiveSetIcon(1) //胡桃
	player.ReceiveSetIcon(2) //温迪
	player.ReceiveSetIcon(3) //钟离

	player.ReceiveSetCard(11) //胡桃
	player.ReceiveSetCard(22) //温迪
	player.ReceiveSetCard(33) //钟离

	player.ReceiveSetName("好人")
	player.ReceiveSetName("坏人")
	player.ReceiveSetName("求外挂")
	player.ReceiveSetName("感觉不如原神。。画质")
}
