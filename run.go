// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

import (
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

// GameLoop methods are called during the main loop to process events, Update
// the game state and render it.
type GameLoop interface {
	// The loop
	Update() error
	Draw() error

	// Window events
	WindowShown()
	WindowHidden()
	WindowResized(width, height int32)
	WindowMinimized()
	WindowMaximized()
	WindowRestored()
	WindowMouseEnter()
	WindowMouseLeave()
	WindowFocusGained()
	WindowFocusLost()
	WindowQuit()

	// Keyboard events
	KeyDown(l key.Label, p key.Position)
	KeyUp(l key.Label, p key.Position)

	// Mouse events
	MouseMotion(deltaX, deltaY int32, posX, posY int32)
	MouseButtonDown(b mouse.Button, clicks int)
	MouseButtonUp(b mouse.Button, clicks int)
	MouseWheel(deltaX, deltaY int32)

	// Pixel events
	ScreenResized(width, height int16, pixel int32)
}

// Handlers implements default handlers for all events.
//
// It's an empty struct intended to be embedded in the user-defined GameLoop:
//
//  type loop struct {
//    glam.Handlers
//  }
//
// This way it's possible to implement the GameLoop interface without writing a
// method for each event.
type Handlers = internal.Handlers

//------------------------------------------------------------------------------

func SetTimeStep(t float64) {
	internal.TimeStep = t
}

func TimeStep() float64 {
	return internal.TimeStep
}

//------------------------------------------------------------------------------

// Run initializes the framework, calls setup and then starts the game loop.
//
// The Update callback is called with a fixed time step, while the Draw callback
// is tied to the framerate. Event callbacks are called before each Update, but
// at least once for every frame. The loop runs until Stop() is called.
//
// Important: must be called from main.main, or at least from a function that is
// known to run on the main OS thread.
func Run(setup func() error, loop GameLoop) error {
	defer internal.SDLQuit()
	defer internal.DestroyWindow()

	internal.Loop = loop

	// Setup

	err := internal.Setup()
	if err != nil {
		return internal.Error("in internal setup", err)
	}

	err = gl.Setup(internal.Config.Debug)
	if err != nil {
		return internal.Error("in OpenGL setup", err)
	}

	err = internal.PaletteSetup()
	if err != nil {
		return internal.Error("in palette setup", err)
	}

	err = internal.PixelSetup()
	if err != nil {
		return internal.Error("in pixel Setup", err)
	}

	if setup != nil {
		err = setup()
		if err != nil {
			return internal.Error("in game loop Setup", err)
		}
	}

	// First, send a fake resize window event
	internal.Loop.WindowResized(internal.Window.Width, internal.Window.Height)
	internal.ResizeScreen()

	// Main Loop

	internal.FrameTime = 0.0
	internal.UpdateLag = 0.0

	then := internal.GetSeconds()
	now := then
	gametime := 0.0

	for !internal.QuitRequested {
		internal.FrameTime = now - then
		countFrames()
		if internal.FrameTime > 4*internal.TimeStep {
			// Prevent "spiral of death" when Draw can't keep up with Update
			internal.FrameTime = 4 * internal.TimeStep
		}

		// Update and Events

		internal.UpdateLag += internal.FrameTime
		if internal.UpdateLag < internal.TimeStep {
			// Process events even if there is no Update this frame
			internal.ProcessEvents()
		}
		for internal.UpdateLag >= internal.TimeStep {
			// Do the Time Step
			internal.UpdateLag -= internal.TimeStep
			gametime += internal.TimeStep
			internal.GameTime = gametime
			// Events
			internal.ProcessEvents()
			// Update
			err = internal.Loop.Update()
			if err != nil {
				return internal.Error("in Update callback", err)
			}
		}

		// Draw

		internal.GameTime = gametime + internal.UpdateLag
		err = internal.Loop.Draw()
		if err != nil {
			return internal.Error("in Draw callback", err)
		}

		err = internal.PixelDraw()
		if err != nil {
			return internal.Error("in pixel Draw", err)
		}

		internal.SwapWindow()

		then = now
		now = internal.GetSeconds()
	}
	return nil
}

//------------------------------------------------------------------------------

// GameTime returns the time elapsed in the game. It is updated before each call
// to Update and before each call to Draw.
func GameTime() float64 {
	return internal.GameTime
}

// FrameTime returns the time elapsed between the previous frame and the one
// being drawn. See also UpdateLag.
func FrameTime() float64 {
	return internal.FrameTime
}

// UpdateLag returns the time elapsed between the last update and the frame
// being drawn. It should be used during Draw to extrapolate (or interpolate)
// the game state.
//
// Note: if called during Update (or an event callback), it returns the time
// between the current update and the next Draw call.
func UpdateLag() float64 {
	return internal.UpdateLag
}

//------------------------------------------------------------------------------

func countFrames() {
	frCount++
	frSum += internal.FrameTime
	if internal.FrameTime > xrunThreshold {
		xrunCount++
	}
	if frSum >= frInterval {
		frAverage = frSum / float64(frCount)
		xrunPrevious = xrunCount
		//TODO: microtext.PrintFrameTime(frAverage, xrunCount)
		frSum = 0
		frCount = 0
		xrunCount = 0
	}
}

// FrameStats returns the average durations of frames; it is updated 4
// times per second. It also returns the number of overruns (i.e. frame time
// longer than the threshold) during the last measurment interval.
func FrameStats() (t float64, overruns int) {
	return frAverage, xrunPrevious
}

const frInterval = 1.0 / 4.0

var frAverage float64
var frSum float64
var frCount int

const xrunThreshold float64 = 17 / 1000.0

var xrunCount, xrunPrevious int

//------------------------------------------------------------------------------

// Stop request the game loop to stop.
func Stop() {
	internal.QuitRequested = true
}

//------------------------------------------------------------------------------

// Path returns the (slash-separated) path of the executable, with a trailing
// slash.
func Path() string {
	return internal.Path
}

//------------------------------------------------------------------------------
