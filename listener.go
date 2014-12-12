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

// Set the orientation of the forward vector in the scene
//
// The direction (also called "at vector") is the vector
// pointing forward from the listener's perspective. Together
// with the up vector, it defines the 3D orientation of the
// listener in the scene. The direction vector doesn't
// have to be normalized.
// The default listener's direction is (0, 0, -1).
//
// 	direction New listener's direction
func ListenerSetDirection(direction Vector3f) {
	C.sfListener_setDirection(direction.toC())
}

// Get the listener's forward vector (not normalized)
func ListenerGetDirection() (direction Vector3f) {
	direction.fromC(C.sfListener_getDirection())
	return
}

// Set the upward vector of the listener in the scene
//
// The up vector is the vector that points upward from the
// listener's perspective. Together with the direction, it
// defines the 3D orientation of the listener in the scene.
// The up vector doesn't have to be normalized.
// The default listener's up vector is (0, 1, 0). It is usually
// not necessary to change it, especially in 2D scenarios.
//
// 	upVec New listener's up vector
func ListenerSetUpVector(upVec Vector3f) {
	C.sfListener_setUpVector(upVec.toC())
}

// Get the current upward vector of the listener in the scene
func ListenerGetUpVector() (upVec Vector3f) {
	upVec.fromC(C.sfListener_getUpVector())
	return
}
