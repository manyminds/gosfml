// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Audio/SoundBuffer.h>
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

type SoundBuffer struct {
	cptr *C.sfSoundBuffer
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Copy a C soundBuffer into a Go SoundBuffer
func newSoundBufferFromPtr(cbuffer *C.sfSoundBuffer) *SoundBuffer {
	buffer := &SoundBuffer{C.sfSoundBuffer_copy(cbuffer)}
	runtime.SetFinalizer(buffer, (*SoundBuffer).destroy)

	return buffer
}

// Create a new sound buffer and load it from a file
//
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
func NewSoundBufferFromFile(file string) (*SoundBuffer, error) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))

	if cptr := C.sfSoundBuffer_createFromFile(cFile); cptr != nil {
		buffer := &SoundBuffer{cptr}
		runtime.SetFinalizer(buffer, (*SoundBuffer).destroy)

		return buffer, nil
	}

	return nil, genericError
}

// Create a new sound buffer and load it from a file in memory
//
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
//
// 	data: Slice of file data
func NewSoundBufferFromMemory(data []byte) (*SoundBuffer, error) {
	if len(data) == 0 {
		return nil, errors.New("NewSoundBufferFromMemory: len(data)==0")
	}

	if cptr := C.sfSoundBuffer_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data))); cptr != nil {
		buffer := &SoundBuffer{cptr}
		runtime.SetFinalizer(buffer, (*SoundBuffer).destroy)

		return buffer, nil
	}

	return nil, genericError
}

// Create a new sound buffer and load it from an array of samples in memory
//
// The assumed format of the audio samples is 16 bits signed integer
// (int16).
//
// 	samples:      Slice of samples
// 	channelCount: Number of channels (1 = mono, 2 = stereo, ...)
// 	sampleRate:   Sample rate (number of samples to play per second)
func NewSoundBufferFromSamples(samples []int16, channelCount, sampleRate uint) (*SoundBuffer, error) {
	if len(samples) == 0 {
		return nil, errors.New("NewSoundBufferFromSamples: len(data)==0")
	}

	if cptr := C.sfSoundBuffer_createFromSamples((*C.sfInt16)(unsafe.Pointer(&samples[0])), C.size_t(len(samples)), C.uint(channelCount), C.uint(sampleRate)); cptr != nil {
		buffer := &SoundBuffer{cptr}
		runtime.SetFinalizer(buffer, (*SoundBuffer).destroy)

		return buffer, nil
	}
	return nil, genericError
}

// Create a new sound buffer by copying an existing one
func (this *SoundBuffer) Copy() *SoundBuffer {
	buffer := &SoundBuffer{C.sfSoundBuffer_copy(this.cptr)}
	runtime.SetFinalizer(buffer, (*SoundBuffer).destroy)
	return buffer
}

// Destroy a sound buffer
func (this *SoundBuffer) destroy() {
	C.sfSoundBuffer_destroy(this.cptr)
}

// Save a sound buffer to an audio file
//
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
//
// 	file: Path of the sound file to write
func (this *SoundBuffer) SaveToFile(file string) error {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))

	if !sfBool2Go(C.sfSoundBuffer_saveToFile(this.cptr, cFile)) {
		return genericError
	}
	return nil
}

// Get the number of samples stored in a sound buffer
//
// The array of samples can be accessed with the
// SoundBuffer.GetSamples function.
func (this *SoundBuffer) GetSampleCount() uint {
	return uint(C.sfSoundBuffer_getSampleCount(this.cptr))
}

// Get the slice of audio samples stored in a sound buffer
//
// The format of the returned samples is 16 bits signed integer
// (int16).
func (this *SoundBuffer) GetSamples() []int16 {
	data := make([]int16, this.GetSampleCount())
	if len(data) > 0 {
		memcopy(unsafe.Pointer(&data[0]), unsafe.Pointer(C.sfSoundBuffer_getSamples(this.cptr)), len(data)*int(unsafe.Sizeof(int16(0))))
	}
	return data
}

// Get the sample rate of a sound buffer
//
// The sample rate is the number of samples played per second.
// The higher, the better the quality (for example, 44100
// samples/s is CD quality).
func (this *SoundBuffer) GetSampleRate() uint {
	return uint(C.sfSoundBuffer_getSampleRate(this.cptr))
}

// Get the number of channels used by a sound buffer
//
// If the sound is mono then the number of channels will
// be 1, 2 for stereo, etc.
func (this *SoundBuffer) GetChannelCount() uint {
	return uint(C.sfSoundBuffer_getChannelCount(this.cptr))
}

// Get the total duration of a sound buffer
func (this *SoundBuffer) GetDuration() time.Duration {
	return time.Duration(C.sfSoundBuffer_getDuration(this.cptr).microseconds) * time.Microsecond
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *SoundBuffer) toCPtr() *C.sfSoundBuffer {
	if this != nil {
		return this.cptr
	}
	return nil
}
