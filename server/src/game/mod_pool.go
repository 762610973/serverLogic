package game

import (
	"fmt"
	"serverLogic/server/src/csvs"
)

type PoolInfo struct {
	PoolId        int
	FiveStarTimes int
	FourStarTimes int
	IsMustUp      int
}

type ModPool struct {
	UpPoolInfo *PoolInfo
}

func (m *ModPool) AddTimes() {
	m.UpPoolInfo.FiveStarTimes++
	m.UpPoolInfo.FourStarTimes++
}

func (m *ModPool) DoUpPool() {
	result := make(map[int]int)
	fourNum := 0
	fiveNum := 0
	resultEach := make(map[int]int)
	resultEachTest := make(map[int]int)
	fiveTest := 0
	for i := 0; i < 100000000; i++ {
		m.AddTimes()
		if i%10 == 0 {
			fiveTest = 0
		}
		dropGroup := csvs.ConfigDropGroupMap[1000]
		if dropGroup == nil {
			return
		}

		if m.UpPoolInfo.FiveStarTimes > csvs.FiveStarTimesLimit || m.UpPoolInfo.FourStarTimes > csvs.FourStarTimesLimit {
			newDropGroup := new(csvs.DropGroup)
			newDropGroup.DropId = dropGroup.DropId
			newDropGroup.WeightAll = dropGroup.WeightAll
			addFiveWeight := (m.UpPoolInfo.FiveStarTimes - csvs.FiveStarTimesLimit) * csvs.FiveStarTimesLimitEachValue
			if addFiveWeight < 0 {
				addFiveWeight = 0
			}
			addFourWeight := (m.UpPoolInfo.FourStarTimes - csvs.FourStarTimesLimit) * csvs.FourStarTimesLimitEachValue
			if addFourWeight < 0 {
				addFourWeight = 0
			}

			for _, config := range dropGroup.DropConfigs {
				newConfig := new(csvs.ConfigDrop)
				newConfig.Result = config.Result
				newConfig.DropId = config.DropId
				newConfig.IsEnd = config.IsEnd
				if config.Result == 10001 {
					newConfig.Weight = config.Weight + addFiveWeight
				} else if config.Result == 10002 {
					newConfig.Weight = config.Weight + addFourWeight
				} else if config.Result == 10003 {
					newConfig.Weight = config.Weight - addFiveWeight - addFourWeight
				}
				newDropGroup.DropConfigs = append(newDropGroup.DropConfigs, newConfig)
			}
			dropGroup = newDropGroup
		}
		roleIdConfig := csvs.GetRandDropNew(dropGroup)
		if roleIdConfig != nil {
			roleConfig := csvs.GetRoleConfig(roleIdConfig.Result)
			if roleConfig != nil {
				if roleConfig.Star == 5 {
					fiveTest++
					resultEach[m.UpPoolInfo.FiveStarTimes]++
					m.UpPoolInfo.FiveStarTimes = 0
					fiveNum++
					if m.UpPoolInfo.IsMustUp == csvs.LoginTrue {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("数据异常")
								return
							}
						}
					}
					if roleIdConfig.DropId == 100012 {
						m.UpPoolInfo.IsMustUp = csvs.LoginFalse
					} else {
						m.UpPoolInfo.IsMustUp = csvs.LoginTrue
					}
				} else if roleConfig.Star == 4 {
					m.UpPoolInfo.FourStarTimes = 0
					fourNum++
				}
			}
			result[roleIdConfig.Result]++
		}
		if i%10 == 9 {
			resultEachTest[fiveTest]++
		}
	}

	for k, v := range result {
		fmt.Println(fmt.Sprintf("抽中%s次数：%d", csvs.GetItemName(k), v))
	}
	fmt.Println(fmt.Sprintf("抽中4星：%d", fourNum))
	fmt.Println(fmt.Sprintf("抽中5星：%d", fiveNum))

	for k, v := range resultEach {
		fmt.Println(fmt.Sprintf("第%d抽抽出5星的次数：%d", k, v))
	}

	for k, v := range resultEachTest {
		fmt.Println(fmt.Sprintf("10连%d黄次数：%d", k, v))
	}
}

