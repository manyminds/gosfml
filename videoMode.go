package GoSFML2

/*
 #include <SFML/Window.h>
 #include <stdlib.h>

 sfVideoMode videoModeAt(size_t index, sfVideoMode* modes) {
		return modes[index];
 }

*/
import "C"
import "unsafe"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type VideoMode struct {
	Width        uint ///< Video mode width, in pixels
	Height       uint ///< Video mode height, in pixels
	BitsPerPixel uint ///< Video mode pixel depth, in bits per pixels
}

type fullscreenModesCount struct {
	count uint
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func (this *VideoMode) SetAsDesktopVideoMode() {
	this.fromC(C.sfVideoMode_getDesktopMode())
}

func (this *VideoMode) IsValid() bool {
	return sfBool2Go(C.sfVideoMode_isValid(this.toC()))
}

func (this *VideoMode) GetFullscreenModes() []VideoMode {
	c := &fullscreenModesCount{}
	cVideoModes := C.sfVideoMode_getFullscreenModes((*C.size_t)(unsafe.Pointer(c)))

	modes := make([]VideoMode, c.count)
	for i := uint(0); i < c.count; i++ {
		modes[i].fromC(C.videoModeAt(C.size_t(i), cVideoModes))
	}
	return modes
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *VideoMode) fromC(videoMode C.sfVideoMode) {
	this.Width = uint(videoMode.width)
	this.Height = uint(videoMode.height)
	this.BitsPerPixel = uint(videoMode.bitsPerPixel)
}

func (this *VideoMode) toC() C.sfVideoMode {
	return C.sfVideoMode{width: C.uint(this.Width), height: C.uint(this.Height), bitsPerPixel: C.uint(this.BitsPerPixel)}
}
