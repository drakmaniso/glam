// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

import (
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

// Identity matrix.
func Identity() Matrix {
	return Matrix{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}

////////////////////////////////////////////////////////////////////////////////

// Apply a tranformation matrix to a vector (i.e. returns matrix multiplied by
// column vector).
func Apply(m Matrix, v coord.XYZ) coord.XYZ {
	return coord.XYZ{
		X: m[0][0]*v.X + m[1][0]*v.Y + m[2][0]*v.Z,
		Y: m[0][1]*v.X + m[1][1]*v.Y + m[2][1]*v.Z,
		Z: m[0][2]*v.X + m[1][2]*v.Y + m[2][2]*v.Z,
	}
}

////////////////////////////////////////////////////////////////////////////////

// Translation by a vector.
func Translation(t coord.XY) Matrix {

	return Matrix{
		{1, 0, 0},
		{0, 1, 0},
		{t.X, t.Y, 1},
	}
}

////////////////////////////////////////////////////////////////////////////////

// Rotation by an angle.
func Rotation(angle float32) Matrix {
	c := math32.Cos(angle)
	s := math32.Sin(angle)

	return Matrix{
		{c, -s, 0},
		{s, c, 0},
		{0, 0, 1},
	}
}

////////////////////////////////////////////////////////////////////////////////

// RotationAround a point.
func RotationAround(angle float32, center coord.XY) Matrix {
	c := math32.Cos(angle)
	s := math32.Sin(angle)

	x, y := center.X, center.Y

	return Matrix{
		{c, -s, 0},
		{s, c, 0},
		{x - c*x - s*y, y + s*x - c*y, 1},
	}
}

////////////////////////////////////////////////////////////////////////////////

// Scaling and/or mirror along both axis.
func Scaling(s coord.XY) Matrix {
	return Matrix{
		{s.X, 0, 0},
		{0, s.Y, 0},
		{0, 0, 1},
	}
}

////////////////////////////////////////////////////////////////////////////////

// ScalingAround a point (and/or mirror).
func ScalingAround(s coord.XY, center coord.XY) Matrix {
	sx, sy := s.X, s.Y
	cx, cy := center.X, center.Y

	return Matrix{
		{sx, 0, 0},
		{0, sy, 0},
		{cx - cx*sx, cy - cy*sy, 1},
	}
}

////////////////////////////////////////////////////////////////////////////////

// Shearing along both axis.
func Shearing(s coord.XY) Matrix {
	return Matrix{
		{1, s.Y, 0},
		{s.X, 1, 0},
		{0, 0, 1},
	}
}

////////////////////////////////////////////////////////////////////////////////

// Viewport returns a transformation matrix that scale to an aspect ratio and
// zoom.
func Viewport(zoom, aspectRatio float32) Matrix {
	height := zoom / 2
	width := height * aspectRatio
	return Scaling(coord.XY{X: width, Y: height})
}

////////////////////////////////////////////////////////////////////////////////
