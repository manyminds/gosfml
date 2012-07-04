package GoSFML2

// #include <SFML/Graphics.h>
import "C"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Color struct {
	R byte
	G byte
	B byte
	A byte // 0=transparent
}

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

var (
	Color_Black = Color{0, 0, 0, 255}
	Color_Red   = Color{255, 0, 0, 255}
	Color_Green = Color{0, 255, 0, 255}
	Color_Blue  = Color{0, 0, 255, 255}
	Color_White = Color{255, 255, 255, 255}
)

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Color) fromC(color C.sfColor) {
	this.R = byte(color.r)
	this.G = byte(color.g)
	this.B = byte(color.b)
	this.A = byte(color.a)
}

func (this *Color) toC() C.sfColor {
	return C.sfColor{r: C.sfUint8(this.R), g: C.sfUint8(this.G), b: C.sfUint8(this.B), a: C.sfUint8(this.A)}
}
