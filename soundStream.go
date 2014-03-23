// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

/*
#include <SFML/Audio/SoundStream.h>

extern sfSoundStream* sfSoundStream_createEx(unsigned int channelCount, unsigned int sampleRate,void* obj);
*/
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

type SoundStream struct {
	cptr         *C.sfSoundStream
	dataCallback SoundStreamDataCallback
	seekCallback SoundStreamSeekCallback
	userData     interface{}
}

type SoundStreamDataCallback func(userData interface{}) (proceed bool, samples []int16)
type SoundStreamSeekCallback func(time time.Duration, userData interface{})

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Create a new sound stream
//
// onGetData    Function called when the stream needs more data (can't be nil)
// onSeek       Function called when the stream seeks (can't be nil)
// channelCount Number of channels to use (1 = mono, 2 = stereo)
// sampleRate   Sample rate of the sound (44100 = CD quality)
// userData     Data to pass to the callback function (can be nil)
func NewSoundStream(onGetData SoundStreamDataCallback, onSeek SoundStreamSeekCallback, channelCount, sampleRate uint, userData interface{}) (*SoundStream, error) {
	if onGetData == nil || onSeek == nil {
		return nil, errors.New("NewSoundStream: Callbacks cannot be nil")
	}

	soundStream := &SoundStream{
		dataCallback: onGetData,
		seekCallback: onSeek,
		userData:     userData,
	}

	if cptr := C.sfSoundStream_createEx(C.uint(channelCount), C.uint(sampleRate), unsafe.Pointer(soundStream)); cptr != nil {
		soundStream.cptr = cptr
		runtime.SetFinalizer(soundStream, (*SoundStream).destroy)
		return soundStream, nil
	}

	return nil, genericError
}

// Destroy a sound stream
func (this *SoundStream) destroy() {
	C.sfSoundStream_destroy(this.cptr)
}

// Start or resume playing a sound stream
//
// This function starts the stream if it was stopped, resumes
// it if it was paused, and restarts it from beginning if it
// was it already playing.
// This function uses its own thread so that it doesn't block
// the rest of the program while the music is played.
func (this *SoundStream) Play() {
	C.sfSoundStream_play(this.cptr)
}

// Pause a sound stream
//
// This function pauses the stream if it was playing,
// otherwise (stream already paused or stopped) it has no effect.
func (this *SoundStream) Pause() {
	C.sfSoundStream_pause(this.cptr)
}

// Stop playing a sound stream
//
// This function stops the stream if it was playing or paused,
// and does nothing if it was already stopped.
// It also resets the playing position (unlike SoundStream.Pause).
func (this *SoundStream) Stop() {
	C.sfSoundStream_stop(this.cptr)
}

// Get the current status of a sound stream (stopped, paused, playing)
func (this *SoundStream) GetStatus() SoundStatus {
	return (SoundStatus)(C.sfSoundStream_getStatus(this.cptr))
}

// Return the number of channels of a sound stream
//
// 1 channel means a mono sound, 2 means stereo, etc.
func (this *SoundStream) GetChannelCount() uint {
	return (uint)(C.sfSoundStream_getChannelCount(this.cptr))
}

// Get the sample rate of a sound stream
//
// The sample rate is the number of audio samples played per
// second. The higher, the better the quality.
func (this *SoundStream) GetSampleRate() uint {
	return (uint)(C.sfSoundStream_getSampleRate(this.cptr))
}

// Set the pitch of a sound stream
//
// The pitch represents the perceived fundamental frequency
// of a sound; thus you can make a stream more acute or grave
// by changing its pitch. A side effect of changing the pitch
// is to modify the playing speed of the stream as well.
// The default value for the pitch is 1.
func (this *SoundStream) SetPitch(pitch float32) {
	C.sfSoundStream_setPitch(this.cptr, C.float(pitch))
}

// Set the volume of a sound stream
//
// The volume is a value between 0 (mute) and 100 (full volume).
// The default value for the volume is 100.
func (this *SoundStream) SetVolume(volume float32) {
	C.sfSoundStream_setVolume(this.cptr, C.float(volume))
}

// Set the 3D position of a sound stream in the audio scene
//
// Only streams with one channel (mono streams) can be
// spatialized.
// The default position of a stream is (0, 0, 0).
func (this *SoundStream) SetPosition(position Vector3f) {
	C.sfSoundStream_setPosition(this.cptr, position.toC())
}

