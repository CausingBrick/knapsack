package twodimension

import (
	"testing"
)

var dynSets = map[string]*dataset{
	"set05":  datasetNew("set05", "./testdata/"),
	"set20":  datasetNew("set20", "./testdata/"),
	"set40":  datasetNew("set40", "./testdata/"),
	"set100": datasetNew("set100", "./testdata/"),
	"set250": datasetNew("set250", "./testdata/"),
	"set500": datasetNew("set500", "./testdata/"),
}

func TestDynamicRec(t *testing.T) {
	t.Parallel()
	t.Run(dynSets["set05"].name, func(t *testing.T) {
		if got := DynamicRec(dynSets["set05"].its, dynSets["set05"].kp); got != dynSets["set05"].slt {
			t.Parallel()
			t.Errorf("Set %s Want %d, got %d\n", dynSets["set05"].name, dynSets["set05"].slt, got)
		}
	})
	t.Run(dynSets["set20"].name, func(t *testing.T) {
		if got := DynamicRec(dynSets["set20"].its, dynSets["set20"].kp); got != dynSets["set20"].slt {
			t.Parallel()
			t.Errorf("Set %s Want %d, got %d\n", dynSets["set20"].name, dynSets["set20"].slt, got)
		}
	})
	t.Run(dynSets["set40"].name, func(t *testing.T) {
		if got := DynamicRec(dynSets["set40"].its, dynSets["set40"].kp); got != dynSets["set40"].slt {
			t.Parallel()
			t.Errorf("Set %s Want %d, got %d\n", dynSets["set40"].name, dynSets["set40"].slt, got)
		}
	})
}

func TestDynamic(t *testing.T) {
	t.Parallel()
	t.Run(dynSets["set05"].name, func(t *testing.T) {
		if got := Dynamic(dynSets["set05"].its, dynSets["set05"].kp); got != dynSets["set05"].slt {
			t.Parallel()
			t.Errorf("Set %s Want %d, got %d\n", dynSets["set05"].name, dynSets["set05"].slt, got)
		}
	})
	t.Run(dynSets["set20"].name, func(t *testing.T) {
		if got := Dynamic(dynSets["set20"].its, dynSets["set20"].kp); got != dynSets["set20"].slt {
			t.Parallel()
			t.Errorf("Set %s Want %d, got %d\n", dynSets["set20"].name, dynSets["set20"].slt, got)
		}
	})
	t.Run(dynSets["set40"].name, func(t *testing.T) {
		if got := Dynamic(dynSets["set40"].its, dynSets["set40"].kp); got != dynSets["set40"].slt {
			t.Parallel()
			t.Errorf("Set %s Want %d, got %d\n", dynSets["set40"].name, dynSets["set40"].slt, got)
		}
	})
}

func TestDynamicCompress(t *testing.T) {
	t.Parallel()
	t.Run(dynSets["set05"].name, func(t *testing.T) {
		if got := DynamicCompress(dynSets["set05"].its, dynSets["set05"].kp); got != dynSets["set05"].slt {
			t.Parallel()
			t.Errorf("Set %s Want %d, got %d\n", dynSets["set05"].name, dynSets["set05"].slt, got)
		}
	})
	t.Run(dynSets["set20"].name, func(t *testing.T) {
		if got := DynamicCompress(dynSets["set20"].its, dynSets["set20"].kp); got != dynSets["set20"].slt {
			t.Parallel()
			t.Errorf("Set %s Want %d, got %d\n", dynSets["set20"].name, dynSets["set20"].slt, got)
		}
	})
	t.Run(dynSets["set40"].name, func(t *testing.T) {
		if got := DynamicCompress(dynSets["set40"].its, dynSets["set40"].kp); got != dynSets["set40"].slt {
			t.Parallel()
			t.Errorf("Set %s Want %d, got %d\n", dynSets["set40"].name, dynSets["set40"].slt, got)
		}
	})
	t.Run(dynSets["set100"].name, func(t *testing.T) {
		if got := DynamicCompress(dynSets["set100"].its, dynSets["set100"].kp); got != dynSets["set100"].slt {
			t.Parallel()
			t.Errorf("Set %s Want %d, got %d\n", dynSets["set100"].name, dynSets["set100"].slt, got)
		}
	})
}
