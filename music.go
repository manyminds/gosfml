// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Audio/Music.h>
// #include <stdlib.h>
import "C"

import (
	"errors"
	"runtime"
	"time"
	"unsafe"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Music struct {
	cptr *C.sfMusic
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Create a new music and load it from a file
//
///This function doesn't start playing the music (call
// Music.Play to do so).
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
//
// 	file: Path of the music file to open
func NewMusicFromFile(file string) (*Music, error) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))

	if cptr := C.sfMusic_createFromFile(cFile); cptr != nil {
		music := &Music{cptr}
		runtime.SetFinalizer(music, (*Music).destroy)
		return music, nil
	}

	return nil, genericError
}

// Create a new music and load it from a file in memory
//
// This function doesn't start playing the music (call
// Music.Play to do so).
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
//
// 	data: Slice of files data
func NewMusicFromMemory(data []byte) (*Music, error) {
	if len(data) == 0 {
		return nil, errors.New("NewMusicFromMemory: len(data)==0")
	}

	if cptr := C.sfMusic_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data))); cptr != nil {
		music := &Music{cptr}
		runtime.SetFinalizer(music, (*Music).destroy)

		return music, nil
	}

	return nil, genericError
}

// Destroy a music
func (this *Music) destroy() {
	C.sfMusic_destroy(this.cptr)
}

// Start or resume playing a music
//
// This function starts the music if it was stopped, resumes
// it if it was paused, and restarts it from beginning if it
// was it already playing.
// This function uses its own thread so that it doesn't block
// the rest of the program while the music is played.
func (this *Music) Play() {
	C.sfMusic_play(this.cptr)
}

// Pause a music
//
// This function pauses the music if it was playing,
// otherwise (music already paused or stopped) it has no effect.
func (this *Music) Pause() {
	C.sfMusic_pause(this.cptr)
}

// Stop playing a music
//
// This function stops the music if it was playing or paused,
// and does nothing if it was already stopped.
// It also resets the playing position (unlike Music.Pause).
func (this *Music) Stop() {
	C.sfMusic_stop(this.cptr)
}

// Set whether or not a music should loop after reaching the end
//
// If set, the music will restart from beginning after
// reaching the end and so on, until it is stopped or
// Music.SetLoop(false) is called.
// The default looping state for musics is false.
//
// 	loop:  true to play in loop, false to play once
func (this *Music) SetLoop(loop bool) {
	C.sfMusic_setLoop(this.cptr, goBool2C(loop))
}

// Get the current status of a music (stopped, paused, playing)
func (this *Music) GetStatus() SoundStatus {
	return SoundStatus(C.sfMusic_getStatus(this.cptr))
}

// Set the pitch of a music
//
// The pitch represents the perceived fundamental frequency
// of a sound; thus you can make a music more acute or grave
// by changing its pitch. A side effect of changing the pitch
// is to modify the playing speed of the music as well.
// The default value for the pitch is 1.
//
// 	pitch: New pitch to apply to the music
func (this *Music) SetPitch(pitch float32) {
	C.sfMusic_setPitch(this.cptr, C.float(pitch))
}

// Set the volume of a music
//
// The volume is a value between 0 (mute) and 100 (full volume).
// The default value for the volume is 100.
//
// 	volume: Volume of the music
func (this *Music) SetVolume(volume float32) {
	C.sfMusic_setVolume(this.cptr, C.float(volume))
}

// Set the 3D position of a music in the audio scene
//
// Only musics with one channel (mono musics) can be
// spatialized.
// The default position of a music is (0, 0, 0).
//
// 	position: Position of the music in the scene
func (this *Music) SetPosition(pos Vector3f) {
	C.sfMusic_setPosition(this.cptr, pos.toC())
}

// Make a musics's position relative to the listener or absolute
//
// Making a music relative to the listener will ensure that it will always
// be played the same way regardless the position of the listener.
// This can be useful for non-spatialized musics, musics that are
// produced by the listener, or musics attached to it.
// The default value is false (position is absolute).
//
// 	relative: true to set the position relative, false to set it absolute
func (this *Music) SetRelativeToListener(relative bool) {
	C.sfMusic_setRelativeToListener(this.cptr, goBool2C(relative))
}

// Set the minimum distance of a music
//
// The "minimum distance" of a music is the maximum
// distance at which it is heard at its maximum volume. Further
// than the minimum distance, it will start to fade out according
// to its attenuation factor. A value of 0 ("inside the head
// of the listener") is an invalid value and is forbidden.
// The default value of the minimum distance is 1.
//
// 	distance: New minimum distance of the music
func (this *Music) SetMinDistance(distance float32) {
	C.sfMusic_setMinDistance(this.cptr, C.float(distance))
}

// Set the attenuation factor of a music
//
// The attenuation is a multiplicative factor which makes
// the music more or less loud according to its distance
// from the listener. An attenuation of 0 will produce a
// non-attenuated music, i.e. its volume will always be the same
// whether it is heard from near or from far. On the other hand,
// an attenuation value such as 100 will make the music fade out
// very quickly as it gets further from the listener.
// The default value of the attenuation is 1.
//
// 	attenuation: New attenuation factor of the music
func (this *Music) SetAttenuation(attenuation float32) {
	C.sfMusic_setAttenuation(this.cptr, C.float(attenuation))
}

// Change the current playing position of a music
//
// The playing position can be changed when the music is
// either paused or playing.
//
// 	timeOffset: New playing position
func (this *Music) SetPlayingOffset(offset time.Duration) {
	C.sfMusic_setPlayingOffset(this.cptr, C.sfTime{microseconds: (C.sfInt64(offset / time.Microsecond))})
}

// Get the pitch of a music
func (this *Music) GetPitch() float32 {
	return float32(C.sfMusic_getPitch(this.cptr))
}

// Get the volume of a music
func (this *Music) GetVolume() float32 {
	return float32(C.sfMusic_getVolume(this.cptr))
}

// Get the 3D position of a music in the audio scene
func (this *Music) GetPosition() (pos Vector3f) {
	pos.fromC(C.sfMusic_getPosition(this.cptr))
	return
}

// Tell whether a music's position is relative to the
// listener or is absolute
func (this *Music) IsRelativeToListner() bool {
	return sfBool2Go(C.sfMusic_isRelativeToListener(this.cptr))
}

// Get the minimum distance of a music
func (this *Music) GetMinDistance() float32 {
	return float32(C.sfMusic_getMinDistance(this.cptr))
}

// Get the attenuation factor of a music
func (this *Music) GetAttenuation() float32 {
	return float32(C.sfMusic_getAttenuation(this.cptr))
}

// Get the current playing position of a music
func (this *Music) GetPlayingOffset() time.Duration {
	return time.Duration(C.sfMusic_getPlayingOffset(this.cptr).microseconds) * time.Microsecond
}

// Get the sample rate of a music
//
// The sample rate is the number of audio samples played per
// second. The higher, the better the quality.
func (this *Music) GetSampleRate() uint {
	return uint(C.sfMusic_getSampleRate(this.cptr))
}

// Return the number of channels of a music
//
// 1 channel means a mono sound, 2 means stereo, etc
func (this *Music) GetChannelCount() uint {
	return uint(C.sfMusic_getChannelCount(this.cptr))
}

// Get the total duration of a music
func (this *Music) GetDuration() time.Duration {
	return time.Duration(C.sfMusic_getPlayingOffset(this.cptr).microseconds) * time.Microsecond
}
