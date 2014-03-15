// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

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
func NewSoundBufferRecorder() (*SoundBufferRecorder, error) {
	if cptr := C.sfSoundBufferRecorder_create(); cptr != nil {
		soundBufferRecorder := &SoundBufferRecorder{cptr}
		runtime.SetFinalizer(soundBufferRecorder, (*SoundBufferRecorder).destroy)

		return soundBufferRecorder, nil
	}

	return nil, genericError
}

// Destroy an existing SoundBufferRecorder
func (this *SoundBufferRecorder) destroy() {
	C.sfSoundBufferRecorder_destroy(this.cptr)
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
