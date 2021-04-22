package twodeminsion

import (
	"testing"
)

var bkSets = []struct {
	name string
	path string
	kp   *KnapSack
	its  Items
	slt  int
}{
	{name: "set05", path: "./testdata/"},
	{name: "set20", path: "./testdata/"},
	{name: "set40", path: "./testdata/"},
	// {name: "set100", path: "./testdata/"},
	// {name: "set500", path: "./testdata/"},
}

//Load data from the datasetNames befrore running the test.
func init() {
	for i, set := range bkSets {
		bkSets[i].kp, bkSets[i].its, bkSets[i].slt = GetSet2D(set.path, set.name)

	}
}

func TestBacktrack(t *testing.T) {
	for _, set := range bkSets {
		if got := Backtrack(set.its, set.kp); got != set.slt {
			t.Errorf("Want %d, got %d\n", set.slt, got)
		}
	}
}
