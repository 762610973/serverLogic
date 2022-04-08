package main

import (
	"fmt"
	"serverLogic/server/src/game"
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
	fmt.Println("数据测试----start")
	player := game.NewTestPlayer()
	player.ReceiveSetIcon(1) //胡桃
	player.ReceiveSetIcon(2) //温迪
	player.ReceiveSetIcon(3) //钟离

	player.ReceiveSetCard(11) //胡桃
	player.ReceiveSetCard(22) //温迪
	player.ReceiveSetCard(33) //钟离

}