func (m *ModPool) HandleUpPoolTen(player *Player) {
	for i := 0; i < 10; i++ {
		m.AddTimes()
		dropGroup := csvs.ConfigDropGroupMap[1000]
		if dropGroup == nil {
			return
		}

		if m.UpPoolInfo.FiveStarTimes > csvs.FiveStarTimesLimit || m.UpPoolInfo.FourStarTimes > csvs.FourStarTimesLimit {
			newDropGroup := new(csvs.DropGroup)
			newDropGroup.DropId = dropGroup.DropId
			newDropGroup.WeightAll = dropGroup.WeightAll
			addFiveWeight := (m.UpPoolInfo.FiveStarTimes - csvs.FiveStarTimesLimit) * csvs.FiveStarTimesLimitEachValue
			if addFiveWeight < 0 {
				addFiveWeight = 0
			}
			addFourWeight := (m.UpPoolInfo.FourStarTimes - csvs.FourStarTimesLimit) * csvs.FourStarTimesLimitEachValue
			if addFourWeight < 0 {
				addFourWeight = 0
			}

			for _, config := range dropGroup.DropConfigs {
				newConfig := new(csvs.ConfigDrop)
				newConfig.Result = config.Result
				newConfig.DropId = config.DropId
				newConfig.IsEnd = config.IsEnd
				if config.Result == 10001 {
					newConfig.Weight = config.Weight + addFiveWeight
				} else if config.Result == 10002 {
					newConfig.Weight = config.Weight + addFourWeight
				} else if config.Result == 10003 {
					newConfig.Weight = config.Weight - addFiveWeight - addFourWeight
				}
				newDropGroup.DropConfigs = append(newDropGroup.DropConfigs, newConfig)
			}
			dropGroup = newDropGroup
		}
		roleIdConfig := csvs.GetRandDropNew(dropGroup)
		if roleIdConfig != nil {
			roleConfig := csvs.GetRoleConfig(roleIdConfig.Result)
			if roleConfig != nil {
				if roleConfig.Star == 5 {
					m.UpPoolInfo.FiveStarTimes = 0
					if m.UpPoolInfo.IsMustUp == csvs.LoginTrue {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("数据异常")
								return
							}
						}
					}
					if roleIdConfig.DropId == 100012 {
						m.UpPoolInfo.IsMustUp = csvs.LoginFalse
					} else {
						m.UpPoolInfo.IsMustUp = csvs.LoginTrue
					}
				} else if roleConfig.Star == 4 {
					m.UpPoolInfo.FourStarTimes = 0
				}
			}
			//fmt.Println(fmt.Sprintf("第%d抽抽中:%s", i+1, csvs.GetItemName(roleIdConfig.Result)))
			player.ModBag.AddItem(roleIdConfig.Result, 1, player)
		}
	}
	if m.UpPoolInfo.IsMustUp == csvs.LoginFalse {
		fmt.Println(fmt.Sprintf("当前处于小保底区间！"))
	} else {
		fmt.Println(fmt.Sprintf("当前处于大保底区间！"))
	}
	fmt.Println(fmt.Sprintf("当前累计未出5星次数：%d", m.UpPoolInfo.FiveStarTimes))
	fmt.Println(fmt.Sprintf("当前累计未出4星次数：%d", m.UpPoolInfo.FourStarTimes))

}

