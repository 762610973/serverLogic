package game

import (
	"fmt"
	"serverLogic/server/src/csvs"
	"time"
)

//每一个模块都对应数据库中的一张表
//玩家模块，包含以下的字段

type ShowRole struct {
	RoleId    int
	RoleLevel int
}

type ModPlayer struct {
	UserId         int         //用户id
	Icon           int         //头像
	Card           int         //名片
	Name           string      //名称
	Sign           string      //签名
	PlayerLevel    int         //等级
	PlayerExp      int         //阅历（经验）
	WorldLevel     int         //世界等级
	WorldLevelNow  int         //大世界当前等级
	WorldLevelCool int64       //世界等级冷却时间
	Birth          int         //生日，只能改一次
	ShowTeam       []*ShowRole //展示阵容
	HideShowTeam   int         //展示隐藏的开关，0/1
	ShowCard       []int       //展示名片
	//看不见的字段
	Prohibit int //是否封禁，很少用bool，int方便扩展
	IsGm     int //是否是游戏管理员
}

func (m *ModPlayer) SetIcon(iconId int, player *Player) {
	//自己的icon模块没有这个头像
	if !player.ModIcon.IsHasIcon(iconId) {
		//没有这个icon，通知客户端，操作非法
		fmt.Println("没有头像:", iconId)
		return
	}
	player.ModPlayer.Icon = iconId
	fmt.Println("当前头像:", player.ModPlayer.Icon)
}

// SetCard 设置名片
func (m *ModPlayer) SetCard(cardId int, player *Player) {
	if !player.ModCard.IsHasCard(cardId) {
		return
	}
	player.ModPlayer.Card = cardId
	fmt.Println("当前名片:", player.ModPlayer.Card)
}

// SetName 设置名字
func (m *ModPlayer) SetName(name string, player *Player) {
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

// SetSign 设置签名
func (m *ModPlayer) SetSign(sign string, player *Player) {
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
func (m *ModPlayer) AddExp(exp int, player *Player) {
	//经验直接加上去，升级用后面的处理
	m.PlayerExp += exp
	for {
		//这里需要循环获取值
		config := csvs.GetNowLevelConfig(m.PlayerLevel)
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
		if m.PlayerExp >= config.PlayerExp {
			m.PlayerLevel += 1              //等级加1
			m.PlayerExp -= config.PlayerExp //扣除经验，重新开始计算
		} else {
			break
		}
	}
	fmt.Println("当前等级：", m.PlayerLevel, "---当前经验：", m.PlayerExp)
}

// ReduceWorldLevel 降低世界等级
func (m *ModPlayer) ReduceWorldLevel(player *Player) {
	// 如果等级小于5级，就不能进行降级
	if m.WorldLevel < csvs.ReduceWorldLevelStart {
		// 当前世界等级过低
		fmt.Println("操作失败，-----当时世界等级", m.WorldLevel)
		return
	}
	// 如果自己的等级-当前世界等级大于等于1级，返回（最多降低一级）
	if m.WorldLevel-m.WorldLevelNow >= csvs.ReduceWorldLevelMax {
		fmt.Println("操作失败：-----当前世界等级", m.WorldLevel, "----真实世界等级，", m.WorldLevelNow)
		return
	}
	if time.Now().Unix() < m.WorldLevelCool {
		//没到达冷却时间
		fmt.Println("操作失败：-----冷却中")
		return
	}
	m.WorldLevel -= 1
	m.WorldLevelCool = time.Now().Unix() + csvs.ReduceWorldCoolTime
	fmt.Println("操作成功：------当前世界等级:", m.WorldLevel, "-----真实世界等级", m.WorldLevelNow)
	return
}

func (m *ModPlayer) ReturnWorldLevel(player *Player) {
	if m.WorldLevelNow == m.WorldLevel {
		fmt.Println("操作失败：------当前世界等级", m.WorldLevel, "-----真实世界等级", m.WorldLevelNow)
		return
	}
	if time.Now().Unix() < m.WorldLevelCool {
		//没到达冷却时间
		fmt.Println("操作失败：-----冷却中")
		return
	}
	m.WorldLevel += 1
	m.WorldLevelCool = time.Now().Unix() + csvs.ReduceWorldCoolTime
	fmt.Println("操作成功：------当前世界等级:", m.WorldLevel, "-----真实世界等级", m.WorldLevelNow)
	return
}

// SetBirth 设置生日
func (m *ModPlayer) SetBirth(birth int, player *Player) {
	if m.Birth > 0 {
		fmt.Println("已设置过生日，不允许再次设置")
		return
	}
	month := birth / 100
	day := birth % 100
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
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
	default:
		fmt.Println("没有", month, "月")
		return
	}

	m.Birth = birth
	fmt.Println("设置成功，生日为：", month, "月", day, "日")
	if m.IsBirthDay() {
		fmt.Println("今天是你的生日，生日快乐")
	} else {
		fmt.Println("期待你生日的到来")
	}
}

// IsBirthDay 判断今天是否是生日
func (m *ModPlayer) IsBirthDay() bool {
	month := time.Now().Month()
	day := time.Now().Day()
	if int(month) == m.Birth/100 && day == m.Birth%100 {
		return true
	}
	return false
}

// SetShowCard 设置展示名片
func (m *ModPlayer) SetShowCard(showCard []int, player *Player) {
	if len(showCard) > csvs.ShowSize {
		return
	}
	// 验证数量(客户端最多展示9个)以及是否存在重复的
	cardExit := make(map[int]int, 9)
	newList := make([]int, 0, 9)
	for _, cardId := range showCard {
		_, ok := cardExit[cardId]
		if ok {
			//存在
			continue
		}
		if !player.ModCard.IsHasCard(cardId) {
			continue
		}
		// 不存在的话就加到这个里面，设置这个的目的主要是保证顺序，map遍历是无序的
		newList = append(newList, cardId)
		cardExit[cardId] = 1
	}
	m.ShowCard = newList
	fmt.Println(m.ShowCard)
}

// SetShowTeam 展示阵容
func (m *ModPlayer) SetShowTeam(showRole []int, player *Player) {
	if len(showRole) > csvs.ShowSize {
		fmt.Println("消息结构错误")
		return
	}
	roleExist := make(map[int]int)
	newList := make([]*ShowRole, 0, 9)
	for _, roleId := range showRole {
		_, ok := roleExist[roleId]
		if ok {
			continue
		}
		if !player.ModRole.IsHasRole(roleId) {
			continue
		}
		showRole := new(ShowRole)
		showRole.RoleId = roleId
		showRole.RoleLevel = player.ModRole.GetRoleLevel(roleId)
		newList = append(newList, showRole)
		roleExist[roleId] = 1
	}
	m.ShowTeam = newList
	fmt.Println(m.ShowTeam)
}

// SetHideShowTeam 展示隐藏阵容
func (m *ModPlayer) SetHideShowTeam(isHide int, player *Player) {
	if isHide != csvs.LoginFalse && isHide != csvs.LoginTrue {
		return
	}
	m.HideShowTeam = isHide
}

func (m *ModPlayer) SetProhibit(prohibit int) {
	m.Prohibit = prohibit
}

func (m *ModPlayer) SetIsGM(isGm int) {
	m.IsGm = isGm
}

func (m *ModPlayer) IsCanEnter() bool {
	return int64(m.Prohibit) < time.Now().Unix()
}
