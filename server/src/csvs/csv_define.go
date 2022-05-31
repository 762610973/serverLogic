package csvs

const (
	ReduceWorldLevelStart = 5  //降低世界等级的要求
	ReduceWorldLevelMax   = 1  //最低能降低多少级
	ReduceWorldCoolTime   = 10 //冷却时间，目前是10秒，应该是24*3600
	ShowSize              = 9
	AddRoleTimeNormalMin  = 2
	AddRoleTimeNormalMax  = 7

	WeaponMaxCount              = 2000
	RelicsMaxCount              = 1500
	FiveStarTimesLimit          = 73
	FiveStarTimesLimitEachValue = 600
	FourStarTimesLimit          = 8
	FourStarTimesLimitEachValue = 5100
	AllEntryRate                = 2000
	WeaponMaxRefine             = 5 //武器的最大精炼等级
)

const (
	DROP_ITEM_TYPE_ITEM   = 1
	DROP_ITEM_TYPE_GROUP  = 2
	DROP_ITEM_TYPE_WEIGHT = 3
)

// 蒲公英只有0/10两种，矿物有三种
const (
	EventStart  = 0
	EventFinish = 9
	EventEnd    = 10

	EVENT_TYPE_NORMAL = 1
	EVENT_TYPE_REWARD = 2
)

const (
	MapRefreshDay    = 1
	MapRefreshWeek   = 2
	MapRefreshSelf   = 3
	MAP_REFRESH_CANT = 4

	MapRefreshDayTime  = 20
	MapRefreshWeekTime = 40
	MapRefreshSelfTime = 60

	REFRESH_SYSTEM = 1
	REFRESH_PLAYER = 2
)

const (
	LoginFalse = 0
	LoginTrue  = 1
	PercentAll = 10000
)
