// Copyright (C) 2012 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

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
