package main

import (
	"github.com/nsf/termbox-go"
)

type Room struct {
	top    int
	bottom int
	left   int
	right  int
}

func (r *Room) Render() {
	var ch rune
	for i := r.left; i <= r.right; i++ {
		for j := r.top; j <= r.bottom; j++ {
			if i == r.left || i == r.right || j == r.top || j == r.bottom {
				ch = rune('x')
			} else {
				ch = rune('.')
			}
			termbox.SetCell(i, j, ch, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

type Player struct {
	x  int
	y  int
	ch rune
}

func (p *Player) Move(x int, y int) {
	w, h := termbox.Size()
	if x >= 0 && x < w {
		p.x = x
	}
	if y >= 0 && y < h {
		p.y = y
	}
}

func redraw_all(p Player, r Room) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	r.Render()

	termbox.SetCell(p.x, p.y, p.ch, termbox.ColorDefault, termbox.ColorDefault)

	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)

	pl1 := Player{x: 0, y: 0, ch: rune('@')}
	r1 := Room{top: 5, bottom: 10, left: 5, right: 20}

	redraw_all(pl1, r1)

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop

			case termbox.KeyArrowUp:
				pl1.Move(pl1.x, pl1.y-1)
			case termbox.KeyArrowDown:
				pl1.Move(pl1.x, pl1.y+1)
			case termbox.KeyArrowLeft:
				pl1.Move(pl1.x-1, pl1.y)
			case termbox.KeyArrowRight:
				pl1.Move(pl1.x+1, pl1.y)
			}

		case termbox.EventError:
			panic(ev.Err)

		case termbox.EventInterrupt:
			break mainloop
		}

		redraw_all(pl1, r1)
	}

	termbox.Close()
}
