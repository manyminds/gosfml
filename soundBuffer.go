/*
Copyright (c) 2012 krepa098 (krepa098 at gmail dot com)
This software is provided 'as-is', without any express or implied warranty.
In no event will the authors be held liable for any damages arising from the use of this software.
Permission is granted to anyone to use this software for any purpose, including commercial applications, 
and to alter it and redistribute it freely, subject to the following restrictions:
	1.	The origin of this software must not be misrepresented; you must not claim that you wrote the original software. 
		If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
	2. 	Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
	3. 	This notice may not be removed or altered from any source distribution.
*/

package GoSFML2

/*
 #include <SFML/Audio.h> 
 #include <stdlib.h>
*/
import "C"

import (
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

func NewSoundBufferFromFile(file string) *SoundBuffer {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	buffer := &SoundBuffer{C.sfSoundBuffer_createFromFile(cFile)}
	runtime.SetFinalizer(buffer, (*SoundBuffer).Destroy)
	return buffer
}

func (this *SoundBuffer) Copy() *SoundBuffer {
	buffer := &SoundBuffer{C.sfSoundBuffer_copy(this.cptr)}
	runtime.SetFinalizer(buffer, (*SoundBuffer).Destroy)
	return buffer
}

func (this *SoundBuffer) Destroy() {
	C.sfSoundBuffer_destroy(this.cptr)
	this.cptr = nil
}

func (this *SoundBuffer) SaveToFile(file string) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))

	C.sfSoundBuffer_saveToFile(this.cptr, cFile)
}

func (this *SoundBuffer) GetSampleCount() uint {
	return uint(C.sfSoundBuffer_getSampleCount(this.cptr))
}

func (this *SoundBuffer) GetSampleRate() uint {
	return uint(C.sfSoundBuffer_getSampleRate(this.cptr))
}

func (this *SoundBuffer) GetChannelCount() uint {
	return uint(C.sfSoundBuffer_getChannelCount(this.cptr))
}

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
