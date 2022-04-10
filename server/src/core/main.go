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

	//公共管理类，每个类都调用一个线程
	//每个玩家都是一个线程
	fmt.Println("数据测试----start")
	/*	头像、卡片、id、name测试
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
		player.ReceiveSetName("感觉不如原神。。画质")*/
	//需要进行服务器的配置
	go game.GetManageBanWord().Run()

}
