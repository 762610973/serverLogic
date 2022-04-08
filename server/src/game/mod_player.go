package game

import "fmt"

//每一个模块都对应数据库中的一张表
//玩家模块，包含以下的字段

type ModPlayer struct {
	UserId         int
	Icon           int //头像
	Card           int //名片
	Name           string
	Sign           string
	PlayerLevel    int   //等级
	Exp            int   //经验
	WorldLevel     int   //世界等级
	WorldLevelCool int64 //世界等级冷却时间
	Birth          int   //生日，只能改一次
	ShowTeam       []int //展示阵容
	ShowCard       int   //展示名片
	//看不见的字段
	IsProhibit int //是否封禁
	IsGm       int //是否是游戏管理员
}

func (self *ModPlayer) SetIcon(iconId int, player *Player) {
	//自己的icon模块没有这个头像
	if !player.ModIcon.IsHasIcon(iconId) {
		//没有这个icon，通知客户端，操作非法
		//公司有一个完整的消息通知机制
		return
	}
	player.ModPlayer.Icon = iconId
	fmt.Println("当前图标:", player.ModPlayer.Icon)
}
func (self *ModPlayer) SetCard(cardId int, player *Player) {
	if !player.ModCard.IsHasCard(cardId) {
		return
	}
	player.ModPlayer.Card = cardId
	fmt.Println("当前名片:", player.ModPlayer.Card)
}
