// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

/*
#include "glad.h"

static void Viewport(GLint x,  GLint y,  GLsizei width,  GLsizei height) {
	glViewport(x, y, width, height);
}
*/
import "C"

import (
	"github.com/drakmaniso/glam/geom"
)

//------------------------------------------------------------------------------

func Viewport(orig, size geom.IVec2) {
	C.Viewport(C.GLint(orig.X), C.GLint(orig.Y), C.GLsizei(size.X), C.GLsizei(size.Y))
}

//------------------------------------------------------------------------------
