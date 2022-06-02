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

func (p *Player) SetEventState(state int) {
	//p.ModMap.SetEventState(state, p)
}

func (p *Player) Run() {
	fmt.Println("从0开始写原神服务器------测试工具")
	fmt.Println("模拟用户创建成功OK------开始测试")
	fmt.Println("↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓")
	for {
		fmt.Println(p.ModPlayer.Name, ",欢迎来到提瓦特大陆,请选择功能：1基础信息2背包3角色(八重神子UP池)4地图5圣遗物6角色7武器")
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
		case 5:
			p.HandleRelics()
		case 6:
			p.HandleRole()
		case 7:
			p.HandleWeapon()
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
		fmt.Println("当前处于基础信息界面,请选择操作：0返回1增加物品2扣除物品3使用物品4升级七天神像(风)")
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
		case 4:
			p.HandleBagWindStatue()
		}
	}
}

// HandlePool 抽卡
func (p *Player) HandlePool() {
	for {
		fmt.Println("当前处于模拟抽卡界面,请选择操作：0返回1角色信息2十连抽(入包)3单抽(可选次数,入包)" +
			"4五星爆点测试5十连多黄测试6视频原版函数(30秒)7单抽(仓检版,独宠一人)8单抽(仓检版,雨露均沾)")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			p.ModRole.HandleSendRoleInfo(p)
		case 2:
			p.ModPool.HandleUpPoolTen(p)
		case 3:
			fmt.Println("请输入抽卡次数,最大值1亿(最大耗时约30秒):")
			var times int
			fmt.Scan(&times)
			p.ModPool.HandleUpPoolSingle(times, p)
		case 4:
			fmt.Println("请输入抽卡次数,最大值1亿(最大耗时约30秒):")
			var times int
			fmt.Scan(&times)
			p.ModPool.HandleUpPoolTimesTest(times)
		case 5:
			fmt.Println("请输入抽卡次数,最大值1亿(最大耗时约30秒):")
			var times int
			fmt.Scan(&times)
			p.ModPool.HandleUpPoolFiveTest(times)
		case 6:
			p.ModPool.DoUpPool()
		case 7:
			fmt.Println("请输入抽卡次数,最大值1亿(最大耗时约30秒):")
			var times int
			fmt.Scan(&times)
			p.ModPool.HandleUpPoolSingleCheck1(times, p)
		case 8:
			fmt.Println("请输入抽卡次数,最大值1亿(最大耗时约30秒):")
			var times int
			fmt.Scan(&times)
			p.ModPool.HandleUpPoolSingleCheck2(times, p)
		}
	}
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

func (p *Player) HandleBagWindStatue() {
	fmt.Println("开始升级七天神像")
	p.ModMap.UpStatue(1, p)
	p.ModRole.CalHpPool()
}

// HandleMap 地图
func (p *Player) HandleMap() {
	fmt.Println("向着星辰与深渊,欢迎来到冒险家协会！")
	for {
		fmt.Println("请选择互动地图1蒙德2璃月1001深入风龙废墟2001无妄引咎密宫")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		default:
			p.HandleMapIn(action)
		}
	}
}

func (p *Player) HandleMapIn(mapId int) {
	config := csvs.ConfigMapMap[mapId]
	if config == nil {
		fmt.Println("无法识别的地图")
		return
	}
	p.ModMap.RefreshByPlayer(mapId)
	for {
		p.ModMap.GetEventList(config)
		fmt.Println("请选择触发事件Id(0返回)")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		default:
			eventConfig := csvs.ConfigMapEventMap[action]
			if eventConfig == nil {
				fmt.Println("无法识别的事件")
				break
			}
			p.ModMap.SetEventState(mapId, eventConfig.EventId, csvs.EventEnd, p)
		}
	}
}

func (p *Player) HandleRelics() {
	for {
		fmt.Println("当前处于圣遗物界面，选择功能0返回1强化测试2满级圣遗物3极品头测试")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			p.ModRelics.RelicsUp(p)
		case 2:
			p.ModRelics.RelicsTop(p)
		case 3:
			p.ModRelics.RelicsTestBest(p)
		default:
			fmt.Println("无法识别在操作")
		}
	}
}

func (p *Player) HandleRole() {
	for {
		fmt.Println("当前处于角色界面，选择功能0返回1查询2穿戴圣遗物3卸下圣遗物4穿戴武器5卸下武器")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			p.ModRole.HandleSendRoleInfo(p)
		case 2:
			p.HandleWearRelics()
		case 3:
			p.HandleTakeOffRelics()
		case 4:
			p.HandleWearWeapon()
		case 5:
			p.HandleTakeOffWeapon()
		default:
			fmt.Println("无法识别在操作")
		}
	}
}

func (p *Player) HandleWearRelics() {
	for {
		fmt.Println("输入操作的目标英雄Id:,0返回")
		var roleId int
		fmt.Scan(&roleId)

		if roleId == 0 {
			return
		}

		RoleInfo := p.ModRole.RoleInfo[roleId]
		if RoleInfo == nil {
			fmt.Println("英雄不存在")
			continue
		}

		RoleInfo.ShowInfo(p)
		fmt.Println("输入需要穿戴的圣遗物key:,0返回")
		var relicsKey int
		fmt.Scan(&relicsKey)
		if relicsKey == 0 {
			return
		}
		relics := p.ModRelics.RelicsInfo[relicsKey]
		if relics == nil {
			fmt.Println("圣遗物不存在")
			continue
		}
		p.ModRole.WearRelics(RoleInfo, relics, p)
	}
}