func (m *ModPool) HandleUpPoolSingle(times int, player *Player) {
	if times <= 0 || times > 100000000 {
		fmt.Println("请输入正确的数值(1~100000000)")
		return
	} else {
		fmt.Println(fmt.Sprintf("累计抽取%d次,结果如下:", times))
	}
	result := make(map[int]int)
	fourNum := 0
	fiveNum := 0
	for i := 0; i < times; i++ {
		m.AddTimes()
		dropGroup := csvs.ConfigDropGroupMap[1000]
		if dropGroup == nil {
			return
		}

		if m.UpPoolInfo.FiveStarTimes > csvs.FiveStarTimesLimit || m.UpPoolInfo.FourStarTimes > csvs.FourStarTimesLimit {
			newDropGroup := new(csvs.DropGroup)
			newDropGroup.DropId = dropGroup.DropId
			newDropGroup.WeightAll = dropGroup.WeightAll
			addFiveWeight := (m.UpPoolInfo.FiveStarTimes - csvs.FiveStarTimesLimit) * csvs.FiveStarTimesLimitEachValue
			if addFiveWeight < 0 {
				addFiveWeight = 0
			}
			addFourWeight := (m.UpPoolInfo.FourStarTimes - csvs.FourStarTimesLimit) * csvs.FourStarTimesLimitEachValue
			if addFourWeight < 0 {
				addFourWeight = 0
			}

			for _, config := range dropGroup.DropConfigs {
				newConfig := new(csvs.ConfigDrop)
				newConfig.Result = config.Result
				newConfig.DropId = config.DropId
				newConfig.IsEnd = config.IsEnd
				if config.Result == 10001 {
					newConfig.Weight = config.Weight + addFiveWeight
				} else if config.Result == 10002 {
					newConfig.Weight = config.Weight + addFourWeight
				} else if config.Result == 10003 {
					newConfig.Weight = config.Weight - addFiveWeight - addFourWeight
				}
				newDropGroup.DropConfigs = append(newDropGroup.DropConfigs, newConfig)
			}
			dropGroup = newDropGroup
		}
		roleIdConfig := csvs.GetRandDropNew(dropGroup)
		if roleIdConfig != nil {
			roleConfig := csvs.GetRoleConfig(roleIdConfig.Result)
			if roleConfig != nil {
				if roleConfig.Star == 5 {
					m.UpPoolInfo.FiveStarTimes = 0
					fiveNum++
					if m.UpPoolInfo.IsMustUp == csvs.LoginTrue {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("数据异常")
								return
							}
						}
					}
					if roleIdConfig.DropId == 100012 {
						m.UpPoolInfo.IsMustUp = csvs.LoginFalse
					} else {
						m.UpPoolInfo.IsMustUp = csvs.LoginTrue
					}
				} else if roleConfig.Star == 4 {
					m.UpPoolInfo.FourStarTimes = 0
					fourNum++
				}
			}
			result[roleIdConfig.Result]++
			player.ModBag.AddItem(roleIdConfig.Result, 1, player)
		}
	}

	for k, v := range result {
		fmt.Println(fmt.Sprintf("抽中%s次数：%d", csvs.GetItemName(k), v))
	}
	fmt.Println(fmt.Sprintf("抽中4星：%d", fourNum))
	fmt.Println(fmt.Sprintf("抽中5星：%d", fiveNum))
}

