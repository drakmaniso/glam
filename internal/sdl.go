// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

// #include "sdl.h"
import "C"

import "errors"

//------------------------------------------------------------------------------

// GetSDLError returns nil or the current SDL Error wrapped in a Go error.
func GetSDLError() error {
	if s := C.SDL_GetError(); s != nil {
		return errors.New(C.GoString(s))
	}
	return nil
}

//------------------------------------------------------------------------------