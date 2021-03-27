package knapsack01

import (
	"fmt"
	"testing"

	ds "github.com/CausingBrick/knapsack/datasets"
)

type testTable struct {
	*ds.KnapSack01
	dsName string
}

// Add test data from here.
// The keys are the path of dataset, the value are the set names.
var datasetNames = map[string][]string{
	"../datasets/KNAPSACK_01": {"p01" /*, "p02", "p03", "p04", "p05"*/},
}

var testTables []*testTable

//Load data from the datasetNames befrore running the test.
func init() {
	for key, names := range datasetNames {
		for _, name := range names {
			testTables = append(testTables, &testTable{ds.KnapSack01New(key, name), name})
		}
	}
	fmt.Println(testTables)
}

func TestKnapsack01Weight(t *testing.T) {
	t.Parallel()
	for _, tt := range testTables {
		t.Run(tt.dsName, func(t *testing.T) {
			t.Parallel()
			got := GreedyWeight(tt.Profits, tt.Weights, tt.Capacity)
			t.Log("Got:", got, "Expected:", tt.BestProfit)
		})
	}
}

func TestKnapsack01Profit(t *testing.T) {
	t.Parallel()
	for _, tt := range testTables {
		t.Run(tt.dsName, func(t *testing.T) {
			t.Parallel()
			got := Knapsack01Profit(tt.Profits, tt.Weights, tt.Capacity)
			t.Log("Got:", got, "Expected:", tt.BestProfit)
		})
	}
}

func TestKnapsack01Ratio(t *testing.T) {
	t.Parallel()
	for _, tt := range testTables {
		t.Run(tt.dsName, func(t *testing.T) {
			t.Parallel()
			got := Knapsack01Ratio(tt.Profits, tt.Weights, tt.Capacity)
			t.Log("Got:", got, "Expected:", tt.BestProfit)
		})
	}
}

func TestKnapsack01Recursive(t *testing.T) {
	t.Parallel()
	for _, tt := range testTables {
		t.Run(tt.dsName, func(t *testing.T) {
			t.Parallel()
			got := Backtrack(tt.Profits, tt.Weights, tt.Capacity)
			t.Log("Got:", got, "Expected:", tt.BestProfit)
			if got != tt.BestProfit {
				t.Errorf("Error in \033[1;31;40m%s\033[0m, got: %d, but want: %d\n", tt.dsName, got, tt.BestProfit)
			}
		})
	}
}

func TestKnapsack01Dynamic(t *testing.T) {
	t.Parallel()
	for _, tt := range testTables {
		t.Run(tt.dsName, func(t *testing.T) {
			t.Parallel()
			got := DynamicRecursive(tt.Profits, tt.Weights, tt.Capacity)
			t.Log("Got:", got, "Expected:", tt.BestProfit)
			if got != tt.BestProfit {
				t.Errorf("Error in %s, got: %d, but want: %d\n", tt.dsName, got, tt.BestProfit)
			}
		})
	}
}

func TestKnapsack01DynamicLoop(t *testing.T) {
	t.Parallel()
	for _, tt := range testTables {
		t.Run(tt.dsName, func(t *testing.T) {
			t.Parallel()
			got := Dynamic(tt.Profits, tt.Weights, tt.Capacity)
			t.Log("Got:", got, "Expected:", tt.BestProfit)
			if got != tt.BestProfit {
				t.Errorf("Error in %s, got: %d, but want: %d\n", tt.dsName, got, tt.BestProfit)
			}
		})
	}
}
func TestKnapsack01DynamicLoopOpt(t *testing.T) {
	t.Parallel()
	for _, tt := range testTables {
		t.Run(tt.dsName, func(t *testing.T) {
			t.Parallel()
			got := DynamicCompress(tt.Profits, tt.Weights, tt.Capacity)
			t.Log("Got:", got, "Expected:", tt.BestProfit)
			if got != tt.BestProfit {
				t.Errorf("Error in %s, got: %d, but want: %d\n", tt.dsName, got, tt.BestProfit)
			}
		})
	}
}

func TestKnapsack01DynamicHash(t *testing.T) {
	t.Parallel()
	for _, tt := range testTables {
		t.Run(tt.dsName, func(t *testing.T) {
			t.Parallel()
			got := DynamicHash(tt.Profits, tt.Weights, tt.Capacity)
			t.Log("Got:", got, "Expected:", tt.BestProfit)
			if got != tt.BestProfit {
				t.Errorf("Error in %s, got: %d, but want: %d\n", tt.dsName, got, tt.BestProfit)
			}
		})
	}
}

func BenchmarkKnapsack01Weight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreedyWeight(testTables[0].Profits, testTables[0].Weights, testTables[0].Capacity)
	}
}

func BenchmarkKnapsack01Profit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Knapsack01Profit(testTables[0].Profits, testTables[0].Weights, testTables[0].Capacity)
	}
}
func BenchmarkKnapsack01Ratio(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Knapsack01Ratio(testTables[0].Profits, testTables[0].Weights, testTables[0].Capacity)
	}
}
func BenchmarkKnapsack01Recursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Backtrack(testTables[0].Profits, testTables[0].Weights, testTables[0].Capacity)
	}
}

func BenchmarkKnapsack01Dynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DynamicRecursive(testTables[0].Profits, testTables[0].Weights, testTables[0].Capacity)
	}
}
func BenchmarkKnapsack01DynamicLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dynamic(testTables[0].Profits, testTables[0].Weights, testTables[0].Capacity)
	}
}

func BenchmarkKnapsack01DynamicLoopOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DynamicCompress(testTables[0].Profits, testTables[0].Weights, testTables[0].Capacity)
	}
}
func BenchmarkKnapsack01DynamicHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DynamicHash(testTables[0].Profits, testTables[0].Weights, testTables[0].Capacity)
	}
}
