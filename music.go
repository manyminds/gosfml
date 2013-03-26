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

func NewMusicFromFile(file string) (music *Music, err error) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	music = &Music{C.sfMusic_createFromFile(cFile)}
	runtime.SetFinalizer(music, (*Music).Destroy)

	if music.cptr == nil {
		err = errors.New("NewMusicFromFile: Cannot load music " + file)
	}

	return
}

func NewMusicFromMemory(data []byte) (music *Music, err error) {
	if len(data) > 0 {
		music = &Music{C.sfMusic_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data)))}
		runtime.SetFinalizer(music, (*Music).Destroy)

		if music.cptr == nil {
			err = errors.New("NewMusicFromMemory: Cannot load music")
		}
		return
	}
	return
}

func (this *Music) Destroy() {
	C.sfMusic_destroy(this.cptr)
	this.cptr = nil
}

func (this *Music) Play() {
	C.sfMusic_play(this.cptr)
}

func (this *Music) Pause() {
	C.sfMusic_pause(this.cptr)
}

func (this *Music) Stop() {
	C.sfMusic_stop(this.cptr)
}

func (this *Music) SetLoop(loop bool) {
	C.sfMusic_setLoop(this.cptr, goBool2C(loop))
}

func (this *Music) GetStatus() SoundStatus {
	return SoundStatus(C.sfMusic_getStatus(this.cptr))
}

func (this *Music) SetPitch(pitch float32) {
	C.sfMusic_setPitch(this.cptr, C.float(pitch))
}

func (this *Music) SetVolume(volume float32) {
	C.sfMusic_setVolume(this.cptr, C.float(volume))
}

func (this *Music) SetPosition(pos Vector3f) {
	C.sfMusic_setPosition(this.cptr, pos.toC())
}

func (this *Music) SetRelativeToListener(relative bool) {
	C.sfMusic_setRelativeToListener(this.cptr, goBool2C(relative))
}

func (this *Music) SetMinDistance(distance float32) {
	C.sfMusic_setMinDistance(this.cptr, C.float(distance))
}

func (this *Music) SetAttenuation(attenuation float32) {
	C.sfMusic_setAttenuation(this.cptr, C.float(attenuation))
}

func (this *Music) SetPlayingOffset(offset time.Duration) {
	C.sfMusic_setPlayingOffset(this.cptr, C.sfTime{microseconds: (C.sfInt64(offset / time.Microsecond))})
}

func (this *Music) GetPitch() float32 {
	return float32(C.sfMusic_getPitch(this.cptr))
}

func (this *Music) GetVolume() float32 {
	return float32(C.sfMusic_getVolume(this.cptr))
}

func (this *Music) GetPosition() (pos Vector3f) {
	pos.fromC(C.sfMusic_getPosition(this.cptr))
	return
}

func (this *Music) IsRelativeToListner() bool {
	return sfBool2Go(C.sfMusic_isRelativeToListener(this.cptr))
}

func (this *Music) GetMinDistance() float32 {
	return float32(C.sfMusic_getMinDistance(this.cptr))
}

func (this *Music) GetAttenuation() float32 {
	return float32(C.sfMusic_getAttenuation(this.cptr))
}

func (this *Music) GetPlayingOffset() time.Duration {
	return time.Duration(C.sfMusic_getPlayingOffset(this.cptr).microseconds) * time.Microsecond
}

func (this *Music) GetSampleRate() uint {
	return uint(C.sfMusic_getSampleRate(this.cptr))
}

func (this *Music) GetChannelCount() uint {
	return uint(C.sfMusic_getChannelCount(this.cptr))
}

func (this *Music) GetDuration() time.Duration {
	return time.Duration(C.sfMusic_getPlayingOffset(this.cptr).microseconds) * time.Microsecond
}
