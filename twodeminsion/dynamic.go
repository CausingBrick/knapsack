package twodeminsion

import "sync"

func max(nums ...int) int {
	m := nums[0]
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}

// DynamicRec solution for the two-dimensional cost knapsack problem.
func DynamicRec(items []*Item, kp *KnapSack) int {
	dp := make([][][]int, len(items)+1)
	for i := range dp {
		dp[i] = make([][]int, kp.Capacity+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, kp.Volume+1)
		}
	}
	return dynamicRec(dp, items, kp.Capacity, kp.Volume, 0)
}

func dynamicRec(dp [][][]int, items []*Item, cap, vol, i int) int {
	value := 0
	switch {
	case dp[i][cap][vol] > 0:
		value = dp[i][cap][vol]
	case i == len(items):
		break
	case items[i].Bulk > vol || items[i].Weight > cap:
		value = dynamicRec(dp, items, cap, vol, i+1)
		dp[i][cap][vol] = value
	default:
		value = max(
			dynamicRec(dp, items, cap, vol, i+1),
			dynamicRec(
				dp,
				items,
				cap-items[i].Weight,
				vol-items[i].Bulk,
				i+1,
			)+items[i].Value,
		)
		dp[i][cap][vol] = value
	}
	return value
}

// Dynamic dynamic programing solution.
func Dynamic(items []*Item, kp *KnapSack) int {
	dp := make([][][]int, len(items)+1)
	var wg sync.WaitGroup
	wg.Add(len(items) + 1)
	for i := range dp {
		go func(i int) {
			dp[i] = make([][]int, kp.Capacity+1)
			for j := range dp[i] {
				dp[i][j] = make([]int, kp.Volume+1)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i, item := range items {
		for j := 0; j <= kp.Capacity; j++ {
			for k := 0; k <= kp.Volume; k++ {
				if item.Weight > j || item.Bulk > k {
					dp[i+1][j][k] = dp[i][j][k]
				} else {
					dp[i+1][j][k] = max(
						dp[i][j][k],
						dp[i][j-item.Weight][k-item.Bulk]+item.Value,
					)
				}
			}
		}
	}
	return dp[len(items)][kp.Capacity][kp.Volume]
}

// DynamicCompress loop soulution with two-dimensional array.
func DynamicCompress(items []*Item, kp *KnapSack) int {
	dp := make([][]int, kp.Capacity+1)
	for i := range dp {
		dp[i] = make([]int, kp.Volume+1)
	}

	for _, item := range items {
		for j := kp.Capacity; j >= item.Weight; j-- {
			for k := kp.Volume; k >= item.Bulk; k-- {
				dp[j][k] = max(
					dp[j][k],
					dp[j-item.Weight][k-item.Bulk]+item.Value,
				)
			}
		}
	}

	return dp[kp.Capacity][kp.Volume]
}
