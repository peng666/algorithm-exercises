package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
* 原博客地址
* https://blog.csdn.net/qq_37102984/article/details/129015565
 */

// RewardInfo 奖励信息
type RewardInfo struct {
	RewardId   int32   `json:"-"`          // 奖品id
	RewardName string  `json:"RewardName"` // 奖品名称
	Weight     float32 `json:"Weight"`     // 抽中的权重（省略末尾百分号的形式：例如某奖品权重为 0.1%，那么这里值也为 0.1）
	GetCnt     int32   `json:"GetCnt"`     // 中奖数
}

// 抽奖模拟器
// 入参：chouJiangCnt-用户的抽取次数
func LotterySimulator(chouJiangCnt int) {
	// 以下权重都是百分制（省略了末尾百分号%），例如 0.1 → 0.1%，79 → 79%，一共是100%
	infos := []RewardInfo{
		{RewardId: 1, RewardName: "幻神-CFS2022(皮肤,不可交易)", Weight: 0.1},
		{RewardId: 2, RewardName: "火麒麟-巅峰荣耀(皮肤,不可交易)", Weight: 0.15},
		{RewardId: 3, RewardName: "黑骑士-CFS2021(皮肤,不可交易)", Weight: 0.25},
		{RewardId: 4, RewardName: "毁灭-CFS2020(皮肤,不可交易)", Weight: 0.25},
		{RewardId: 5, RewardName: "M4A1-雷神-CFS2019(皮肤,不可交易)", Weight: 0.3},
		{RewardId: 6, RewardName: "毁灭-CFS2018(皮肤,不可交易)", Weight: 0.3},
		{RewardId: 7, RewardName: "火麒麟-CFS2018(皮肤,不可交易)", Weight: 0.3},
		{RewardId: 8, RewardName: "屠龙-荣耀之锋(皮肤,不可交易)", Weight: 0.4},
		{RewardId: 9, RewardName: "修罗-CFS2018(皮肤,不可交易)", Weight: 0.85},
		{RewardId: 10, RewardName: "王者之石x1", Weight: 7},
		{RewardId: 11, RewardName: "积分x100", Weight: 0.1},
		{RewardId: 12, RewardName: "积分x10", Weight: 1},
		{RewardId: 13, RewardName: "积分x5", Weight: 10},
		{RewardId: 14, RewardName: "积分x1", Weight: 79},
	}

	// 初始化奖品属性map
	rewardMap := make(map[int32]RewardInfo)
	for _, rewardInfo := range infos {
		rewardMap[rewardInfo.RewardId] = RewardInfo{
			RewardId:   rewardInfo.RewardId,
			RewardName: rewardInfo.RewardName,
			Weight:     rewardInfo.Weight,
			GetCnt:     0,
		}
	}

	// 此处按需添加：将抽奖结果记录文件于当前路径下
	// 注意：我下面是判断抽到幻神才输出结果记录到文件
	testJsonFile, _ := os.OpenFile("./cf_reward.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	defer testJsonFile.Close()

	// 当前抽奖已花费总金额，初始为0；单次抽奖花费金额10元
	costMoneyOfTotal, costMoneyOfEveryTime := 0, 10
	for {
		time.Sleep(time.Second) // 每秒抽一次，控制下速率，方便控制台打印观察
		fmt.Printf("抽奖中...，当前已消费%v元\n", costMoneyOfTotal)

		for i := 0; i < chouJiangCnt; i++ {
			rewardId := SelectByWeight(infos)
			// 组装抽取奖品的数据
			rewardMap[rewardId] = RewardInfo{
				RewardId:   rewardMap[rewardId].RewardId,
				RewardName: rewardMap[rewardId].RewardName,
				Weight:     rewardMap[rewardId].Weight,
				GetCnt:     rewardMap[rewardId].GetCnt + 1,
			}
			costMoneyOfTotal += costMoneyOfEveryTime

			// 每次抽奖都有一定间隔，避免运行过快导致一瞬间的随机数种子都一致
			time.Sleep(time.Nanosecond)
		}

		// 以下逻辑按需添加，只是输出指定抽奖结果到文件
		// 如果 幻神数 > 0，则把抽奖结果记录到文件中，并停止抽奖
		if rewardMap[1].GetCnt > 0 {
			jsonRewardMap, _ := json.Marshal(rewardMap)
			// fmt.Printf("rewardMap: \n%v\n", string(jsonRewardMap))

			writer := bufio.NewWriter(testJsonFile)
			writer.WriteString(string(jsonRewardMap) + "\n" + fmt.Sprintf("一共花费RMB：%v", costMoneyOfTotal))
			writer.Flush()
			break
		}
	}

	return
}

// SelectByWeight 根据不同奖品权重随机抽取奖品并返回
func SelectByWeight(infos []RewardInfo) (RewardId int32) {
	// 先找出所有奖品数要乘的最大公共倍数 maxMultiple
	var maxMultiple int = 1

	for i := 0; i < len(infos); i++ {
		strWeight := fmt.Sprintf("%v", infos[i].Weight)
		// maxMultiple 只关注小数类型权重值，整数类型权重值直接跳过。例如"0.25"，我们通过其中的"25"来得到 maxMultiple=100
		if i := strings.Index(strWeight, "."); i == -1 {
			continue
		}

		var tmpMultiple int = 1
		for i := 0; i < len(strWeight[i+1:]); i++ {
			tmpMultiple *= 10
		}

		if maxMultiple < tmpMultiple {
			maxMultiple = tmpMultiple
		}
	}
	// fmt.Println("最大倍数 maxMultiple: ", maxMultiple)

	// 避免反复扩容，提前预估cap容量：倍率*100
	// 例如权重为0.85，乘以100倍率后为85；那么对于总权重为100，乘以100倍率后就为1w
	rewardList := make([]int32, 0, maxMultiple*100)

	// 根据不同奖品的权重值，分配对应个数的奖品到奖池中
	for _, info := range infos {
		for i := 0; i < int(info.Weight*float32(maxMultiple)); i++ { // 填充每个奖品的整数样本个数到容量为1w的奖池中
			rewardList = append(rewardList, info.RewardId)
		}
	}
	// fmt.Printf("奖池奖品总数 = %v, 奖池奖品明细 = %v \n", len(rewardList), rewardList)
	// // os.Exit(0)

	rand.Seed(time.Now().UnixNano())              // UnixNano将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位纳秒）
	return rewardList[rand.Intn(len(rewardList))] // rand.Intn(100)类似于rand.Int()%100
}

func main() {
	// 模拟cf抽奖
	LotterySimulator(10)
}
