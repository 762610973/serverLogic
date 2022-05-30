package testMod

import "serverLogic/server/src/game"

func bag() {
	playerGM := game.NewTestPlayer()
	playerGM.ModBag.RemoveItemToBagGM(1000003, 1000)
	playerGM.ModBag.RemoveItemToBagGM(1000003, 1000)
	playerGM.ModBag.AddItemToBag(1000003, 100)
	playerGM.ModBag.AddItemToBag(1000003, 500)
	playerGM.ModBag.RemoveItemToBagGM(1000003, 1000)
	playerGM.ModBag.RemoveItemToBagGM(1000003, 1000)

	//playerGM.ModPlayer.SetCard(4000001, playerGM)
	//playerGM.ModBag.AddItem(4000001, playerGM)
	//playerGM.ModPlayer.SetCard(4000001, playerGM)
}