func (m *ModPool) HandleUpPoolTimesTest(times int) {
	if times <= 0 || times > 100000000 {
		fmt.Println("请输入正确的数值(1~100000000)")
		return
	} else {
		fmt.Println(fmt.Sprintf("累计抽取%d次,结果如下:", times))
	}
	resultEach := make(map[int]int)
	for i := 0; i < times; i++ {
		m.AddTimes()
		dropGroup := csvs.ConfigDropGroupMap[1000]
		if dropGroup == nil {
			return
		}

		if m.UpPoolInfo.FiveStarTimes > csvs.FiveStarTimesLimit || m.UpPoolInfo.FourStarTimes > csvs.FourStarTimesLimit {
			newDropGroup := new(csvs.DropGroup)
			newDropGroup.DropId = dropGroup.DropId
			newDropGroup.WeightAll = dropGroup.WeightAll
			addFiveWeight := (m.UpPoolInfo.FiveStarTimes + 1 - csvs.FiveStarTimesLimit) * csvs.FiveStarTimesLimitEachValue
			if addFiveWeight < 0 {
				addFiveWeight = 0
			}
			addFourWeight := (m.UpPoolInfo.FourStarTimes + 1 - csvs.FourStarTimesLimit) * csvs.FourStarTimesLimitEachValue
			if addFourWeight < 0 {
				addFourWeight = 0
			}

			for _, config := range dropGroup.DropConfigs {
				newConfig := new(csvs.ConfigDrop)
				newConfig.Result = config.Result
				newConfig.DropId = config.DropId
				newConfig.IsEnd = config.IsEnd
				if config.Result == 10001 {
					newConfig.Weight = config.Weight + addFiveWeight
				} else if config.Result == 10002 {
					newConfig.Weight = config.Weight + addFourWeight
				} else if config.Result == 10003 {
					newConfig.Weight = config.Weight - addFiveWeight - addFourWeight
				}
				newDropGroup.DropConfigs = append(newDropGroup.DropConfigs, newConfig)
			}
			dropGroup = newDropGroup
		}
		roleIdConfig := csvs.GetRandDropNew(dropGroup)
		if roleIdConfig != nil {
			roleConfig := csvs.GetRoleConfig(roleIdConfig.Result)
			if roleConfig != nil {
				if roleConfig.Star == 5 {
					resultEach[m.UpPoolInfo.FiveStarTimes]++
					m.UpPoolInfo.FiveStarTimes = 0
					if m.UpPoolInfo.IsMustUp == csvs.LoginTrue {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("数据异常")
								return
							}
						}
					}
					if roleIdConfig.DropId == 100012 {
						m.UpPoolInfo.IsMustUp = csvs.LoginFalse
					} else {
						m.UpPoolInfo.IsMustUp = csvs.LoginTrue
					}
				} else if roleConfig.Star == 4 {
					m.UpPoolInfo.FourStarTimes = 0
				}
			}
		}
	}

	for k, v := range resultEach {
		fmt.Println(fmt.Sprintf("第%d抽抽出5星的次数：%d", k, v))
	}
}

func (m *ModPool) HandleUpPoolFiveTest(times int) {
	if times <= 0 || times > 100000000 {
		fmt.Println("请输入正确的数值(1~100000000)")
		return
	} else {
		fmt.Println(fmt.Sprintf("累计抽取%d次,结果如下:", times))
	}
	resultEachTest := make(map[int]int)
	fiveTest := 0
	for i := 0; i < times; i++ {
		m.AddTimes()
		if i%10 == 0 {
			fiveTest = 0
		}
		dropGroup := csvs.ConfigDropGroupMap[1000]
		if dropGroup == nil {
			return
		}

		if m.UpPoolInfo.FiveStarTimes > csvs.FiveStarTimesLimit || m.UpPoolInfo.FourStarTimes > csvs.FourStarTimesLimit {
			newDropGroup := new(csvs.DropGroup)
			newDropGroup.DropId = dropGroup.DropId
			newDropGroup.WeightAll = dropGroup.WeightAll
			addFiveWeight := (m.UpPoolInfo.FiveStarTimes - csvs.FiveStarTimesLimit) * csvs.FiveStarTimesLimitEachValue
			if addFiveWeight < 0 {
				addFiveWeight = 0
			}
			addFourWeight := (m.UpPoolInfo.FourStarTimes - csvs.FourStarTimesLimit) * csvs.FourStarTimesLimitEachValue
			if addFourWeight < 0 {
				addFourWeight = 0
			}

			for _, config := range dropGroup.DropConfigs {
				newConfig := new(csvs.ConfigDrop)
				newConfig.Result = config.Result
				newConfig.DropId = config.DropId
				newConfig.IsEnd = config.IsEnd
				if config.Result == 10001 {
					newConfig.Weight = config.Weight + addFiveWeight
				} else if config.Result == 10002 {
					newConfig.Weight = config.Weight + addFourWeight
				} else if config.Result == 10003 {
					newConfig.Weight = config.Weight - addFiveWeight - addFourWeight
				}
				newDropGroup.DropConfigs = append(newDropGroup.DropConfigs, newConfig)
			}
			dropGroup = newDropGroup
		}
		roleIdConfig := csvs.GetRandDropNew(dropGroup)
		if roleIdConfig != nil {
			roleConfig := csvs.GetRoleConfig(roleIdConfig.Result)
			if roleConfig != nil {
				if roleConfig.Star == 5 {
					fiveTest++
					m.UpPoolInfo.FiveStarTimes = 0
					if m.UpPoolInfo.IsMustUp == csvs.LoginTrue {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("数据异常")
								return
							}
						}
					}
					if roleIdConfig.DropId == 100012 {
						m.UpPoolInfo.IsMustUp = csvs.LoginFalse
					} else {
						m.UpPoolInfo.IsMustUp = csvs.LoginTrue
					}
				} else if roleConfig.Star == 4 {
					m.UpPoolInfo.FourStarTimes = 0
				}
			}
		}
		if i%10 == 9 {
			resultEachTest[fiveTest]++
		}
	}

	for k, v := range resultEachTest {
		fmt.Println(fmt.Sprintf("10连%d黄次数：%d", k, v))
	}
}

