// Copyright (C) 2012 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Audio/Listener.h>
import "C"

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Change the global volume of all the sounds and musics
//
// The volume is a number between 0 and 100; it is combined with
// the individual volume of each sound / music.
// The default value for the volume is 100 (maximum).
//
// 	volume: New global volume, in the range [0, 100]
func ListenerSetGlobalVolume(volume float32) {
	C.sfListener_setGlobalVolume(C.float(volume))
}

// Get the current value of the global volume
//
// return Current global volume, in the range [0, 100]
func ListenerGetGlobalVolume() float32 {
	return float32(C.sfListener_getGlobalVolume())
}

// Set the position of the listener in the scene
//
// The default listener's position is (0, 0, 0).
//
// 	position: New position of the listener
func ListenerSetPosition(pos Vector3f) {
	C.sfListener_setPosition(pos.toC())
}

// Get the current position of the listener in the scene
func ListenerGetPosition() (pos Vector3f) {
	pos.fromC(C.sfListener_getPosition())
	return
}

// Set the orientation of the listener in the scene
//
// The orientation defines the 3D axes of the listener
// (left, up, front) in the scene. The orientation vector
// doesn't have to be normalized.
// The default listener's orientation is (0, 0, -1).
//
// 	position: New direction of the listener
func ListenerSetDirection(dir Vector3f) {
	C.sfListener_setPosition(dir.toC())
}

// Get the current orientation of the listener in the scene
func ListenerGetDirection() (dir Vector3f) {
	dir.fromC(C.sfListener_getDirection())
	return
}
