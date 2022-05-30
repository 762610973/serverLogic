package game

import (
	"time"
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

	player.ModPlayer.PlayerLevel = 1 //初始等级是1级
	player.ModPlayer.WorldLevel = 6
	player.ModPlayer.WorldLevelNow = 6

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

func (p *Player) Run() {
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			if time.Now().Unix()%5 == 0 {
				p.ModBag.AddItem(2000017, 7, p)
			}
			//fmt.Println(time.Now().Unix())
		}
	}
}
