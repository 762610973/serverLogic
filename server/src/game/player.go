package game

import (
	"fmt"
	"serverLogic/server/src/csvs"
)

//每一个玩家就是一个结构,目前是包含两个模块
//客户端直接和player模块交互，player收到消息。调用其他模块处理
//receive函数，跟客户端打交道的函数
//增加经验的接口一定是内部接口

type Player struct {
	ModPlayer     *ModPlayer     //基础模块
	ModIcon       *ModIcon       //头像模块
	ModCard       *ModCard       //名片模块
	ModUniqueTask *ModUniqueTask //唯一任务模块
	ModRole       *ModRole       //角色模块
	ModBag        *ModBag        //背包模块
	ModWeapon     *ModWeapon     //武器模块
	ModRelics     *ModRelics     //圣遗物模块
	ModCook       *ModCook       //烹饪模块
	ModHome       *ModHome       //家园模块
	ModPool       *ModPool       //池子
	ModMap        *ModMap        //地图模块
}

// NewTestPlayer 生成玩家
// NewTestPlayer :new一个测试玩家
func NewTestPlayer() *Player {
	//模块初始化
	player := new(Player)
	player.ModPlayer = new(ModPlayer)

	player.ModIcon = new(ModIcon)
	player.ModIcon.IconInfo = make(map[int]*Icon)

	player.ModCard = new(ModCard)
	player.ModCard.CardInfo = make(map[int]*Card)

	player.ModUniqueTask = new(ModUniqueTask)
	player.ModUniqueTask.MyTaskInfo = make(map[int]*TaskInfo)
	//player.ModUniqueTask.Locker = new(sync.RWMutex)

	player.ModRole = new(ModRole)
	player.ModRole.RoleInfo = make(map[int]*RoleInfo)

	player.ModBag = new(ModBag)
	player.ModBag.BagInfo = make(map[int]*ItemInfo)

	player.ModWeapon = new(ModWeapon)
	player.ModWeapon.WeaponInfo = make(map[int]*Weapon)

	player.ModRelics = new(ModRelics)
	player.ModRelics.RelicsInfo = make(map[int]*Relics)

	player.ModCook = new(ModCook)
	player.ModCook.CookInfo = make(map[int]*Cook)

	player.ModHome = new(ModHome)
	player.ModHome.HomeItemIdInfo = make(map[int]*HomeItemId)

	player.ModPool = new(ModPool)
	player.ModPool.UpPoolInfo = new(PoolInfo)

	player.ModMap = new(ModMap)
	player.ModMap.InitData()

	player.ModPlayer.PlayerLevel = 1 //初始等级是1级
	player.ModPlayer.WorldLevel = 1
	player.ModPlayer.WorldLevelNow = 1
	player.ModPlayer.Name = "旅行者"
	//以上是模块初始化，下面是数据初始化
	//player.ModPlayer.Icon = 0

	return player
}

// ReceiveSetIcon ... 对外接口
// 判断这个玩家是否含有这个icon，如果有就赋值，没有就返回
// ReceiveSetIcon 与客户端交互，意思是接收到要更改Icon的信息，那么就调用SetIcon，然后将iconId和这个玩家的指针传进去
func (p *Player) ReceiveSetIcon(iconId int) {
	p.ModPlayer.SetIcon(iconId, p)
}

// ReceiveSetCard 设置名片
func (p *Player) ReceiveSetCard(cardId int) {
	p.ModPlayer.SetCard(cardId, p)
}

// ReceiveSetName 修改名字
func (p *Player) ReceiveSetName(name string) {
	p.ModPlayer.SetName(name, p)
}

// ReceiveSetSign 设置签名
func (p *Player) ReceiveSetSign(sign string) {
	p.ModPlayer.SetSign(sign, p)
}

// ReduceWorldLevel 降低世界等级，是一个对外接口
func (p *Player) ReduceWorldLevel() {
	p.ModPlayer.ReduceWorldLevel(p)
}

// ReturnWorldLevel 返回世界等级
func (p *Player) ReturnWorldLevel() {
	p.ModPlayer.ReturnWorldLevel(p)
}

// SetBirth 设置生日
func (p *Player) SetBirth(birth int) {
	p.ModPlayer.SetBirth(birth, p)
}

// SetShowCard 设置展示的名片
func (p *Player) SetShowCard(showCard []int) {
	p.ModPlayer.SetShowCard(showCard, p)
}

func (p *Player) SetShowTeam(showRole []int) {
	p.ModPlayer.SetShowTeam(showRole, p)
}

func (p *Player) SetHideShowTeam(isHead int) {
	p.ModPlayer.SetHideShowTeam(isHead, p)
}

// Run /*func (p *Player) Run() {
func (p *Player) Run() {
	fmt.Println("从0开始写原神服务器------测试工具v0.1")
	fmt.Println("作者:B站------golang大海葵")
	fmt.Println("模拟用户创建成功OK------开始测试")
	fmt.Println("↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓")
	for {
		fmt.Println(p.ModPlayer.Name, ",欢迎来到提瓦特大陆,请选择功能：1基础信息2背包3(优菈UP池)模拟抽卡1000W次4地图(未开放)")
		var modChoose int
		fmt.Scan(&modChoose)
		switch modChoose {
		case 1:
			p.HandleBase()
		case 2:
			p.HandleBag()
		case 3:
			p.HandlePool()
		case 4:
			p.HandleMap()
		}
	}
}

