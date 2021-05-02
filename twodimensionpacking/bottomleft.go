package twodimensionpacking

import (
	"sort"
)

// BottomLeft bottom left first algoritms
func BottomLeft(boxH, boxW int, items []*Item) []*Box {
	sortArea := items
	sort.Slice(sortArea, func(i int, j int) bool {
		return items[i].Area() > items[j].Area()
	})

	var boxes []*Box
	for len(sortArea) != 0 {
		box := BoxNew(boxH, boxW)
		for i := 0; i < len(sortArea); i++ {
			item := sortArea[i]
			// set starting point of item
			sp := PointNew(boxW-item.Width, boxH-item.Height)
			if isOverlap(box, item, sp) {
				continue
			}
			movedPoint := getMovedPoint(box, item, sp)
			if sp == movedPoint {
				continue
			}
			// intsert item to box
			box.StoredItem = append(box.StoredItem, RectangleNew(&item.Size, movedPoint))
			// delete item form sortHeight
			sortArea = append(sortArea[:i], sortArea[i+1:]...)
			i--
		}
		boxes = append(boxes, box)
	}
	return boxes
}

// maxDropHeight returns the maximum drop height of the item in the box.
func maxDropHeight(box *Box, item *Item, sp *Point) (height int) {
	if len(box.StoredItem) == 0 {
		return sp.Y
	}

	dists := []int{}
	itemRec := RectangleNew(&item.Size, sp)
	for _, stored := range box.StoredItem {
		isInter, dist := itemRec.BottomLine().VerticalIntersect(stored.TopLine())
		if isInter {
			dists = append(dists, dist)
		}
	}
	if len(dists) == 0 {
		height = sp.Y
	} else {
		height = min(dists...)
	}
	return
}

// maxShiftWeight returns the maximum shift left with of the item which will
//  be placed in the box.
func maxShiftWeight(box *Box, item *Item, sp *Point) (height int) {
	if len(box.StoredItem) == 0 {
		return sp.X
	}

	dists := []int{}
	itemRec := RectangleNew(&item.Size, sp)
	for _, stored := range box.StoredItem {
		isInter, dist := itemRec.LeftLine().HorizontalIntersect(stored.RightLine())
		if isInter {
			dists = append(dists, dist)
		}
	}
	if len(dists) == 0 {
		height = sp.X
	} else {
		height = min(dists...)
	}
	return
}

// getMovedPoint returns the point of item after moving with BL strategy.
func getMovedPoint(box *Box, item *Item, sp *Point) *Point {
	final := *sp
	for {
		maxh := maxDropHeight(box, item, &final)
		if maxh > 0 {
			final.Y -= maxh
		}
		maxw := maxShiftWeight(box, item, &final)
		if maxw > 0 {
			final.X -= maxw
		}
		if maxh <= 0 && maxw <= 0 {
			break
		}
	}
	return &final
}

// overlap checks if item overlap with items in box.
func isOverlap(box *Box, item *Item, sp *Point) bool {
	for _, inbox := range box.StoredItem {
		if RectangleNew(&item.Size, sp).Intersect(inbox) {
			return true
		}
	}
	return false
}

func min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}