func (p *Player) HandleTakeOffRelics() {
	for {
		fmt.Println("输入操作的目标英雄Id:,0返回")
		var roleId int
		fmt.Scan(&roleId)

		if roleId == 0 {
			return
		}

		RoleInfo := p.ModRole.RoleInfo[roleId]
		if RoleInfo == nil {
			fmt.Println("英雄不存在")
			continue
		}

		RoleInfo.ShowInfo(p)
		fmt.Println("输入需要卸下的圣遗物key:,0返回")
		var relicsKey int
		fmt.Scan(&relicsKey)
		if relicsKey == 0 {
			return
		}
		relics := p.ModRelics.RelicsInfo[relicsKey]
		if relics == nil {
			fmt.Println("圣遗物不存在")
			continue
		}
		p.ModRole.TakeOffRelics(RoleInfo, relics, p)
	}
}

func (p *Player) HandleWeapon() {
	for {
		fmt.Println("当前处于武器界面，选择功能0返回1强化测试2突破测试3精炼测试")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			p.HandleWeaponUp()
		case 2:
			p.HandleWeaponStarUp()
		case 3:
			p.HandleWeaponRefineUp()
		default:
			fmt.Println("无法识别在操作")
		}
	}
}

func (p *Player) HandleWeaponUp() {
	for {
		fmt.Println("输入操作的目标武器keyId:,0返回")
		for _, v := range p.ModWeapon.WeaponInfo {
			fmt.Println(fmt.Sprintf("武器keyId:%d,等级:%d,突破等级:%d,精炼:%d",
				v.KeyId, v.Level, v.StarLevel, v.RefineLevel))
		}
		var weaponKeyId int
		fmt.Scan(&weaponKeyId)
		if weaponKeyId == 0 {
			return
		}
		p.ModWeapon.WeaponUp(weaponKeyId, p)
	}
}

func (p *Player) HandleWeaponStarUp() {
	for {
		fmt.Println("输入操作的目标武器keyId:,0返回")
		for _, v := range p.ModWeapon.WeaponInfo {
			fmt.Println(fmt.Sprintf("武器keyId:%d,等级:%d,突破等级:%d,精炼:%d",
				v.KeyId, v.Level, v.StarLevel, v.RefineLevel))
		}
		var weaponKeyId int
		fmt.Scan(&weaponKeyId)
		if weaponKeyId == 0 {
			return
		}
		p.ModWeapon.WeaponUpStar(weaponKeyId, p)
	}
}

func (p *Player) HandleWeaponRefineUp() {
	for {
		fmt.Println("输入操作的目标武器keyId:,0返回")
		for _, v := range p.ModWeapon.WeaponInfo {
			fmt.Println(fmt.Sprintf("武器keyId:%d,等级:%d,突破等级:%d,精炼:%d",
				v.KeyId, v.Level, v.StarLevel, v.RefineLevel))
		}
		var weaponKeyId int
		fmt.Scan(&weaponKeyId)
		if weaponKeyId == 0 {
			return
		}
		for {
			fmt.Println("输入作为材料的武器keyId:,0返回")
			var weaponTargetKeyId int
			fmt.Scan(&weaponTargetKeyId)
			if weaponTargetKeyId == 0 {
				return
			}
			p.ModWeapon.WeaponUpRefine(weaponKeyId, weaponTargetKeyId, p)
		}
	}
}

func (p *Player) HandleWearWeapon() {
	for {
		fmt.Println("输入操作的目标英雄Id:,0返回")
		var roleId int
		fmt.Scan(&roleId)

		if roleId == 0 {
			return
		}

		RoleInfo := p.ModRole.RoleInfo[roleId]
		if RoleInfo == nil {
			fmt.Println("英雄不存在")
			continue
		}

		RoleInfo.ShowInfo(p)
		fmt.Println("输入需要穿戴的武器key:,0返回")
		var weaponKey int
		fmt.Scan(&weaponKey)
		if weaponKey == 0 {
			return
		}
		weaponInfo := p.ModWeapon.WeaponInfo[weaponKey]
		if weaponInfo == nil {
			fmt.Println("武器不存在")
			continue
		}
		p.ModRole.WearWeapon(RoleInfo, weaponInfo, p)
		RoleInfo.ShowInfo(p)
	}
}

func (p *Player) HandleTakeOffWeapon() {
	for {
		fmt.Println("输入操作的目标英雄Id:,0返回")
		var roleId int
		fmt.Scan(&roleId)

		if roleId == 0 {
			return
		}

		RoleInfo := p.ModRole.RoleInfo[roleId]
		if RoleInfo == nil {
			fmt.Println("英雄不存在")
			continue
		}

		RoleInfo.ShowInfo(p)
		fmt.Println("输入需要卸下的武器key:,0返回")
		var weaponKey int
		fmt.Scan(&weaponKey)
		if weaponKey == 0 {
			return
		}
		weapon := p.ModWeapon.WeaponInfo[weaponKey]
		if weapon == nil {
			fmt.Println("武器不存在")
			continue
		}
		p.ModRole.TakeOffWeapon(RoleInfo, weapon, p)
		RoleInfo.ShowInfo(p)
	}
}
