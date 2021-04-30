package twodimensionpacking

//BottomLeft bottom left first algoritms
// func BottomLeft(Weight *, items *Items) (int, float32) {

// 	return 0, 0
// }

// MAXDropHeight returns the maximum drop height of the item in the box.
func MAXDropHeight(box *Box, item *Item, sp *Point) (height int) {
	if len(box.StoredItem) == 0 {
		return sp.Y
	}

	dists := []int{}
	itemBottomLine := RectangleNew(&item.Size, sp).BottomLine()
	for _, stored := range box.StoredItem {
		isInter, dist := itemBottomLine.VerticalIntersect(stored.TopLine())
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

// MAXShiftWeight returns the maximum shift left with of the item which will
//  be placed in the box.
func MAXShiftWeight(box *Box, item *Item, sp *Point) (height int) {
	if len(box.StoredItem) == 0 {
		return sp.X
	}

	dists := []int{}
	itemLeftLine := RectangleNew(&item.Size, sp).LeftLine()
	for _, stored := range box.StoredItem {
		isInter, dist := itemLeftLine.HorizontalIntersect(stored.RightLine())
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

func BLInsertItem(box *Box, item *Item, sp *Point) *Point {
	final := *sp
	for {
		maxh := MAXDropHeight(box, item, sp)
		final.Y -= maxh
		maxw := MAXShiftWeight(box, item, sp)
		final.X -= maxw
		if maxh == 0 && maxw == 0 {
			return &final
		}
	}
}

// func Overlap(box *Box, item *Item, sp *Point) bool {

// }

func max(nums ...int) int {
	m := nums[0]
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}

func min(nums ...int) int {
	m := nums[0]
	for _, num := range nums {
		if num < m {
			m = num
		}
	}
	return m
}
