package GoSFML2

/*
 #include <SFML/Audio.h> 
*/
import "C"

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func Listener_SetGlobalVolume(volume float32) {
	C.sfListener_setGlobalVolume(C.float(volume))
}

func Listener_GetGlobalVolume() float32 {
	return float32(C.sfListener_getGlobalVolume())
}

func Listener_SetPosition(pos Vector3f) {
	C.sfListener_setPosition(pos.toC())
}

func Listener_GetPosition() (pos Vector3f) {
	pos.fromC(C.sfListener_getPosition())
	return
}

func Listener_SetDirection(dir Vector3f) {
	C.sfListener_setPosition(dir.toC())
}

func Listener_GetDirection() (dir Vector3f) {
	dir.fromC(C.sfListener_getDirection())
	return
}
