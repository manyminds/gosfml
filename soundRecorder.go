// Copyright (C) 2012 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

/*
#include <SFML/Audio/SoundRecorder.h>
extern sfSoundRecorder* sfSoundRecorder_createEx(void*);
extern void copyData(void*, void*, size_t);
extern size_t sizeofInt16();
*/
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type SoundRecorder struct {
	cptr *C.sfSoundRecorder

	startCallback    SoundRecorderCallbackStart
	stopCallback     SoundRecorderCallbackStop
	progressCallback SoundRecorderCallbackProgress
}

type SoundRecorderCallbackStart func() bool
type SoundRecorderCallbackStop func()
type SoundRecorderCallbackProgress func([]int16) bool

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Construct a new sound recorder from callback functions
//
// 	onStart   Callback function which will be called when a new capture starts (can be nil)
// 	onProcess Callback function which will be called each time there's audio data to process (cannot be nil)
// 	onStop    Callback function which will be called when the current capture stops (can be nil)
func NewSoundRecorder(onStart SoundRecorderCallbackStart, onProgress SoundRecorderCallbackProgress, onStop SoundRecorderCallbackStop) (*SoundRecorder, error) {
	soundRecorder := &SoundRecorder{}
	soundRecorder.startCallback = onStart
	soundRecorder.stopCallback = onStop
	soundRecorder.progressCallback = onProgress
	soundRecorder.cptr = C.sfSoundRecorder_createEx(unsafe.Pointer(soundRecorder))

	runtime.SetFinalizer(soundRecorder, (*SoundRecorder).destroy)

	if soundRecorder.cptr == nil || onProgress == nil {
		return nil, errors.New("NewSoundRecorder: Cannot create SoundRecorder")
	}

	return soundRecorder, nil
}

// Destroy an existing SoundRecorder
func (this *SoundRecorder) destroy() {
	C.sfSoundRecorder_destroy(this.cptr)
	this.cptr = nil
	this.startCallback = nil
	this.stopCallback = nil
	this.progressCallback = nil
}

// The sampleRate parameter defines the number of audio samples
// captured per second. The higher, the better the quality
// (for example, 44100 samples/sec is CD quality).
// This function uses its own thread so that it doesn't block
// the rest of the program while the capture runs.
// Please note that only one capture can happen at the same time.
//
// 	sampleRate    Desired capture rate, in number of samples per second
func (this *SoundRecorder) Start(sampleRate uint) {
	C.sfSoundRecorder_start(this.cptr, C.uint(sampleRate))
}

// Stop the capture of a sound recorder
func (this *SoundRecorder) Stop() {
	C.sfSoundRecorder_stop(this.cptr)
}

// Get the sample rate of a sound recorder
//
// The sample rate defines the number of audio samples
// captured per second. The higher, the better the quality
// (for example, 44100 samples/sec is CD quality).
func (this *SoundRecorder) GetSampleRate() uint {
	return uint(C.sfSoundRecorder_getSampleRate(this.cptr))
}

// Check if the system supports audio capture
//
// This function should always be called before using
// the audio capture features. If it returns false, then
// any attempt to use SoundRecorder will fail.
func SoundRecorderIsAvailable() bool {
	return sfBool2Go(C.sfSoundRecorder_isAvailable())
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

//export go_callbackStart
func go_callbackStart(ptr unsafe.Pointer) C.sfBool {
	if (*(*SoundRecorder)(ptr)).startCallback != nil {
		return goBool2C((*(*SoundRecorder)(ptr)).startCallback())
	}
	return C.sfTrue
}

//export go_callbackStop
func go_callbackStop(ptr unsafe.Pointer) {
	if (*(*SoundRecorder)(ptr)).stopCallback != nil {
		(*(*SoundRecorder)(ptr)).stopCallback()
	}
}

//export go_callbackProgress
func go_callbackProgress(data *C.sfInt16, count C.size_t, ptr unsafe.Pointer) C.sfBool {
	buffer := make([]int16, count)
	if count > 0 {
		C.copyData(unsafe.Pointer(data), unsafe.Pointer(&buffer[0]), count*C.sizeofInt16())
	}
	return goBool2C((*(*SoundRecorder)(ptr)).progressCallback(buffer))
}
