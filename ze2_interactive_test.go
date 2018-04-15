package cozely_test

import (
	"math/rand"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

// Declarations ////////////////////////////////////////////////////////////////

var (
	quit    = input.Bool("Quit")
	start   = input.Bool("Start")
	context = input.Context("Default", start, quit)
)

var bindings = input.Bindings{
	"Default": {
		"Start": {"Space", "Mouse Left"},
		"Quit":  {"Escape"},
	},
}

var (
	canv = pixel.Canvas(pixel.Resolution(160, 100))
	logo = pixel.Picture("graphics/cozely")
	pal1 = color.PaletteFrom("graphics/cozely")
	pal2 = color.Palette()
)

var started = false

var reverse = false

type loop2 struct{}

// Initialization //////////////////////////////////////////////////////////////

func Example_interactive() {
	cozely.Configure(cozely.UpdateStep(1.0 / 3))
	cozely.Run(loop2{})
	// Output:
}

func (loop2) Enter() error {
	bindings.Load()
	context.Activate(1)
	pal1.Activate()
	shufflecolors()
	return nil
}

func (loop2) Leave() error {
	return nil
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (loop2) React() error {
	if start.JustPressed(1) {
		started = !started
	}
	if quit.JustPressed(1) {
		cozely.Stop()
	}
	return nil
}

func (loop2) Update() error {
	if started {
		shufflecolors()
	}
	return nil
}

func shufflecolors() {
	dark := [12]bool{
		true, false, true, false, true, false,
		false, true, false, true, false, true,
	}
	reverse = !reverse
	for i := 2; i < 14; i++ {
		r := .2 + .8*rand.Float32()
		g := .2 + .8*rand.Float32()
		b := .2 + .8*rand.Float32()
		if dark[i-2] != reverse {
			pal2.Set(uint8(i), color.SRGB{r, g, b})
		} else {
			pal2.Set(uint8(i), color.LRGB{r, g, b})
		}
	}
}

func (loop2) Render() error {
	canv.Clear(0)

	if started {
		pal2.Activate()
	} else {
		pal1.Activate()
	}

	o := canv.Size().Minus(logo.Size()).Slash(2)
	canv.Picture(logo, 0, o)

	canv.Display()
	return nil
}