// The bounded knapsack problem (BKP) places  upper bound
// on the number of copies of each kind of item and the
// only restriction on each item is that its number is a n.
// on-negative integer less than upper bund.
package boundedKnapsack

import "github.com/CausingBrick/knapsack"

// Dynamic dynamic programing solution.
// Transfer equaltion is as bellow:
// 		dp[i+1][j] = max(dp[i][j-K*w[i]]+k*v[i]) and the 0 <= k <= b[j]
func Dynamic(items []*knapsack.Item, capacity int, bound []int) int {
	num := len(items)
	dp := make([][]int, num+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 0; i < num; i++ {
		for j := 0; j <= capacity; j++ {
			for k := 0; k <= min(j/items[i].Weight, bound[i]); k++ {
				dp[i+1][j] = max(dp[i+1][j],
					dp[i][j-k*items[i].Weight]+k*items[i].Value)
			}
		}
	}
	return dp[num][capacity]
}

// DynamicCompress dynamic programing soulution with compressed dp table.
// The transfer equation is as follow:
// 		 dp[j] = max(dp[j], dp[j-w[i]]*k+v[i]*k)
func DynamicCompress(items []*knapsack.Item, capacity int, bound []int) int {
	num := len(items)
	dp := make([]int, capacity+1)

	for i := 0; i < num; i++ {
		for j := capacity; j >= items[i].Weight; j-- {
			for k := 1; k <= bound[i]; k++ {
				if j-k*items[i].Weight >= 0 {
					dp[j] = max(dp[j], dp[j-k*items[i].Weight]+k*items[i].Value)
				}
			}
		}
	}
	return dp[capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
