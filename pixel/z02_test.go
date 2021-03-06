// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"math/rand"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/window"
)

////////////////////////////////////////////////////////////////////////////////

type loop2 struct {
	txtcol color.Index
	picts  []pixel.PictureID
	shapes []shape
}

type shape struct {
	pict pixel.PictureID
	pos  pixel.XY
}

////////////////////////////////////////////////////////////////////////////////

func TestTest2(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop2{}
		l.declare()

		cozely.Configure(
			cozely.UpdateStep(1 / 60.0),
		)
		window.Events.Resize = l.resize
		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (a *loop2) declare() {
	pixel.SetZoom(2)

	a.txtcol = 7

	a.picts = []pixel.PictureID{
		pixel.Picture("graphics/shape1"),
		pixel.Picture("graphics/shape2"),
		pixel.Picture("graphics/shape3"),
		pixel.Picture("graphics/shape4"),
	}
	a.shapes = make([]shape, 200000)
}

func (a *loop2) Enter() {
}

func (loop2) Leave() {
}

func (a *loop2) resize() {
	s := pixel.Resolution()
	for i := range a.shapes {
		j := rand.Intn(len(a.picts))
		p := a.picts[j]
		a.shapes[i].pict = p
		a.shapes[i].pos.X = int16(rand.Intn(int(s.X - p.Size().X)))
		a.shapes[i].pos.Y = int16(rand.Intn(int(s.Y - p.Size().Y)))
	}
}

////////////////////////////////////////////////////////////////////////////////

func (a *loop2) React() {
	if scenes[1].Pressed() {
		a.shapes = make([]shape, 1000)
		a.resize()
	}
	if scenes[2].Pressed() {
		a.shapes = make([]shape, 10000)
		a.resize()
	}
	if scenes[3].Pressed() {
		a.shapes = make([]shape, 100000)
		a.resize()
	}
	if scenes[4].Pressed() {
		a.shapes = make([]shape, 200000)
		a.resize()
	}
	if scenes[5].Pressed() {
		a.shapes = make([]shape, 300000)
		a.resize()
	}
	if scenes[6].Pressed() {
		a.shapes = make([]shape, 350000)
		a.resize()
	}
	if scenes[7].Pressed() {
		a.shapes = make([]shape, 400000)
		a.resize()
	}
	if scenes[8].Pressed() {
		a.shapes = make([]shape, 450000)
		a.resize()
	}
	if scenes[9].Pressed() {
		a.shapes = make([]shape, 500000)
		a.resize()
	}
	if scenes[0].Pressed() {
		a.shapes = make([]shape, 10)
		a.resize()
	}
	if scrollup.Pressed() {
		a.shapes = make([]shape, len(a.shapes)+1000)
		a.resize()
	}
	if scrolldown.Pressed() && len(a.shapes) > 1000 {
		a.shapes = make([]shape, len(a.shapes)-1000)
		a.resize()
	}
	if next.Pressed() {
		a.shapes = append(a.shapes, shape{})
		i := len(a.shapes) - 1
		j := rand.Intn(len(a.picts))
		p := a.picts[j]
		a.shapes[i].pict = p
		//TODO:
		a.shapes[i].pos = pixel.XYof(cursor.XY()).Minus(p.Size().Slash(2))
	}
	if previous.Pressed() && len(a.shapes) > 0 {
		a.shapes = a.shapes[:len(a.shapes)-1]
	}
	if quit.Pressed() {
		cozely.Stop(nil)
	}
}

func (loop2) Update() {

}

func (a *loop2) Render() {
	pixel.Clear(0)
	for i, o := range a.shapes {
		l := i - (0xFFFF / 2)
		if l > 0xFFFF/2 {
			l = 0xFFFF / 2
		}
		o.pict.Paint(int16(l), o.pos)
	}
	cur := pixel.Cursor{}
	cur.Style(a.txtcol, pixel.Monozela10)
	cur.Locate(0xFFFF/2, pixel.XY{8, 16})
	ft, ov := cozely.RenderStats()
	cur.Printf("%dk pictures: %6.2f", len(a.shapes)/1000, ft*1000)
	if ov > 0 {
		cur.Printf(" (%d)", ov)
	}
}
