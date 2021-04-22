package twodeminsion

import (
	"testing"
)

var bkSets = []*dataset{
	datasetNew("set05", "./testdata/"),
	datasetNew("set20", "./testdata/"),
	// datasetNew("set40", "./testdata/"),
	// datasetNew("set100", "./testdata/"),
}

func TestBacktrack(t *testing.T) {
	for _, set := range bkSets {
		if got := Backtrack(set.its, set.kp); got != set.slt {
			t.Errorf("Want %d, got %d\n", set.slt, got)
		}
	}
}
