// Copyright (c) 2012 krepa098 (krepa098 at gmail dot com)
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
// Permission is granted to anyone to use this software for any purpose, including commercial applications,
// and to alter it and redistribute it freely, subject to the following restrictions:
// 	1.	The origin of this software must not be misrepresented; you must not claim that you wrote the original software.
//			If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
// 	2. Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
// 	3. This notice may not be removed or altered from any source distribution.

package gosfml2

/*
#include <SFML/Audio/SoundRecorder.h>
extern sfSoundRecorder* sfSoundRecorder_createEx(void*);
extern sfInt16 accessSampleData(sfInt16*,size_t);
*/
import "C"

import (
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
// 	onStart   Callback function which will be called when a new capture starts
// 	onProcess Callback function which will be called each time there's audio data to process
// 	onStop    Callback function which will be called when the current capture stops
func NewSoundRecorder(onStart SoundRecorderCallbackStart, onProgress SoundRecorderCallbackProgress, onStop SoundRecorderCallbackStop) *SoundRecorder {
	soundRecorder := &SoundRecorder{}
	soundRecorder.startCallback = onStart
	soundRecorder.stopCallback = onStop
	soundRecorder.progressCallback = onProgress

	soundRecorder.cptr = C.sfSoundRecorder_createEx(unsafe.Pointer(soundRecorder))

	runtime.SetFinalizer(soundRecorder, (*SoundRecorder).destroy)

	return soundRecorder
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

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

// private proxy functions
func (this *SoundRecorder) callCallackStart() bool {
	return this.startCallback()
}

func (this *SoundRecorder) callCallackStop() {
	this.stopCallback()
}

func (this *SoundRecorder) callCallackProgress(data []int16) bool {
	return this.progressCallback(data)
}

func SoundRecorderIsAvailable() bool {
	return sfBool2Go(C.sfSoundRecorder_isAvailable())
}

//export go_callbackStart
func go_callbackStart(ptr unsafe.Pointer) C.sfBool {
	return goBool2C((*(*SoundRecorder)(ptr)).callCallackStart())
}

//export go_callbackStop
func go_callbackStop(ptr unsafe.Pointer) {
	(*(*SoundRecorder)(ptr)).callCallackStop()
}

//export go_callbackProgress
func go_callbackProgress(data *C.sfInt16, count C.size_t, ptr unsafe.Pointer) C.sfBool {
	buffer := make([]int16, count)
	for i := 0; i < int(count); i++ {
		buffer[i] = int16(C.accessSampleData(data, C.size_t(i)))
	}

	return goBool2C((*(*SoundRecorder)(ptr)).callCallackProgress(buffer))
}
