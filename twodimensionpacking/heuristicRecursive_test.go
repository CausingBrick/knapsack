package twodimensionpacking

import (
	"fmt"
	"testing"
)

var hrsets = []*dataset{
	datasetNew("M1a", "./dataset/M1a.txt"),
	datasetNew("M3c", "./dataset/M3c.txt"),
}

func TestHeuristic(t *testing.T) {
	t.Parallel()
	t.Run(hrsets[0].name, func(t *testing.T) {
		t.Parallel()
		var itsArea, boxesArea float64
		boxes := Heuristic(20, 15, hrsets[0].its)
		for _, box := range boxes {
			boxesArea += float64(box.Area())
			for _, item := range box.StoredItem {
				itsArea += float64(item.Area())
			}
		}
		DrawBoxes(boxes, "img/HR_M1a", 10)
		fmt.Printf("Set: %s\nBoxes nums: %d, Area usage: %f %% .\n", hrsets[0].name, len(boxes), itsArea/boxesArea*100)
		fmt.Printf("Area box:%f, items:%f.\n", boxesArea, itsArea)
	})
	t.Run(hrsets[1].name, func(t *testing.T) {
		t.Parallel()
		var itsArea, boxesArea float64
		boxes := Heuristic(60, 100, hrsets[1].its)
		for _, box := range boxes {
			boxesArea += float64(box.Area())
			for _, item := range box.StoredItem {
				itsArea += float64(item.Area())
			}
		}
		DrawBoxes(boxes, "img/HR_M3c", 2)
		fmt.Printf("Set: %s\nBoxes nums: %d, Area usage: %f %% .\n", hrsets[1].name, len(boxes), itsArea/boxesArea*100)
		fmt.Printf("Area box:%f, items:%f.\n", boxesArea, itsArea)
	})
}
