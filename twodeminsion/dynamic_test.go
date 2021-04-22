package twodeminsion

import (
	"testing"
)

var dynamicSets = []*dataset{
	datasetNew("set05", "./testdata/"),
	datasetNew("set20", "./testdata/"),
	datasetNew("set40", "./testdata/"),
	datasetNew("set100", "./testdata/"),
}

func TestDynamicRec(t *testing.T) {
	for _, set := range dynamicSets {
		if got := DynamicRec(set.its, set.kp); got != set.slt {
			t.Errorf("Set %s Want %d, got %d\n", set.name, set.slt, got)
		}
	}
}

func TestDynamic(t *testing.T) {
	for _, set := range dynamicSets {
		if got := Dynamic(set.its, set.kp); got != set.slt {
			t.Errorf("Set %s Want %d, got %d\n", set.name, set.slt, got)
		}
	}
}

func TestDynamicCompress(t *testing.T) {
	for _, set := range dynamicSets {
		if got := DynamicCompress(set.its, set.kp); got != set.slt {
			t.Errorf("Set %s Want %d, got %d\n", set.name, set.slt, got)
		}
	}
}
