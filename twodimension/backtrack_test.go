package twodimension

import (
	"testing"
)

var bkSets = map[string]*dataset{
	"set05": datasetNew("set05", "./testdata/"),
	"set20": datasetNew("set20", "./testdata/"),
}

func TestBacktrack(t *testing.T) {
	for _, set := range bkSets {
		set := set //Capturing iterator variables
		t.Run(set.name, func(t *testing.T) {
			if got := Backtrack(set.its, set.kp); got != set.slt {
				t.Parallel()
				t.Errorf("Set %s Want %d, got %d\n", set.name, set.slt, got)
			}
		})
	}
}
