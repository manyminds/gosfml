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

package gosfml2

import (
	"testing"
	"unsafe"
)

func TestStructSizes(t *testing.T) {
	if int(unsafe.Sizeof(Vector2i{})) != sizeofVector2i() {
		t.Fatal("Vector2i size mismatch: ", unsafe.Sizeof(Vector2i{}), " != ", sizeofVector2i())
	}

	if int(unsafe.Sizeof(Vector2u{})) != sizeofVector2u() {
		t.Fatal("Vector2u size mismatch: ", unsafe.Sizeof(Vector2u{}), " != ", sizeofVector2u())
	}

	if int(unsafe.Sizeof(Vector2f{})) != sizeofVector2f() {
		t.Fatal("Vector2f size mismatch: ", unsafe.Sizeof(Vector2f{}), " != ", sizeofVector2f())
	}

	if int(unsafe.Sizeof(Vector3f{})) != sizeofVector3f() {
		t.Fatal("Vector3f size mismatch: ", unsafe.Sizeof(Vector3f{}), " != ", sizeofVector3f())
	}

	if int(unsafe.Sizeof(IntRect{})) != sizeofIntRect() {
		t.Fatal("IntRect size mismatch: ", unsafe.Sizeof(IntRect{}), " != ", sizeofIntRect())
	}

	if int(unsafe.Sizeof(FloatRect{})) != sizeofFloatRect() {
		t.Fatal("FloatRect size mismatch: ", unsafe.Sizeof(FloatRect{}), " != ", sizeofFloatRect())
	}

	if int(unsafe.Sizeof(Color{})) != sizeofColor() {
		t.Fatal("Color size mismatch: ", unsafe.Sizeof(Color{}), " != ", sizeofColor())
	}

	if int(unsafe.Sizeof(ContextSettings{})) != sizeofContextSettings() {
		t.Fatal("ContextSettings size mismatch: ", unsafe.Sizeof(ContextSettings{}), " != ", sizeofContextSettings())
	}
}