func (m *ModPool) HandleUpPoolSingleCheck1(times int, player *Player) {
	if times <= 0 || times > 100000000 {
		fmt.Println("请输入正确的数值(1~100000000)")
		return
	} else {
		fmt.Println(fmt.Sprintf("累计抽取%d次,结果如下:", times))
	}
	result := make(map[int]int)
	fourNum := 0
	fiveNum := 0
	for i := 0; i < times; i++ {
		m.AddTimes()
		dropGroup := csvs.ConfigDropGroupMap[1000]
		if dropGroup == nil {
			return
		}

		if m.UpPoolInfo.FiveStarTimes > csvs.FiveStarTimesLimit || m.UpPoolInfo.FourStarTimes > csvs.FourStarTimesLimit {
			newDropGroup := new(csvs.DropGroup)
			newDropGroup.DropId = dropGroup.DropId
			newDropGroup.WeightAll = dropGroup.WeightAll
			addFiveWeight := (m.UpPoolInfo.FiveStarTimes - csvs.FiveStarTimesLimit) * csvs.FiveStarTimesLimitEachValue
			if addFiveWeight < 0 {
				addFiveWeight = 0
			}
			addFourWeight := (m.UpPoolInfo.FourStarTimes - csvs.FourStarTimesLimit) * csvs.FourStarTimesLimitEachValue
			if addFourWeight < 0 {
				addFourWeight = 0
			}

			for _, config := range dropGroup.DropConfigs {
				newConfig := new(csvs.ConfigDrop)
				newConfig.Result = config.Result
				newConfig.DropId = config.DropId
				newConfig.IsEnd = config.IsEnd
				if config.Result == 10001 {
					newConfig.Weight = config.Weight + addFiveWeight
				} else if config.Result == 10002 {
					newConfig.Weight = config.Weight + addFourWeight
				} else if config.Result == 10003 {
					newConfig.Weight = config.Weight - addFiveWeight - addFourWeight
				}
				newDropGroup.DropConfigs = append(newDropGroup.DropConfigs, newConfig)
			}
			dropGroup = newDropGroup
		}
		// 获取角色信息，用于仓检功能，两个map记录玩家当前拥有角色情况
		fiveInfo, fourInfo := player.ModRole.GetRoleInfoForPoolCheck()
		roleIdConfig := csvs.GetRandDropNew1(dropGroup, fiveInfo, fourInfo)

		if roleIdConfig != nil {
			roleConfig := csvs.GetRoleConfig(roleIdConfig.Result)
			if roleConfig != nil {
				if roleConfig.Star == 5 {
					m.UpPoolInfo.FiveStarTimes = 0
					fiveNum++
					if m.UpPoolInfo.IsMustUp == csvs.LoginTrue {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("数据异常")
								return
							}
						}
					}
					if roleIdConfig.DropId == 100012 {
						m.UpPoolInfo.IsMustUp = csvs.LoginFalse
					} else {
						m.UpPoolInfo.IsMustUp = csvs.LoginTrue
					}
				} else if roleConfig.Star == 4 {
					m.UpPoolInfo.FourStarTimes = 0
					fourNum++
				}
			}
			weaponConfig := csvs.GetWeaponConfig(roleIdConfig.Result)
			if weaponConfig != nil {
				if weaponConfig.Star == 4 {
					m.UpPoolInfo.FourStarTimes = 0
					fourNum++
				}
			}
			result[roleIdConfig.Result]++
			player.ModBag.AddItem(roleIdConfig.Result, 1, player)
		}
	}

	for k, v := range result {
		fmt.Println(fmt.Sprintf("抽中%s次数：%d", csvs.GetItemName(k), v))
	}
	fmt.Println(fmt.Sprintf("抽中4星角色：%d", fourNum))
	fmt.Println(fmt.Sprintf("抽中5星：%d", fiveNum))
}

