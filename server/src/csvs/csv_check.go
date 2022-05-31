package csvs

import (
	"fmt"
	"math/rand"
)

// CheckLoadCsv 通常建立一个check函数，来表示初始化的结束
// CheckLoadCsv 初始化完成之后，表示结束

var (
	ConfigDropGroupMap map[int]*DropGroup
)

type DropGroup struct {
	DropId      int
	WeightAll   int // 权重总和
	DropConfigs []*ConfigDrop
}

func CheckLoadCsv() {
	//二次处理
	MakeDropGroupMap()
	fmt.Println("csv配置读取完成---ok")
}

// MakeDropGroupMap 生成掉落组的map
func MakeDropGroupMap() {
	ConfigDropGroupMap = make(map[int]*DropGroup)
	// 遍历这个map，读取响应配置即可
	for _, v := range ConfigDropSlice {
		dropGroup, ok := ConfigDropGroupMap[v.DropId]
		if !ok {
			dropGroup = new(DropGroup)
			dropGroup.DropId = v.DropId
			ConfigDropGroupMap[v.DropId] = dropGroup
		}
		// 权重总和
		dropGroup.WeightAll += v.Weight
		dropGroup.DropConfigs = append(dropGroup.DropConfigs, v)
	}
	//RandDropTest()
	return
}

// RandDropTest 随机掉落测试
func RandDropTest() {
	dropGroup := ConfigDropGroupMap[1000]
	if dropGroup == nil {
		return
	}
	num := 0
	for {
		config := GetRandDropNew(dropGroup)
		if config.IsEnd == LoginTrue {
			fmt.Println(GetItemName(config.Result))
			num++
			dropGroup = ConfigDropGroupMap[1000]
			if num >= 100 {
				break
			} else {
				continue
			}
		}
		dropGroup = ConfigDropGroupMap[config.Result]
		if dropGroup == nil {
			break
		}
	}
}

// GetRandDrop 获得随机掉落物，返回掉落物配置信息
func GetRandDrop(dropGroup *DropGroup) *ConfigDrop {
	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			return v
		}
	}
	return nil
}

// GetRandDropNew 使用递归
func GetRandDropNew(dropGroup *DropGroup) *ConfigDrop {
	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			if v.IsEnd == LoginTrue {
				// 结束的话直接返回
				return v
			}
			dropGroup := ConfigDropGroupMap[v.Result]
			// 增加容错
			if dropGroup == nil {
				return nil
			}
			// 将参数刷新为递归后的参数
			return GetRandDropNew(dropGroup)
		}
	}
	return nil
}
func GetRandDropNew1(dropGroup *DropGroup, fiveInfo map[int]int, fourInfo map[int]int) *ConfigDrop {
	for _, v := range dropGroup.DropConfigs {
		_, ok := fiveInfo[v.Result]
		if ok {
			index := 0
			maxGetTime := 0
			for k, config := range dropGroup.DropConfigs {
				_, nowOK := fiveInfo[config.Result]
				if !nowOK {
					continue
				}
				if maxGetTime < fiveInfo[config.Result] {
					maxGetTime = fiveInfo[config.Result]
					index = k
				}
			}
			return dropGroup.DropConfigs[index]
		}

		_, ok = fourInfo[v.Result]
		if ok {
			index := 0
			maxGetTime := 0
			for k, config := range dropGroup.DropConfigs {
				_, nowOK := fourInfo[config.Result]
				if !nowOK {
					continue
				}
				if maxGetTime < fourInfo[config.Result] {
					maxGetTime = fourInfo[config.Result]
					index = k
				}
			}
			return dropGroup.DropConfigs[index]
		}
	}

	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			if v.IsEnd == LoginTrue {
				return v
			}
			dropGroup := ConfigDropGroupMap[v.Result]
			if dropGroup == nil {
				return nil
			}
			return GetRandDropNew1(dropGroup, fiveInfo, fourInfo)
		}
	}
	return nil
}

func GetRandDropNew2(dropGroup *DropGroup, fiveInfo map[int]int, fourInfo map[int]int) *ConfigDrop {
	for _, v := range dropGroup.DropConfigs {
		_, ok := fiveInfo[v.Result]
		if ok {
			index := 0
			minGetTime := 0
			for k, config := range dropGroup.DropConfigs {
				_, nowOK := fiveInfo[config.Result]
				if !nowOK {
					index = k
					break
				}
				if minGetTime == 0 || minGetTime > fiveInfo[config.Result] {
					minGetTime = fiveInfo[config.Result]
					index = k
				}
			}
			return dropGroup.DropConfigs[index]
		}

		_, ok = fourInfo[v.Result]
		if ok {
			index := 0
			minGetTime := 0
			for k, config := range dropGroup.DropConfigs {
				_, nowOK := fourInfo[config.Result]
				if !nowOK {
					index = k
					break
				}
				if minGetTime == 0 || minGetTime > fourInfo[config.Result] {
					minGetTime = fourInfo[config.Result]
					index = k
				}
			}
			return dropGroup.DropConfigs[index]
		}
	}

	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			if v.IsEnd == LoginTrue {
				return v
			}
			dropGroup := ConfigDropGroupMap[v.Result]
			if dropGroup == nil {
				return nil
			}
			return GetRandDropNew2(dropGroup, fiveInfo, fourInfo)
		}
	}
	return nil
}
