package twodimensionpacking

import (
	"sort"
)

func Heuristic(boxH, boxW int, items []*Item) []*Box {
	sortArea := items
	sort.Slice(sortArea, func(i int, j int) bool {
		return items[i].Height > items[j].Height
	})
	var boxes []*Box
	for len(sortArea) != 0 {
		box := BoxNew(boxH, boxW)
		heuristic(box, &sortArea, *box.Rectangle)
		boxes = append(boxes, box)
	}
	return boxes
}

func heuristic(b *Box, items *[]*Item, r Rectangle) {
	for i := 0; i < len(*items); i++ {
		item := (*items)[i]
		if r.CanHold(&item.Size) {
			// move item form items to the boxes
			b.StoredItem = append(b.StoredItem, RectangleNew(&item.Size, r.Sp))
			(*items) = append((*items)[:i], (*items)[i+1:]...)
			heuristic(b, items, *RectangleNew(
				SizeNew(item.Height, r.Width-item.Width),
				PointNew(r.Sp.X+item.Width, r.Sp.Y),
			))
			heuristic(b, items, *RectangleNew(
				SizeNew(r.Height-item.Height, r.Width),
				PointNew(r.Sp.X, r.Sp.Y+item.Height),
			))
			break
		}
	}
}
