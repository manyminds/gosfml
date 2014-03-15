// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Audio/Sound.h>
import "C"

import (
	"runtime"
	"time"
)

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	SoundStatusStopped = iota ///< Sound / music is not playing
	SoundStatusPaused         ///< Sound / music is paused
	SoundStatusPlaying        ///< Sound / music is playing
)

type SoundStatus int

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Sound struct {
	cptr   *C.sfSound
	buffer *SoundBuffer
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Create a new sound with a given SoundBuffer
func NewSound(buffer *SoundBuffer) *Sound {
	sound := &Sound{C.sfSound_create(), nil}
	sound.SetBuffer(buffer)
	runtime.SetFinalizer(sound, (*Sound).destroy)

	return sound
}

// Create a new sound by copying an existing one
func (this *Sound) Copy() *Sound {
	sound := &Sound{C.sfSound_copy(this.cptr), this.buffer}
	runtime.SetFinalizer(sound, (*Sound).destroy)
	return sound
}

// Destroy a sound
func (this *Sound) destroy() {
	C.sfSound_destroy(this.cptr)
}

// Start or resume playing a sound
//
// This function starts the sound if it was stopped, resumes
// it if it was paused, and restarts it from beginning if it
// was it already playing.
// This function uses its own thread so that it doesn't block
// the rest of the program while the sound is played.
func (this *Sound) Play() {
	C.sfSound_play(this.cptr)
}

// Pause a sound
//
// This function pauses the sound if it was playing,
// otherwise (sound already paused or stopped) it has no effect.
func (this *Sound) Pause() {
	C.sfSound_pause(this.cptr)
}

// Stop playing a sound
//
// This function stops the sound if it was playing or paused,
// and does nothing if it was already stopped.
// It also resets the playing position (unlike Sound.Pause).
func (this *Sound) Stop() {
	C.sfSound_stop(this.cptr)
}

// Set the source buffer containing the audio data to play
func (this *Sound) SetBuffer(buffer *SoundBuffer) {
	C.sfSound_setBuffer(this.cptr, buffer.toCPtr())
	this.buffer = buffer
}

// Get the audio buffer attached to a sound
func (this *Sound) GetBuffer() *SoundBuffer {
	return this.buffer
}

// Set whether or not a sound should loop after reaching the end
//
// If set, the sound will restart from beginning after
// reaching the end and so on, until it is stopped or
// Sound.SetLoop(false) is called.
// The default looping state for sounds is false.
func (this *Sound) SetLoop(loop bool) {
	C.sfSound_setLoop(this.cptr, goBool2C(loop))
}

// Get the current status of a sound (stopped, paused, playing)
func (this *Sound) GetStatus() SoundStatus {
	return SoundStatus(C.sfSound_getStatus(this.cptr))
}

// Set the pitch of a sound
//
// The pitch represents the perceived fundamental frequency
// of a sound; thus you can make a sound more acute or grave
// by changing its pitch. A side effect of changing the pitch
// is to modify the playing speed of the sound as well.
// The default value for the pitch is 1.
func (this *Sound) SetPitch(pitch float32) {
	C.sfSound_setPitch(this.cptr, C.float(pitch))
}

// Set the volume of a sound
//
// The volume is a value between 0 (mute) and 100 (full volume).
// The default value for the volume is 100.
func (this *Sound) SetVolume(volume float32) {
	C.sfSound_setVolume(this.cptr, C.float(volume))
}

// Set the 3D position of a sound in the audio scene
//
// Only sounds with one channel (mono sounds) can be
// spatialized.
// The default position of a sound is (0, 0, 0).
func (this *Sound) SetPosition(pos Vector3f) {
	C.sfSound_setPosition(this.cptr, pos.toC())
}

// Make the sound's position relative to the listener or absolute
//
// Making a sound relative to the listener will ensure that it will always
// be played the same way regardless the position of the listener.
// This can be useful for non-spatialized sounds, sounds that are
// produced by the listener, or sounds attached to it.
// The default value is false (position is absolute).
func (this *Sound) SetRelativeToListener(relative bool) {
	C.sfSound_setRelativeToListener(this.cptr, goBool2C(relative))
}

// Set the minimum distance of a sound
//
// The "minimum distance" of a sound is the maximum
// distance at which it is heard at its maximum volume. Further
// than the minimum distance, it will start to fade out according
// to its attenuation factor. A value of 0 ("inside the head
// of the listener") is an invalid value and is forbidden.
// The default value of the minimum distance is 1.
func (this *Sound) SetMinDistance(distance float32) {
	C.sfSound_setMinDistance(this.cptr, C.float(distance))
}

// Set the attenuation factor of a sound
//
// The attenuation is a multiplicative factor which makes
// the sound more or less loud according to its distance
// from the listener. An attenuation of 0 will produce a
// non-attenuated sound, i.e. its volume will always be the same
// whether it is heard from near or from far. On the other hand,
// an attenuation value such as 100 will make the sound fade out
// very quickly as it gets further from the listener.
// The default value of the attenuation is 1.
func (this *Sound) SetAttenuation(attenuation float32) {
	C.sfSound_setAttenuation(this.cptr, C.float(attenuation))
}

// Change the current playing position of a sound
//
// The playing position can be changed when the sound is
// either paused or playing.
func (this *Sound) SetPlayingOffset(offset time.Duration) {
	C.sfSound_setPlayingOffset(this.cptr, C.sfTime{microseconds: (C.sfInt64(offset / time.Microsecond))})
}

// Tell whether or not a sound is in loop mode
func (this *Sound) GetLoop() bool {
	return sfBool2Go(C.sfSound_getLoop(this.cptr))
}

// Get the pitch of a sound
func (this *Sound) GetPitch() float32 {
	return float32(C.sfSound_getPitch(this.cptr))
}

// Get the volume of a sound
func (this *Sound) GetVolume() float32 {
	return float32(C.sfSound_getVolume(this.cptr))
}

// Get the 3D position of a sound in the audio scene
func (this *Sound) GetPosition() (pos Vector3f) {
	pos.fromC(C.sfSound_getPosition(this.cptr))
	return
}

// Tell whether a sound's position is relative to the
// listener or is absolute
func (this *Sound) IsRelativeToListner() bool {
	return sfBool2Go(C.sfSound_isRelativeToListener(this.cptr))
}

// Get the minimum distance of a sound
func (this *Sound) GetMinDistance() float32 {
	return float32(C.sfSound_getMinDistance(this.cptr))
}

// Get the attenuation factor of a sound
func (this *Sound) GetAttenuation() float32 {
	return float32(C.sfSound_getAttenuation(this.cptr))
}

// Get the current playing position of a sound
func (this *Sound) GetPlayingOffset() time.Duration {
	return time.Duration(C.sfSound_getPlayingOffset(this.cptr).microseconds) * time.Microsecond
}
