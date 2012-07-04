package GoSFML2

// #include <SFML/System.h>
import "C"

/////////////////////////////////////
///		WRAPPING HELPERS
/////////////////////////////////////

func sfBool2Go(b C.sfBool) bool {
	return b == 1
}

func goBool2C(b bool) C.sfBool {
	if b {
		return C.sfBool(1)
	}
	return C.sfBool(0)
}
