package engine

import (
	"math/rand"
	"time"
)

type DrawableRoom interface {
	Draw()
	GetRandWallTile()
}

type Rectangle struct {
	tl     [2]int
	width  int
	height int
}

func (r *Rectangle) Draw(b *Board) {
	for i := 0; i < r.height; i++ {
		for j := 0; j < r.width; j++ {
			if i == 0 || i == r.height-1 || j == 0 || j == r.width-1 {
				b.Tiles[i+r.tl[0]][j+r.tl[1]].kind = "wall"
			} else {
				b.Tiles[i+r.tl[0]][j+r.tl[1]].kind = "earth"
			}
		}
	}
}

func (r *Rectangle) GetRandWallTile() [2]int {
	var edgeNS bool
	var randX int
	var randY int

	edgeNS = rand.Intn(2) == 0

	if edgeNS {
		randX = rand.Intn(r.width-2) + (r.tl[1] + 1)
		randY = rand.Intn(2)*r.height + r.tl[0]
	} else {
		randX = rand.Intn(2)*r.width + r.tl[1]
		randY = rand.Intn(r.height-2) + (r.tl[0] + 1)
	}

	return [2]int{randY, randX}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
