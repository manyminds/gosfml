package GoSFML2

/*
 #include <SFML/Audio.h> 
*/
import "C"

import (
	"runtime"
	"time"
)

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	SoundStatus_Stopped = iota ///< Sound / music is not playing
	SoundStatus_Paused         ///< Sound / music is paused
	SoundStatus_Playing        ///< Sound / music is playing
)

type SoundStatus int

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Sound struct {
	cptr   *C.sfSound
	buffer *SoundBuffer
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func CreateSound() *Sound {
	sound := &Sound{C.sfSound_create(), nil}
	runtime.SetFinalizer(sound, (*Sound).Destroy)
	return sound
}

func (this *Sound) Copy() *Sound {
	sound := &Sound{C.sfSound_copy(this.cptr), this.buffer}
	runtime.SetFinalizer(sound, (*Sound).Destroy)
	return sound
}

func (this *Sound) Destroy() {
	C.sfSound_destroy(this.cptr)
	this.cptr = nil
}

func (this *Sound) Play() {
	C.sfSound_play(this.cptr)
}

func (this *Sound) Pause() {
	C.sfSound_pause(this.cptr)
}

func (this *Sound) Stop() {
	C.sfSound_stop(this.cptr)
}

func (this *Sound) SetBuffer(buffer *SoundBuffer) {
	C.sfSound_setBuffer(this.cptr, buffer.cptr)
	this.buffer = buffer
}

func (this *Sound) GetBuffer() *SoundBuffer {
	return this.buffer
}

func (this *Sound) SetLoop(loop bool) {
	C.sfSound_setLoop(this.cptr, goBool2C(loop))
}

func (this *Sound) GetStatus() SoundStatus {
	return SoundStatus(C.sfSound_getStatus(this.cptr))
}

func (this *Sound) SetPitch(pitch float32) {
	C.sfSound_setPitch(this.cptr, C.float(pitch))
}

func (this *Sound) SetVolume(volume float32) {
	C.sfSound_setVolume(this.cptr, C.float(volume))
}

func (this *Sound) SetPosition(pos Vector3f) {
	C.sfSound_setPosition(this.cptr, pos.toC())
}

func (this *Sound) SetRelativeToListener(relative bool) {
	C.sfSound_setRelativeToListener(this.cptr, goBool2C(relative))
}

func (this *Sound) SetMinDistance(distance float32) {
	C.sfSound_setMinDistance(this.cptr, C.float(distance))
}

func (this *Sound) SetAttenuation(attenuation float32) {
	C.sfSound_setAttenuation(this.cptr, C.float(attenuation))
}

func (this *Sound) SetPlayingOffset(offset time.Duration) {
	C.sfSound_setPlayingOffset(this.cptr, C.sfMicroseconds(C.sfInt64(offset/time.Microsecond)))
}

func (this *Sound) GetPitch() float32 {
	return float32(C.sfSound_getPitch(this.cptr))
}

func (this *Sound) GetVolume() float32 {
	return float32(C.sfSound_getVolume(this.cptr))
}

func (this *Sound) GetPosition() (pos Vector3f) {
	pos.fromC(C.sfSound_getPosition(this.cptr))
	return
}

func (this *Sound) IsRelativeToListner() bool {
	return sfBool2Go(C.sfSound_isRelativeToListener(this.cptr))
}

func (this *Sound) GetMinDistance() float32 {
	return float32(C.sfSound_getMinDistance(this.cptr))
}

func (this *Sound) GetAttenuation() float32 {
	return float32(C.sfSound_getAttenuation(this.cptr))
}

func (this *Sound) GetPlayingOffset() time.Duration {
	return time.Duration(C.sfTime_asMicroseconds(C.sfSound_getPlayingOffset(this.cptr))) * time.Microsecond
}
