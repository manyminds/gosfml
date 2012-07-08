package GoSFML2

/*
 #include <SFML/Audio.h> 
 #include <stdlib.h>
*/
import "C"

import (
	"runtime"
	"unsafe"
	"time"
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

func CreateSoundBufferFromFile(file string) *SoundBuffer {
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
	return time.Duration(C.sfTime_asMicroseconds(C.sfSoundBuffer_getDuration(this.cptr))) * time.Microsecond
}
