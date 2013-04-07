// Copyright (c) 2012 krepa098 (krepa098 at gmail dot com)
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
// Permission is granted to anyone to use this software for any purpose, including commercial applications, 
// and to alter it and redistribute it freely, subject to the following restrictions:
// 	1.	The origin of this software must not be misrepresented; you must not claim that you wrote the original software. 
//		If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
// 	2. 	Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
// 	3. 	This notice may not be removed or altered from any source distribution.

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

//MISSING: 	sfSoundBuffer_createFromMemory
//			sfSoundBuffer_createFromStream
//			sfSoundBuffer_createFromSamples
//			sfSoundBuffer_getSamples

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type SoundBuffer struct {
	cptr *C.sfSoundBuffer
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Create a new sound buffer and load it from a file
//
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
func NewSoundBufferFromFile(file string) (*SoundBuffer, error) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	buffer := &SoundBuffer{C.sfSoundBuffer_createFromFile(cFile)}
	runtime.SetFinalizer(buffer, (*SoundBuffer).Destroy)

	if buffer.cptr == nil {
		return nil, errors.New("NewSoundBufferFromFile: Cannot create SoundBuffer")
	}

	return buffer, nil
}

// Create a new sound buffer and load it from a file in memory
//
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
//
// 	data: Slice of file data
func NewSoundBufferFromMemory(data []byte) (buffer *SoundBuffer, err error) {
	if len(data) > 0 {
		buffer = &SoundBuffer{C.sfSoundBuffer_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data)))}
		runtime.SetFinalizer(buffer, (*SoundBuffer).Destroy)

		if buffer.cptr == nil {
			err = errors.New("NewSoundBufferFromMemory: Cannot create SoundBuffer")
		}
		return
	}
	return nil, errors.New("NewSoundBufferFromMemory: NewSoundBufferFromMemory: no data")
}

// Create a new sound buffer and load it from an array of samples in memory
//
// The assumed format of the audio samples is 16 bits signed integer
// (int16).
//
// 	samples:      Slice of samples
// 	sampleCount:  Number of samples in the array
// 	channelCount: Number of channels (1 = mono, 2 = stereo, ...)
// 	sampleRate:   Sample rate (number of samples to play per second)
func NewSoundBufferFromSamples(samples []int16, sampleCount, channelCount, sampleRate uint) (buffer *SoundBuffer, err error) {
	if len(samples) > 0 {
		buffer = &SoundBuffer{C.sfSoundBuffer_createFromSamples((*C.sfInt16)(unsafe.Pointer(&samples[0])),C.size_t(sampleCount),C.uint(channelCount),C.uint(sampleRate))}
		runtime.SetFinalizer(buffer, (*SoundBuffer).Destroy)

		if buffer.cptr == nil {
			err = errors.New("NewSoundBufferFromSamples: Cannot create SoundBuffer")
		}
		return
	}
	return nil, errors.New("NewSoundBufferFromSamples: NewSoundBufferFromMemory: no data")
}

// Create a new sound buffer by copying an existing one
func (this *SoundBuffer) Copy() *SoundBuffer {
	buffer := &SoundBuffer{C.sfSoundBuffer_copy(this.cptr)}
	runtime.SetFinalizer(buffer, (*SoundBuffer).Destroy)
	return buffer
}

// Destroy a sound buffer
func (this *SoundBuffer) Destroy() {
	C.sfSoundBuffer_destroy(this.cptr)
	this.cptr = nil
}

// Save a sound buffer to an audio file
//
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
//
// 	file: Path of the sound file to write
func (this *SoundBuffer) SaveToFile(file string) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))

	C.sfSoundBuffer_saveToFile(this.cptr, cFile)
}

// Get the number of samples stored in a sound buffer
//
// The array of samples can be accessed with the
// SoundBuffer.GetSamples function.
func (this *SoundBuffer) GetSampleCount() uint {
	return uint(C.sfSoundBuffer_getSampleCount(this.cptr))
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
