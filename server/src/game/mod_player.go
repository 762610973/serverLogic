package game

import (
	"fmt"
	"serverLogic/server/src/csvs"
)

//每一个模块都对应数据库中的一张表
//玩家模块，包含以下的字段

type ModPlayer struct {
	UserId         int
	Icon           int //头像
	Card           int //名片
	Name           string
	Sign           string
	PlayerLevel    int   //等级
	PlayerExp      int   //阅历（经验）
	WorldLevel     int   //世界等级
	WorldLevelNow  int   //大世界当前等级
	WorldLevelCool int64 //世界等级冷却时间
	Birth          int   //生日，只能改一次
	ShowTeam       []int //展示阵容
	ShowCard       int   //展示名片
	HideShowTeam   int   //隐藏开关
	//看不见的字段
	IsProhibit int //是否封禁
	IsGm       int //是否是游戏管理员
}

func (self *ModPlayer) SetIcon(iconId int, player *Player) {
	//自己的icon模块没有这个头像
	if !player.ModIcon.IsHasIcon(iconId) {
		//没有这个icon，通知客户端，操作非法
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

func (self *ModPlayer) SetName(name string, player *Player) {
	/*if !player.ModCard.IsHasCard(name) {
		return
	}*/
	// 判断违禁词库中是否存在name这个违禁词语
	if GetManageBanWord().IsBanWord(name) {
		return
	}
	//调用外部链接，http地址接口，验证字符串是否合法(外部接口可能出现问题)
	//正确做法：有一个违禁词库
	player.ModPlayer.Name = name
	fmt.Println("当前名字:", player.ModPlayer.Name)
}
func (self *ModPlayer) SetSign(sign string, player *Player) {
	/*if !player.ModCard.IsHasCard(name) {
		return
	}*/
	if GetManageBanWord().IsBanWord(sign) {
		return
	}
	player.ModPlayer.Sign = sign
	fmt.Println("当前签名:", player.ModPlayer.Sign)
}

// AddExp 增加等级
func (self *ModPlayer) AddExp(exp int, player *Player) {
	//经验直接加上去，升级用后面的处理
	self.PlayerExp += exp
	for {
		//这里需要循环获取值
		config := csvs.GetNowLevelConfig(self.PlayerLevel)
		if config == nil {
			break
		}
		//这里是达到上限了，跳出去即可
		if config.PlayerExp == 0 {
			break
		}
		//是否任务完成 todo
		// 确保大于0，并且没有完成，退出,不允许升级
		if config.ChapterId > 0 && !player.ModUniqueTask.IsTaskFinish(config.ChapterId) {
			break
		}
		//计算任务升级
		//如果当前经验大于配置中的经验
		if self.PlayerExp >= config.PlayerExp {
			self.PlayerLevel += 1              //等级加1
			self.PlayerExp -= config.PlayerExp //扣除经验，重新开始计算
		} else {
			break
		}
	}
	fmt.Println("当前等级：", self.PlayerLevel, "---当前经验：", self.PlayerExp)
}
