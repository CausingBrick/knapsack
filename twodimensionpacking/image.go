package twodimensionpacking

import (
	"fmt"
	"log"
	"math/rand"
	"sync"

	"github.com/fogleman/gg"
)

func DrawBox(box *Box, imgPath string, scale int) error {
	dc := gg.NewContext(scale*box.Width, scale*box.Height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	for _, rec := range box.StoredItem {
		dc.DrawRectangle(float64(scale*rec.Sp.X), float64(scale*rec.Sp.Y), float64(scale*rec.Width), float64(scale*rec.Height))
		dc.SetRGB255(rand.Intn(256), rand.Intn(256), rand.Intn(256))
		// dc.Stroke()
		dc.Fill()
	}
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	dc.SetRGB(0, 0, 0)
	dc.Stroke()
	dc.SavePNG(imgPath + ".png")
	return nil
}

func DrawBoxes(boxes []*Box, imgPath string, scale int) {
	var wg sync.WaitGroup
	wg.Add(len(boxes))
	for i := range boxes {
		go func(i int) {
			err := DrawBox(boxes[i], imgPath+fmt.Sprint(i+1), scale)
			if err != nil {
				log.Println(err)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