// HandleBase 基础信息
func (p *Player) HandleBase() {
	for {
		fmt.Println("当前处于基础信息界面,请选择操作：0返回1查询信息2设置名字3设置签名4头像5名片6设置生日")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			p.HandleBaseGetInfo()
		case 2:
			p.HandleBagSetName()
		case 3:
			p.HandleBagSetSign()
		case 4:
			p.HandleBagSetIcon()
		case 5:
			p.HandleBagSetCard()
		case 6:
			p.HandleBagSetBirth()
		}
	}
}

func (p *Player) HandleBaseGetInfo() {
	fmt.Println("名字:", p.ModPlayer.Name)
	fmt.Println("等级:", p.ModPlayer.PlayerLevel)
	fmt.Println("大世界等级:", p.ModPlayer.WorldLevelNow)
	if p.ModPlayer.Sign == "" {
		fmt.Println("签名:", "未设置")
	} else {
		fmt.Println("签名:", p.ModPlayer.Sign)
	}

	if p.ModPlayer.Icon == 0 {
		fmt.Println("头像:", "未设置")
	} else {
		fmt.Println("头像:", csvs.GetItemConfig(p.ModPlayer.Icon), p.ModPlayer.Icon)
	}

	if p.ModPlayer.Card == 0 {
		fmt.Println("名片:", "未设置")
	} else {
		fmt.Println("名片:", csvs.GetItemConfig(p.ModPlayer.Card), p.ModPlayer.Card)
	}

	if p.ModPlayer.Birth == 0 {
		fmt.Println("生日:", "未设置")
	} else {
		fmt.Println("生日:", p.ModPlayer.Birth/100, "月", p.ModPlayer.Birth%100, "日")
	}
}

func (p *Player) HandleBagSetName() {
	fmt.Println("请输入名字:")
	var name string
	fmt.Scan(&name)
	p.ReceiveSetName(name)
}

func (p *Player) HandleBagSetSign() {
	fmt.Println("请输入签名:")
	var sign string
	fmt.Scan(&sign)
	p.ReceiveSetSign(sign)
}

func (p *Player) HandleBagSetIcon() {
	for {
		fmt.Println("当前处于基础信息--头像界面,请选择操作：0返回1查询头像背包2设置头像")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			p.HandleBagSetIconGetInfo()
		case 2:
			p.HandleBagSetIconSet()
		}
	}
}

func (p *Player) HandleBagSetIconGetInfo() {
	fmt.Println("当前拥有头像如下:")
	for _, v := range p.ModIcon.IconInfo {
		config := csvs.GetItemConfig(v.IconId)
		if config != nil {
			fmt.Println(config.ItemName, ":", config.ItemId)
		}
	}
}

func (p *Player) HandleBagSetIconSet() {
	fmt.Println("请输入头像id:")
	var icon int
	fmt.Scan(&icon)
	p.ReceiveSetIcon(icon)
}

func (p *Player) HandleBagSetCard() {
	for {
		fmt.Println("当前处于基础信息--名片界面,请选择操作：0返回1查询名片背包2设置名片")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			p.HandleBagSetCardGetInfo()
		case 2:
			p.HandleBagSetCardSet()
		}
	}
}

func (p *Player) HandleBagSetCardGetInfo() {
	fmt.Println("当前拥有名片如下:")
	for _, v := range p.ModCard.CardInfo {
		config := csvs.GetItemConfig(v.CardId)
		if config != nil {
			fmt.Println(config.ItemName, ":", config.ItemId)
		}
	}
}

func (p *Player) HandleBagSetCardSet() {
	fmt.Println("请输入名片id:")
	var card int
	fmt.Scan(&card)
	p.ReceiveSetCard(card)
}

func (p *Player) HandleBagSetBirth() {
	if p.ModPlayer.Birth > 0 {
		fmt.Println("已设置过生日!")
		return
	}
	fmt.Println("生日只能设置一次，请慎重填写,输入月:")
	var month, day int
	fmt.Scan(&month)
	fmt.Println("请输入日:")
	fmt.Scan(&day)
	p.ModPlayer.SetBirth(month*100+day, p)
}

// HandleBag 背包
func (p *Player) HandleBag() {
	for {
		fmt.Println("当前处于基础信息界面,请选择操作：0返回1增加物品2扣除物品3使用物品")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			p.HandleBagAddItem()
		case 2:
			p.HandleBagRemoveItem()
		case 3:
			p.HandleBagUseItem()
		}
	}
}

// HandlePool 抽卡
func (p *Player) HandlePool() {
	p.ModPool.DoUpPool()
}

func (p *Player) HandleBagAddItem() {
	itemId := 0
	itemNum := 0
	fmt.Println("物品ID")
	fmt.Scan(&itemId)
	fmt.Println("物品数量")
	fmt.Scan(&itemNum)
	p.ModBag.AddItem(itemId, int64(itemNum), p)
}

func (p *Player) HandleBagRemoveItem() {
	itemId := 0
	itemNum := 0
	fmt.Println("物品ID")
	fmt.Scan(&itemId)
	fmt.Println("物品数量")
	fmt.Scan(&itemNum)
	p.ModBag.RemoveItemToBag(itemId, int64(itemNum), p)
}

func (p *Player) HandleBagUseItem() {
	itemId := 0
	itemNum := 0
	fmt.Println("物品ID")
	fmt.Scan(&itemId)
	fmt.Println("物品数量")
	fmt.Scan(&itemNum)
	p.ModBag.UseItem(itemId, int64(itemNum), p)
}

// HandleMap 地图
func (p *Player) HandleMap() {
	fmt.Println("向着星辰与深渊,欢迎来到冒险家协会！")
	fmt.Println("当前位置:", "蒙德城")
	fmt.Println("地图模块还没写到......")
}
