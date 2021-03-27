// The unbounded knapsack problem (UKP) places no upper bound
// on the number of copies of each kind of item and
// the only restriction on  is that it is a non-negative integer.
package unboundedKnapsack

import (
	common "github.com/CausingBrick/knapsack"
)

// DP dynamic programing solution.
// Transfer equaltion is as bellow:
// 		dp[i+1][j] = max(dp[i][j-K*w[i]]+k*v[i]) and the k >= 0
func DP(items []*common.Item, capacity int) int {
	num := len(items)
	dp := make([][]int, num+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 0; i < num; i++ {
		for j := 0; j <= capacity; j++ {
			maxV := 0
			//Found the number k of items we need.
			for k := 0; ; k++ {
				sumW, sumV := k*items[i].Weight, k*items[i].Value
				if sumW <= j {
					if maxV < sumV {
						maxV = sumV
					}
				} else { //case: K*w[i] > j
					break
				}
			}
			dp[i+1][j] = maxV
		}
	}
	return dp[num][capacity]
}

// DPOpt dynamic programing soulution with optimize the number of iterations.
// The transfer equation is as follow:
// 		dp[i+1][j] = max(dp[i][j],dp[i+1][j-w[i]]+v[i])
func DPOpt(items []*common.Item, capacity int) int {
	num := len(items)
	dp := make([][]int, num+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 0; i < num; i++ {
		for j := 0; j <= capacity; j++ {
			if items[i].Weight > j {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = common.Max(dp[i][j], dp[i+1][j-items[i].Weight]+items[i].Value)
			}
		}
	}
	return dp[num][capacity]
}

// DPCompress dynamic programing soulution with compressed dp table.
// The transfer equation is as follow:
// 		dp[j] = max(dp[j], dp[j-w[i]]+v[i])
func DPCompress(items []*common.Item, capacity int) int {
	num := len(items)
	dp := make([]int, capacity+1)

	for i := 0; i < num; i++ {
		for j := items[i].Weight; j <= capacity; j++ {
			if s := dp[j-items[i].Weight] + items[i].Value; dp[j] < s {
				dp[j] = s
			}
		}
	}
	return dp[capacity]
}
