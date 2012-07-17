/*
Copyright (c) 2012 krepa098 (krepa098 at gmail dot com)
This software is provided 'as-is', without any express or implied warranty.
In no event will the authors be held liable for any damages arising from the use of this software.
Permission is granted to anyone to use this software for any purpose, including commercial applications, 
and to alter it and redistribute it freely, subject to the following restrictions:
	1.	The origin of this software must not be misrepresented; you must not claim that you wrote the original software. 
		If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
	2. 	Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
	3. 	This notice may not be removed or altered from any source distribution.
*/

package GoSFML2

// #include <SFML/System.h>
import "C"
import "math"

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

func (this Vector2i) Plus(other Vector2i) Vector2i {
	return Vector2i{X: this.X + other.X, Y: this.Y + other.Y}
}

func (this Vector2i) Minus(other Vector2i) Vector2i {
	return Vector2i{X: this.X - other.X, Y: this.Y - other.Y}
}

func (this Vector2i) Length() float32 {
	return float32(math.Sqrt(float64(this.X*this.X + this.Y*this.Y)))
}

/////////////////////////////////////
// Vector2u

func (this Vector2u) Plus(other Vector2u) Vector2u {
	return Vector2u{X: this.X + other.X, Y: this.Y + other.Y}
}

func (this Vector2u) Minus(other Vector2u) Vector2u {
	return Vector2u{X: this.X - other.X, Y: this.Y - other.Y}
}

func (this Vector2u) Length() float32 {
	return float32(math.Sqrt(float64(this.X*this.X + this.Y*this.Y)))
}

/////////////////////////////////////
// Vector2f

func (this Vector2f) Plus(other Vector2f) Vector2f {
	return Vector2f{X: this.X + other.X, Y: this.Y + other.Y}
}

func (this Vector2f) Minus(other Vector2f) Vector2f {
	return Vector2f{X: this.X - other.X, Y: this.Y - other.Y}
}

func (this Vector2f) Length() float32 {
	return float32(math.Sqrt(float64(this.X*this.X + this.Y*this.Y)))
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
