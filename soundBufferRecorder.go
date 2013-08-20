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

// #include <SFML/Audio/SoundBufferRecorder.h>
import "C"

import (
	"runtime"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type SoundBufferRecorder struct {
	cptr *C.sfSoundBufferRecorder
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

/// Create a new sound buffer recorder
func NewSoundBufferRecorder() *SoundBufferRecorder {
	soundBufferRecorder := &SoundBufferRecorder{cptr: C.sfSoundBufferRecorder_create()}
	runtime.SetFinalizer(soundBufferRecorder, (*SoundBufferRecorder).destroy)

	return soundBufferRecorder
}

// Destroy an existing SoundBufferRecorder
func (this *SoundBufferRecorder) destroy() {
	C.sfSoundBufferRecorder_destroy(this.cptr)
	this.cptr = nil
}

// Start the capture of a sound recorder recorder
//
// The sampleRate parameter defines the number of audio samples
// captured per second. The higher, the better the quality
// (for example, 44100 samples/sec is CD quality).
// This function uses its own thread so that it doesn't block
// the rest of the program while the capture runs.
// Please note that only one capture can happen at the same time.
//
// 	soundBufferRecorder Sound buffer recorder object
// 	sampleRate          Desired capture rate, in number of samples per second
func (this *SoundBufferRecorder) Start(sampleRate uint) {
	C.sfSoundBufferRecorder_start(this.cptr, C.uint(sampleRate))
}

// Stop the capture of a sound recorder
func (this *SoundBufferRecorder) Stop() {
	C.sfSoundBufferRecorder_stop(this.cptr)
}

// Get the sample rate of a sound buffer recorder
//
// The sample rate defines the number of audio samples
// captured per second. The higher, the better the quality
// (for example, 44100 samples/sec is CD quality).
func (this *SoundBufferRecorder) GetSampleRate() uint {
	return uint(C.sfSoundBufferRecorder_getSampleRate(this.cptr))
}

// Get the sound buffer containing the captured audio data
//
// The sound buffer is valid only after the capture has ended.
func (this *SoundBufferRecorder) GetBuffer() *SoundBuffer {
	return newSoundBufferFromPtr(C.sfSoundBufferRecorder_getBuffer(this.cptr))
}
