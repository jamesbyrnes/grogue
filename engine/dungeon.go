package engine

import (
	"math/rand"
	"time"
)

type Dungeon struct {
	Levels []Board
}

func (d *Dungeon) Initialize(numLevels int) {
	for i := 0; i < numLevels; i++ {
		var b Board
		b.Initialize()
		d.Levels = append(d.Levels, b)
	}
}

func (d *Dungeon) Build(level int, numRooms int) {
	for i := 0; i < numRooms; i++ {
		x := rand.Intn(98)
		y := rand.Intn(98)
		h := rand.Intn(16)
		w := rand.Intn(16)

		if x+w > 99 {
			w = 99 - x
		} else if y+h > 99 {
			h = 99 - y
		}

		r := Rectangle{
			tl:     [2]int{y, x},
			width:  w,
			height: h,
		}

		r.Draw(&d.Levels[level])
	}
}

// func (d *Dungeon) build_rooms(level int, numRooms int, x int, y int, w int, h int) {
// 	if numRooms > 0 {
// 		r := Rectangle {
// 			tl: 	[2]int{y, x},
// 			width: 	w,
// 			height: h,
// 		}
// 		r.Draw[d.levels[level]]

// 	}
// }

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
