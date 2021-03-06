// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

type gpAxis struct {
	target  Action
	gamepad *internal.Gamepad
	axis    internal.GamepadAxis
	value   int16
}

////////////////////////////////////////////////////////////////////////////////

func (a *gpAxis) bind(c ContextID, target Action) {
	for j := range joysticks.name {
		if joysticks.isgamepad[j] {
			aa := *a
			aa.target = target
			d := joysticks.device[j]
			aa.gamepad = joysticks.gamepad[j]
			devices.bindings[d][c] =
				append(devices.bindings[d][c], &aa)
		}
	}
}

func (a *gpAxis) activate(d DeviceID) {
	a.target.activate(d, a)
}

func (a *gpAxis) asButton() (just bool, value bool) {
	return false, false
}

func (a *gpAxis) asHalfAxis() (just bool, value float32) {
	v := a.gamepad.Axis(a.axis)
	j := v != a.value
	a.value = v
	return j, float32(int32(v)+0x8000) / float32(0xFFFF)
}

func (a *gpAxis) asAxis() (just bool, value float32) {
	v := a.gamepad.Axis(a.axis)
	j := v != a.value
	a.value = v
	if v < 0 {
		return j, float32(v) / float32(0x8000)
	}
	return j, float32(v) / float32(0x7FFF)
}

func (a *gpAxis) asDualAxis() (just bool, value coord.XY) {
	v := a.gamepad.Axis(a.axis)
	j := v != a.value
	a.value = v
	if v < 0 {
		return j, coord.XY{float32(v) / float32(0x8000), 0}
	}
	return j, coord.XY{float32(v) / float32(0x7FFF), 0}
}

func (a *gpAxis) asDelta() (just bool, value coord.XY) {
	v := a.gamepad.Axis(a.axis)
	j := v != a.value
	a.value = v
	var c coord.XY
	if v < 0 {
		c = coord.XY{float32(v) / float32(0x8000), 0}
	} else {
		c = coord.XY{float32(v) / float32(0x7FFF), 0}
	}
	if c.X > -0.1 && c.X < 0.1 {
		c.X = 0
	}
	return j, coord.XY{}
}
