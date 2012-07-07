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

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Music struct {
	cptr *C.sfMusic
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func CreateMusicFromFile(file string) *Music {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	music := &Music{C.sfMusic_createFromFile(cFile)}
	runtime.SetFinalizer(music, (*Music).Destroy)
	return music
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

func (this *Music) SetPlayingOffset(time time.Duration) {
	C.sfMusic_setPlayingOffset(this.cptr, C.sfSeconds(C.float(float32(time.Seconds()))))
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

//return: time in milliseconds
func (this *Music) GetPlayingOffset() uint {
	time := C.sfMusic_getPlayingOffset(this.cptr)
	return uint(C.sfTime_asMilliseconds(time))
}
