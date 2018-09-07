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
	quit = input.Digital("Quit")
	play = input.Digital("Play")
)

var bindings = input.Bindings{
	"Default": {
		"Play": {"Space", "Mouse Left", "Button A"},
		"Quit": {"Escape", "Button Back"},
	},
}

type loop2 struct {
	logo       pixel.PictureID
	monochrome pixel.PaletteID
	colorful   pixel.PaletteID

	playing bool
}

// Initialization //////////////////////////////////////////////////////////////

func Example_interactive() {
	defer cozely.Recover()

	l := loop2{}
	l.setup()

	input.Load(bindings)
	cozely.Configure(cozely.UpdateStep(1.0 / 3))
	err := cozely.Run(&l)
	if err != nil {
		panic(err)
	}
	// Output:
}

func (l *loop2) setup() {
	pixel.SetResolution(160, 100)
	l.logo = pixel.Picture("graphics/cozely")
	l.monochrome = pixel.Palette("graphics/cozely")
	l.colorful = pixel.Palette("")
}

func (l *loop2) Enter() {
	l.monochrome.Use()
	l.shufflecolors()
}

func (loop2) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (l *loop2) React() {
	if play.Started(input.Any) {
		l.playing = !l.playing
	}
	if quit.Started(input.Any) {
		cozely.Stop(nil)
	}
}

func (l *loop2) Update() {
	if l.playing {
		l.shufflecolors()
	}
}

func (l *loop2) shufflecolors() {
	for i := 2; i < 14; i++ {
		g := 0.2 + 0.7*rand.Float32()
		r := 0.2 + 0.7*rand.Float32()
		b := 0.2 + 0.7*rand.Float32()
		l.colorful.Set(pixel.Color(i), color.SRGB{r, g, b})
	}
}

func (l *loop2) Render() {
	pixel.Clear(0)

	if l.playing {
		l.colorful.Use()
	} else {
		l.monochrome.Use()
	}

	o := pixel.Resolution().Minus(l.logo.Size()).Slash(2)
	pixel.Paint(l.logo, o)
}
