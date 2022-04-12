package game

import (
	"fmt"
	"serverLogic/server/src/csvs"
	"time"
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

// ReduceWorldLevel 降低世界等级
func (self *ModPlayer) ReduceWorldLevel(player *Player) {
	// 如果等级小于5级，就不能进行降级
	if self.WorldLevel < csvs.ReduceWorldLevelStart {
		fmt.Println("操作失败，-----当时世界等级", self.WorldLevel)
		return
	}
	// 如果自己的等级-当前世界等级大于等于1级，返回（最多降低一级）
	if self.WorldLevel-self.WorldLevelNow >= csvs.ReduceWorldLevelMax {
		fmt.Println("操作失败：-----当前世界等级", self.WorldLevel, "----真实世界等级，", self.WorldLevelNow)
		return
	}
	if time.Now().Unix() < self.WorldLevelCool {
		//没到达冷却时间
		fmt.Println("操作失败：-----冷却中")
		return
	}
	self.WorldLevel -= 1
	self.WorldLevelCool = time.Now().Unix() + csvs.ReduceWorldCoolTime
	fmt.Println("操作成功：------当前世界等级:", self.WorldLevel, "-----真实世界等级", self.WorldLevelNow)
	return
}
func (self *ModPlayer) ReturnWorldLevel(player *Player) {
	if self.WorldLevelNow == self.WorldLevel {
		fmt.Println("操作失败：------当前世界等级", self.WorldLevel, "-----真实世界等级", self.WorldLevelNow)
		return
	}
	if time.Now().Unix() < self.WorldLevelCool {
		//没到达冷却时间
		fmt.Println("操作失败：-----冷却中")
		return
	}
	self.WorldLevel += 1
	self.WorldLevelCool = time.Now().Unix() + csvs.ReduceWorldCoolTime
	fmt.Println("操作成功：------当前世界等级:", self.WorldLevel, "-----真实世界等级", self.WorldLevelNow)
	return
}

func (self *ModPlayer) SetBirth(birth int, player *Player) {
	month := birth / 100
	day := birth % 100
	switch month {
	case 1, 3, 5, 7, 8, 0, 12:
		if day <= 0 || day > 31 {
			fmt.Println(month, "月没有", day, "日！")
			return
		}
	case 4, 6, 9, 11:
		if day <= 0 || day > 30 {
			fmt.Println(month, "月没有", day, "日！")
			return
		}
	case 2:
		if day <= 0 || day > 29 {
			fmt.Println(month, "月没有", day, "日！")
			return
		}
	}
	//4
}