// Make a sound stream's position relative to the listener or absolute
//
// Making a stream relative to the listener will ensure that it will always
// be played the same way regardless the position of the listener.
// This can be useful for non-spatialized streams, streams that are
// produced by the listener, or streams attached to it.
// The default value is false (position is absolute).
func (this *SoundStream) SetRelativeToListener(relative bool) {
	C.sfSoundStream_setRelativeToListener(this.cptr, goBool2C(relative))
}

// Set the minimum distance of a sound stream
//
// The "minimum distance" of a stream is the maximum
// distance at which it is heard at its maximum volume. Further
// than the minimum distance, it will start to fade out according
// to its attenuation factor. A value of 0 ("inside the head
// of the listener") is an invalid value and is forbidden.
// The default value of the minimum distance is 1.
func (this *SoundStream) SetMinDistance(distance float32) {
	C.sfSoundStream_setMinDistance(this.cptr, C.float(distance))
}

// Set the attenuation factor of a sound stream
//
// The attenuation is a multiplicative factor which makes
// the stream more or less loud according to its distance
// from the listener. An attenuation of 0 will produce a
// non-attenuated stream, i.e. its volume will always be the same
// whether it is heard from near or from far. On the other hand,
// an attenuation value such as 100 will make the stream fade out
// very quickly as it gets further from the listener.
// The default value of the attenuation is 1.
func (this *SoundStream) SetAttenuation(attenuation float32) {
	C.sfSoundStream_setAttenuation(this.cptr, C.float(attenuation))
}

// Change the current playing position of a sound stream
//
// The playing position can be changed when the stream is
// either paused or playing.
func (this *SoundStream) SetPlayingOffset(offset time.Duration) {
	C.sfSoundStream_setPlayingOffset(this.cptr, C.sfTime{microseconds: (C.sfInt64(offset / time.Microsecond))})
}

// Set whether or not a sound stream should loop after reaching the end
//
// If set, the stream will restart from beginning after
// reaching the end and so on, until it is stopped or
// SoundStream.SetLoop(false) is called.
// The default looping state for sound streams is false.
func (this *SoundStream) SetLoop(loop bool) {
	C.sfSoundStream_setLoop(this.cptr, goBool2C(loop))
}

// Get the pitch of a sound stream
func (this *SoundStream) GetPitch() float32 {
	return (float32)(C.sfSoundStream_getPitch(this.cptr))
}

// Get the volume of a sound stream, in the range [0, 100]
func (this *SoundStream) GetVolume() float32 {
	return (float32)(C.sfSoundStream_getVolume(this.cptr))
}

// Get the 3D position of a sound stream in the audio scene
func (this *SoundStream) GetPosition() (pos Vector3f) {
	pos.fromC((C.sfSoundStream_getPosition(this.cptr)))
	return
}

// Tell whether a sound stream's position is relative to the
// listener or is absolute
func (this *SoundStream) IsRelativeToListener() bool {
	return sfBool2Go((C.sfSoundStream_isRelativeToListener(this.cptr)))
}

// Get the minimum distance of a sound stream
func (this *SoundStream) GetMinDistance() float32 {
	return (float32)(C.sfSoundStream_getMinDistance(this.cptr))
}

// Get the attenuation factor of a sound stream
func (this *SoundStream) GetAttenuation() float32 {
	return (float32)(C.sfSoundStream_getAttenuation(this.cptr))
}

// Tell whether or not a sound stream is in loop mode
func (this *SoundStream) GetLoop() bool {
	return sfBool2Go((C.sfSoundStream_getLoop(this.cptr)))
}

// Get the current playing position of a sound stream
func (this *SoundStream) GetPlayingOffset() time.Duration {
	return time.Duration(C.sfSoundStream_getPlayingOffset(this.cptr).microseconds) * time.Microsecond
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

//export go_callbackGetData
func go_callbackGetData(chunk *C.sfSoundStreamChunk, ptr unsafe.Pointer) C.sfBool {
	if (*(*SoundStream)(ptr)).dataCallback != nil {
		r, goChunk := (*(*SoundStream)(ptr)).dataCallback((*(*SoundStream)(ptr)).userData)
		chunk.sampleCount = C.uint(len(goChunk))
		if len(goChunk) > 0 {
			chunk.samples = (*C.sfInt16)(unsafe.Pointer(&goChunk[0]))
		}
		return goBool2C(r)
	}
	return C.sfFalse //stop playback
}

//export go_callbackSeek
func go_callbackSeek(ctime C.sfTime, ptr unsafe.Pointer) {
	if (*(*SoundStream)(ptr)).seekCallback != nil {
		(*(*SoundStream)(ptr)).seekCallback(time.Duration(ctime.microseconds)*time.Microsecond, (*(*SoundStream)(ptr)).userData)
	}
}
