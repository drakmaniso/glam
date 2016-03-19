// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

import (
	"fmt"
	"log"
	"unsafe"
)

// #include "sdl.h"
import "C"

//------------------------------------------------------------------------------

// Window is the game window.
var Window struct {
	window  *C.SDL_Window
	context C.SDL_GLContext
	Width   int
	Height  int
}

// Focus state
var (
	HasFocus      bool
	HasMouseFocus bool
)

//------------------------------------------------------------------------------

// OpenWindow creates the game window and its associated OpenGL context.
func OpenWindow(
	title string,
	resolution [2]int,
	display int,
	fullscreen bool,
	fullscreenMode string,
	vsync bool,
) error {
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MAJOR_VERSION, 4)
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MINOR_VERSION, 5)
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_PROFILE_MASK,
		C.SDL_GL_CONTEXT_PROFILE_CORE)
	C.SDL_GL_SetAttribute(C.SDL_GL_DOUBLEBUFFER, 1)
	C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLESAMPLES, 8)

	var si C.int
	if vsync {
		si = 1
	}
	C.SDL_GL_SetSwapInterval(si)

	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))

	Window.Width, Window.Height = resolution[0], resolution[1]

	var fs uint32
	if fullscreen {
		if fullscreenMode == "Desktop" {
			fs = C.SDL_WINDOW_FULLSCREEN_DESKTOP
		} else {
			fs = C.SDL_WINDOW_FULLSCREEN
		}
	}
	fl := C.SDL_WINDOW_OPENGL | C.SDL_WINDOW_RESIZABLE | C.Uint32(fs)

	Window.window = C.SDL_CreateWindow(
		t,
		C.int(C.SDL_WINDOWPOS_CENTERED_MASK|display),
		C.int(C.SDL_WINDOWPOS_CENTERED_MASK|display),
		C.int(Window.Width),
		C.int(Window.Height),
		fl,
	)
	if Window.window == nil {
		err := GetSDLError()
		log.Print(err)
		return err
	}

	ctx, err := C.SDL_GL_CreateContext(Window.window)
	if err != nil {
		log.Print(err)
		return err
	}
	Window.context = ctx

	logOpenGLInfos()

	//TODO: Send a fake resize event (for the renderer)

	return nil
}

// logOpenGLInfos displays information about the OpenGL context
func logOpenGLInfos() {
	s := "OpenGL: "
	maj, err1 := sdlGLAttribute(C.SDL_GL_CONTEXT_MAJOR_VERSION)
	min, err2 := sdlGLAttribute(C.SDL_GL_CONTEXT_MINOR_VERSION)
	if err1 == nil && err2 == nil {
		s += fmt.Sprintf("%d.%d", maj, min)
	}

	db, err1 := sdlGLAttribute(C.SDL_GL_DOUBLEBUFFER)
	if err1 == nil {
		if db != 0 {
			s += ", Double Buffer"
		} else {
			s += ", NO Double Buffer"
		}
	}

	av, err1 := sdlGLAttribute(C.SDL_GL_ACCELERATED_VISUAL)
	if err1 == nil {
		if av != 0 {
			s += ", Accelerated"
		} else {
			s += ", NOT Accelerated"
		}
	}

	sw := C.SDL_GL_GetSwapInterval()
	if sw > 0 {
		if sw != 0 {
			s += ", VSync"
		} else {
			s += ", NO VSync"
		}
	} else {
		err1 = GetSDLError()
		log.Print(err1)
	}
	log.Printf(s)
}

func sdlGLAttribute(attr C.SDL_GLattr) (int, error) {
	var v C.int
	errcode := C.SDL_GL_GetAttribute(attr, &v)
	if errcode < 0 {
		return 0, GetSDLError()
	}
	return int(v), nil
}

//------------------------------------------------------------------------------

// DestroyWindow closes the game window and delete the OpenGL context
func DestroyWindow() {
	C.SDL_GL_DeleteContext(Window.context)
	C.SDL_DestroyWindow(Window.window)
}

//------------------------------------------------------------------------------