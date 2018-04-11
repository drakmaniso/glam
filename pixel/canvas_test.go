// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"math/rand"
	"testing"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/input"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
)

//------------------------------------------------------------------------------

var cnvContext = input.Context("TestCanvas", quit)

var cnvBindings = input.Bindings{
	"TestCanvas": {
		"Quit": {"Escape"},
	},
}

//------------------------------------------------------------------------------

var cnvScreen = pixel.Canvas(pixel.Zoom(2))

var shapePictures = []pixel.PictureID{
	pixel.Picture("graphics/shape1"),
	pixel.Picture("graphics/shape2"),
	pixel.Picture("graphics/shape3"),
	pixel.Picture("graphics/shape4"),
}

type shape struct {
	pict  pixel.PictureID
	pos   plane.Pixel
	depth int16
}

var shapes [2048]shape

//------------------------------------------------------------------------------

func TestCanvas_depth(t *testing.T) {
	do(func() {
		glam.Configure(
			glam.UpdateStep(1 / 60.0),
		)
		glam.Events.Resize = resize
		err := glam.Run(cnvLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

//------------------------------------------------------------------------------

type cnvLoop struct{}

//------------------------------------------------------------------------------

func (cnvLoop) Enter() error {
	input.Load(testBindings)
	testContext.Activate(1)
	palette.Load("graphics/shape1")
	return nil
}

func (cnvLoop) Leave() error { return nil }

//------------------------------------------------------------------------------

func (cnvLoop) React() error {
	if quit.JustPressed(1) {
		glam.Stop()
	}
	return nil
}

//------------------------------------------------------------------------------

func (cnvLoop) Update() error { return nil }

//------------------------------------------------------------------------------

func (cnvLoop) Render() error {
	cnvScreen.Clear(0)
	for i, o := range shapes {
		if float64(i)/32 > glam.GameTime() {
			break
		}
		cnvScreen.Picture(o.pict, o.depth, o.pos)
	}
	cnvScreen.Display()
	return nil
}

//------------------------------------------------------------------------------

func resize() {
	s := cnvScreen.Size()
	for i := range shapes {
		j := rand.Intn(len(shapePictures))
		shapes[i].depth = int16(j)
		p := shapePictures[j]
		shapes[i].pict = p
		shapes[i].pos.X = int16(rand.Intn(int(s.X - p.Size().X)))
		shapes[i].pos.Y = int16(rand.Intn(int(s.Y - p.Size().Y)))
	}
}

//------------------------------------------------------------------------------
