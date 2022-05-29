package testMod

import "serverLogic/server/src/game"

func testSetShow() {
	playerGM := game.NewTestPlayer()
	playerGM.ModPlayer.SetShowCard([]int{1001, 1001, 1001, 1002, 1005}, playerGM)
	playerGM.ModPlayer.SetShowCard([]int{}, playerGM)
	playerGM.ModPlayer.SetShowCard([]int{1009}, playerGM)
	playerGM.ModPlayer.SetShowTeam([]int{1009, 1002, 1002, 10021002, 1002, 1002, 1002, 1002, 1002, 1002, 1002, 1002, 100, 1002}, playerGM)
}
