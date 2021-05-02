package twodimensionpacking

import (
	"fmt"
	"log"
	"testing"
)

var sets = []*dataset{
	datasetNew("M1a", "./dataset/M1a.txt"),
	datasetNew("M3c", "./dataset/M3c.txt"),
}

func TestBottomLeft(t *testing.T) {
	t.Parallel()
	t.Run(sets[0].name, func(t *testing.T) {
		t.Parallel()
		var itsArea, boxesArea float64
		boxes := BottomLeft(20, 15, sets[0].its)
		for _, box := range boxes {
			boxesArea += float64(box.Area())
			for _, item := range box.StoredItem {
				itsArea += float64(item.Area())
			}
		}
		DrawBoxes(boxes, "./dataset/M1a/", 10)
		fmt.Printf("Set: %s\nBoxes nums: %d, Area usage: %f %% .\n", sets[0].name, len(boxes), itsArea/boxesArea*100)
		fmt.Printf("Area box:%f, items:%f.\n", boxesArea, itsArea)
	})
	t.Run(sets[1].name, func(t *testing.T) {
		t.Parallel()
		var itsArea, boxesArea float64
		boxes := BottomLeft(60, 100, sets[1].its)
		for _, box := range boxes {
			boxesArea += float64(box.Area())
			for _, item := range box.StoredItem {
				itsArea += float64(item.Area())
			}
		}
		DrawBoxes(boxes, "./dataset/M3c/", 10)
		fmt.Printf("Set: %s\nBoxes nums: %d, Area usage: %f %% .\n", sets[1].name, len(boxes), itsArea/boxesArea*100)
		fmt.Printf("Area box:%f, items:%f.\n", boxesArea, itsArea)
	})
}

func TestReadToLines(t *testing.T) {
	nums, err := readToLines("./dataset/M3c.txt")
	if err != nil {
		log.Println(err, nums)
	}
	for _, num := range nums {
		fmt.Println(num)
	}
}
