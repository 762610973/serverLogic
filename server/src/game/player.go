package game

//每一个玩家就是一个结构,目前是包含两个模块
//客户端直接和player模块交互，player收到消息。调用其他模块处理
//receive函数，跟客户端打交道的函数
//增加经验的接口一定是内部接口

type Player struct {
	ModPlayer *ModPlayer //基础模块
	ModIcon   *ModIcon   //头像模块
	ModCard   *ModCard   //名片模块
}

// NewTestPlayer 生成玩家
// NewTestPlayer :new一个测试玩家
func NewTestPlayer() *Player {
	//模块初始化
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModCard = new(ModCard)

	//以上是模块初始化，下面是数据初始化
	player.ModPlayer.Icon = 0

	return player
}

// ReceiveSetIcon ... 对外接口
// 判断这个玩家是否含有这个icon，如果有就赋值，没有就返回

func (self *Player) ReceiveSetIcon(iconId int) {
	self.ModPlayer.SetIcon(iconId, self)
}
func (self *Player) ReceiveSetCard(cardId int) {
	self.ModPlayer.SetCard(cardId, self)
}
