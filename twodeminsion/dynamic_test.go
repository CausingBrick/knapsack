package twodeminsion

import (
	"testing"
)

var dynamicSets = []struct {
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
	for i, set := range dynamicSets {
		dynamicSets[i].kp, dynamicSets[i].its, dynamicSets[i].slt = GetSet2D(set.path, set.name)
	}
}

func TestDynamicRec(t *testing.T) {
	for _, set := range dynamicSets {
		if got := DynamicRec(set.its, set.kp); got != set.slt {
			t.Errorf("Want %d, got %d\n", set.slt, got)
		}
	}
}

// func TestDynamic(t *testing.T) {
// 	for k, v := range dynamicSets {
// 		log.Println(k, Dynamic(v.its, v.kp))
// 	}
// }

func TestDynamicCompress(t *testing.T) {
	for _, set := range dynamicSets {
		if got := DynamicCompress(set.its, set.kp); got != set.slt {
			t.Errorf("Want %d, got %d\n", set.slt, got)
		}
	}
}