func (m *ModPool) HandleUpPoolSingleCheck2(times int, player *Player) {
	if times <= 0 || times > 100000000 {
		fmt.Println("请输入正确的数值(1~100000000)")
		return
	} else {
		fmt.Println(fmt.Sprintf("累计抽取%d次,结果如下:", times))
	}
	result := make(map[int]int)
	fourNum := 0
	fiveNum := 0
	for i := 0; i < times; i++ {
		m.AddTimes()
		dropGroup := csvs.ConfigDropGroupMap[1000]
		if dropGroup == nil {
			return
		}

		if m.UpPoolInfo.FiveStarTimes > csvs.FiveStarTimesLimit || m.UpPoolInfo.FourStarTimes > csvs.FourStarTimesLimit {
			newDropGroup := new(csvs.DropGroup)
			newDropGroup.DropId = dropGroup.DropId
			newDropGroup.WeightAll = dropGroup.WeightAll
			addFiveWeight := (m.UpPoolInfo.FiveStarTimes - csvs.FiveStarTimesLimit) * csvs.FiveStarTimesLimitEachValue
			if addFiveWeight < 0 {
				addFiveWeight = 0
			}
			addFourWeight := (m.UpPoolInfo.FourStarTimes - csvs.FourStarTimesLimit) * csvs.FourStarTimesLimitEachValue
			if addFourWeight < 0 {
				addFourWeight = 0
			}

			for _, config := range dropGroup.DropConfigs {
				newConfig := new(csvs.ConfigDrop)
				newConfig.Result = config.Result
				newConfig.DropId = config.DropId
				newConfig.IsEnd = config.IsEnd
				if config.Result == 10001 {
					newConfig.Weight = config.Weight + addFiveWeight
				} else if config.Result == 10002 {
					newConfig.Weight = config.Weight + addFourWeight
				} else if config.Result == 10003 {
					newConfig.Weight = config.Weight - addFiveWeight - addFourWeight
				}
				newDropGroup.DropConfigs = append(newDropGroup.DropConfigs, newConfig)
			}
			dropGroup = newDropGroup
		}

		fiveInfo, fourInfo := player.ModRole.GetRoleInfoForPoolCheck()
		roleIdConfig := csvs.GetRandDropNew2(dropGroup, fiveInfo, fourInfo)
		if roleIdConfig != nil {
			roleConfig := csvs.GetRoleConfig(roleIdConfig.Result)
			if roleConfig != nil {
				if roleConfig.Star == 5 {
					m.UpPoolInfo.FiveStarTimes = 0
					fiveNum++
					if m.UpPoolInfo.IsMustUp == csvs.LoginTrue {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("数据异常")
								return
							}
						}
					}
					if roleIdConfig.DropId == 100012 {
						m.UpPoolInfo.IsMustUp = csvs.LoginFalse
					} else {
						m.UpPoolInfo.IsMustUp = csvs.LoginTrue
					}
				} else if roleConfig.Star == 4 {
					m.UpPoolInfo.FourStarTimes = 0
					fourNum++
				}
			}
			weaponConfig := csvs.GetWeaponConfig(roleIdConfig.Result)
			if weaponConfig != nil {
				if weaponConfig.Star == 4 {
					m.UpPoolInfo.FourStarTimes = 0
					fourNum++
				}
			}
			result[roleIdConfig.Result]++
			player.ModBag.AddItem(roleIdConfig.Result, 1, player)
		}
	}

	for k, v := range result {
		fmt.Println(fmt.Sprintf("抽中%s次数：%d", csvs.GetItemName(k), v))
	}
	fmt.Println(fmt.Sprintf("抽中4星角色：%d", fourNum))
	fmt.Println(fmt.Sprintf("抽中5星：%d", fiveNum))
}
