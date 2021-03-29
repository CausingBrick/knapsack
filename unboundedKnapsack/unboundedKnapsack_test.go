package unboundedKnapsack

import (
	"testing"

	"github.com/CausingBrick/knapsack"
	"github.com/CausingBrick/knapsack/datasets"
)

type testTable struct {
	kp     *knapsack.KnapSack
	dsName string
}

// Add test data from here.
// Key: Name for dataset, Value: path + set name
var ds = map[string]string{
	"KNAPSACK_01_P01": "../datasets/KNAPSACK_01/p01",
	"KNAPSACK_01_P02": "../datasets/KNAPSACK_01/p02",
	"KNAPSACK_01_P03": "../datasets/KNAPSACK_01/p03",
}

var testTables []*testTable

//Load data from the datasetNames befrore running the test.
func init() {
	for k, v := range ds {
		testTables = append(testTables, &testTable{kp: datasets.ReadDataset(v), dsName: k})
	}
}

func TestDynamic(t *testing.T) {
	for _, table := range testTables {
		t.Run(table.dsName, func(t *testing.T) {
			got := Dynamic(table.kp.Items, table.kp.Capacity)
			t.Log("\nDataset:", table.dsName, "Got:", got)
		})
	}
}

func TestDynamicOpt(t *testing.T) {
	for _, table := range testTables {
		t.Run(table.dsName, func(t *testing.T) {
			got := DynamicOpt(table.kp.Items, table.kp.Capacity)
			t.Log("\nDataset:", table.dsName, "Got:", got)
		})
	}
}

func TestDynamicCompress(t *testing.T) {
	for _, table := range testTables {
		t.Run(table.dsName, func(t *testing.T) {
			got := DynamicCompress(table.kp.Items, table.kp.Capacity)
			t.Log("\nDataset:", table.dsName, "Got:", got)
		})
	}
}

func TestSGA(t *testing.T) {
	conf := &SGAConf{
		PopulationSize:   10,
		CompetitionScale: 2,
		mutateRatio:      0.9,
		Generations:      3,
	}
	for _, table := range testTables {
		t.Run(table.dsName, func(t *testing.T) {
			got := SGA(table.kp.Items, table.kp.Capacity, conf)
			t.Log("\nDataset:", table.dsName, "Got:", got[len(got)-1])
		})
	}
}

func BenchmarkDynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dynamic(testTables[0].kp.Items, testTables[0].kp.Capacity)
	}
}
func BenchmarkDynamicOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DynamicOpt(testTables[0].kp.Items, testTables[0].kp.Capacity)
	}
}

func BenchmarkDynamicCompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DynamicCompress(testTables[0].kp.Items, testTables[0].kp.Capacity)
	}
}
