package knapsack

import (
	"testing"
)

var (
	pft, wgt []int
	cap      int
	want     int
)

func init() {
	pft, wgt, cap, want = getProfits(), getWeights(), getCapacity(), getBestprofit()
}

func TestKnapsack01Weight(t *testing.T) {
	got := Knapsack01Weight(pft, wgt, cap)
	t.Logf("\nBest profit is:%d\nMaxmium value is:%d\n", want, got)

}
func TestKnapsack01Profit(t *testing.T) {
	got := Knapsack01Profit(pft, wgt, cap)
	t.Logf("\nBest profit is:%d\nMaxmium value is:%d\n", want, got)

}

func TestKnapsack01Ratio(t *testing.T) {
	got := Knapsack01Ratio(pft, wgt, cap)
	t.Logf("\nBest profit is:%d\nMaxmium value is:%d\n", want, got)

}

func TestKnapsack01Recursive(t *testing.T) {
	got := Knapsack01Recursive(pft, wgt, cap)
	if got != want {
		t.Errorf("Knapsack01Recursive(%v,%v,%d) return %d, want %d", pft, wgt, cap, got, want)
	}
	t.Logf("\nBest profit is:%d\nMaxmium value is:%d\n", want, got)
}

func TestKnapsack01Dynamic(t *testing.T) {
	got := Knapsack01Dynamic(pft, wgt, cap)
	if got != want {
		t.Errorf("TestKnapsack01Dynamic(%v,%v,%d) return %d, want %d", pft, wgt, cap, got, want)
	}
	t.Logf("\nBest profit is:%d\nMaxmium value is:%d\n", want, got)
}

func TestKnapsack01DynamicLoop(t *testing.T) {
	got := Knapsack01DynamicLoop(pft, wgt, cap)
	if got != want {
		t.Errorf("TestKnapsack01DynamicLoop(%v,%v,%d) return %d, want %d", pft, wgt, cap, got, want)
	}
	t.Logf("\nBest profit is:%d\nMaxmium value is:%d\n", want, got)
}
func TestKnapsack01DynamicLoopOpt(t *testing.T) {
	got := Knapsack01DynamicLoopOpt(pft, wgt, cap)
	if got != want {
		t.Errorf("TestKnapsack01DynamicLoop(%v,%v,%d) return %d, want %d", pft, wgt, cap, got, want)
	}
	t.Logf("\nBest profit is:%d\nMaxmium value is:%d\n", want, got)
}

func TestKnapsack01DynamicHash(t *testing.T) {
	got := Knapsack01DynamicHash(pft, wgt, cap)
	if got != want {
		t.Errorf("Knapsack01DynamicHash(%v,%v,%d) return %d, want %d", pft, wgt, cap, got, want)
	}
	t.Logf("\nBest profit is:%d\nMaxmium value is:%d\n", want, got)
}

func BenchmarkKnapsack01Weight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Knapsack01Weight(pft, wgt, cap)
	}
}

func BenchmarkKnapsack01Profit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Knapsack01Profit(pft, wgt, cap)
	}
}
func BenchmarkKnapsack01Ratio(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Knapsack01Ratio(pft, wgt, cap)
	}
}
func BenchmarkKnapsack01Recursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Knapsack01Recursive(pft, wgt, cap)
	}
}

func BenchmarkKnapsack01Dynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Knapsack01Dynamic(pft, wgt, cap)
	}
}
func BenchmarkKnapsack01DynamicLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Knapsack01DynamicLoop(pft, wgt, cap)
	}
}

func BenchmarkKnapsack01DynamicLoopOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Knapsack01DynamicLoopOpt(pft, wgt, cap)
	}
}
func BenchmarkKnapsack01DynamicHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Knapsack01DynamicHash(pft, wgt, cap)
	}
}
