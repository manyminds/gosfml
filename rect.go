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

// #include <SFML/Graphics.h>
// int getSizeRecti() { return sizeof(sfIntRect); }
// int getSizeRectf() { return sizeof(sfFloatRect); }
import "C"
import "unsafe"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Rectf struct {
	Left   float32
	Top    float32
	Width  float32
	Height float32
}

type Recti struct {
	Left   int
	Top    int
	Width  int
	Height int
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func (this Rectf) Contains(x, y float32) bool {
	return sfBool2Go(C.sfFloatRect_contains(this.toCPtr(), C.float(x), C.float(y)))
}

func (this Recti) Contains(x, y int) bool {
	return C.sfIntRect_contains(this.toCPtr(), C.int(x), C.int(y)) == 1
}

func (this Rectf) Intersects(other *Rectf) (test bool, intersection *Rectf) {
	test = C.sfFloatRect_intersects(this.toCPtr(), other.toCPtr(), intersection.toCPtr()) == 1
	return
}

func (this Recti) Intersects(other *Recti) (test bool, intersection *Recti) {
	test = C.sfIntRect_intersects(this.toCPtr(), other.toCPtr(), intersection.toCPtr()) == 1
	return
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Rectf) fromC(rect C.sfFloatRect) {
	this.Left = float32(rect.left)
	this.Top = float32(rect.top)
	this.Width = float32(rect.width)
	this.Height = float32(rect.height)
}

func (this *Recti) fromC(rect C.sfIntRect) {
	this.Left = int(rect.left)
	this.Top = int(rect.top)
	this.Width = int(rect.width)
	this.Height = int(rect.height)
}

func (this *Recti) toC() C.sfIntRect {
	return C.sfIntRect{left: C.int(this.Left), top: C.int(this.Top), width: C.int(this.Width), height: C.int(this.Height)}
}

func (this *Rectf) toC() C.sfFloatRect {
	return C.sfFloatRect{left: C.float(this.Left), top: C.float(this.Top), width: C.float(this.Width), height: C.float(this.Height)}
}

func (this *Recti) toCPtr() *C.sfIntRect {
	return (*C.sfIntRect)(unsafe.Pointer(this))
}

func (this *Rectf) toCPtr() *C.sfFloatRect {
	return (*C.sfFloatRect)(unsafe.Pointer(this))
}

/////////////////////////////////////
///		Testing
/////////////////////////////////////

func sizeofRecti() int {
	return int(C.getSizeRecti())
}

func sizeofRectf() int {
	return int(C.getSizeRectf())
}