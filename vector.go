// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/System.h>
import "C"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Vector2i struct {
	X, Y int
}

type Vector2u struct {
	X, Y uint
}

type Vector2f struct {
	X, Y float32
}

type Vector3f struct {
	X, Y, Z float32
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

/////////////////////////////////////
// Vector2i

// Returns the sum of two vectors.
func (this Vector2i) Plus(other Vector2i) Vector2i {
	return Vector2i{X: this.X + other.X, Y: this.Y + other.Y}
}

// Returns the difference of two vectors.
func (this Vector2i) Minus(other Vector2i) Vector2i {
	return Vector2i{X: this.X - other.X, Y: this.Y - other.Y}
}

/////////////////////////////////////
// Vector2u

// Returns the sum of two vectors.
func (this Vector2u) Plus(other Vector2u) Vector2u {
	return Vector2u{X: this.X + other.X, Y: this.Y + other.Y}
}

// Returns the difference of two vectors.
func (this Vector2u) Minus(other Vector2u) Vector2u {
	return Vector2u{X: this.X - other.X, Y: this.Y - other.Y}
}

/////////////////////////////////////
// Vector2f

// Returns the sum of two vectors.
func (this Vector2f) Plus(other Vector2f) Vector2f {
	return Vector2f{X: this.X + other.X, Y: this.Y + other.Y}
}

// Returns the difference of two vectors.
func (this Vector2f) Minus(other Vector2f) Vector2f {
	return Vector2f{X: this.X - other.X, Y: this.Y - other.Y}
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Vector2i) fromC(vec C.sfVector2i) {
	this.X = int(vec.x)
	this.Y = int(vec.y)
}

func (this *Vector2u) fromC(vec C.sfVector2u) {
	this.X = uint(vec.x)
	this.Y = uint(vec.y)
}

func (this *Vector2f) fromC(vec C.sfVector2f) {
	this.X = float32(vec.x)
	this.Y = float32(vec.y)
}

func (this *Vector3f) fromC(vec C.sfVector3f) {
	this.X = float32(vec.x)
	this.Y = float32(vec.y)
	this.Z = float32(vec.z)
}

func (this *Vector2i) toC() C.sfVector2i {
	return C.sfVector2i{x: C.int(this.X), y: C.int(this.Y)}
}

func (this *Vector2u) toC() C.sfVector2u {
	return C.sfVector2u{x: C.uint(this.X), y: C.uint(this.Y)}
}

func (this *Vector2f) toC() C.sfVector2f {
	return C.sfVector2f{x: C.float(this.X), y: C.float(this.Y)}
}

func (this *Vector3f) toC() C.sfVector3f {
	return C.sfVector3f{x: C.float(this.X), y: C.float(this.Y), z: C.float(this.Z)}
}
