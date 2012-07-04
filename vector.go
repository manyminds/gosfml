package GoSFML2

// #include <SFML/System.h>
import "C"
import "math"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Vector2i struct {
	X, Y int
}

type Vector2u struct {
	X, Y uint
}

type Vector2f struct {
	X, Y float32
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

/////////////////////////////////////
// Vector2i

func (this *Vector2i) Plus(other Vector2i) Vector2i {
	return Vector2i{X: this.X + other.X, Y: this.Y + other.Y}
}

func (this *Vector2i) Minus(other Vector2i) Vector2i {
	return Vector2i{X: this.X - other.X, Y: this.Y - other.Y}
}

func (this *Vector2i) Length() float32 {
	return float32(math.Sqrt(float64(this.X*this.X + this.Y*this.Y)))
}

/////////////////////////////////////
// Vector2u

func (this *Vector2u) Plus(other Vector2u) Vector2u {
	return Vector2u{X: this.X + other.X, Y: this.Y + other.Y}
}

func (this *Vector2u) Minus(other Vector2u) Vector2u {
	return Vector2u{X: this.X - other.X, Y: this.Y - other.Y}
}

func (this *Vector2u) Length() float32 {
	return float32(math.Sqrt(float64(this.X*this.X + this.Y*this.Y)))
}

/////////////////////////////////////
// Vector2f

func (this *Vector2f) Plus(other Vector2f) Vector2f {
	return Vector2f{X: this.X + other.X, Y: this.Y + other.Y}
}

func (this *Vector2f) Minus(other Vector2f) Vector2f {
	return Vector2f{X: this.X - other.X, Y: this.Y - other.Y}
}

func (this *Vector2f) Length() float32 {
	return float32(math.Sqrt(float64(this.X*this.X + this.Y*this.Y)))
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Vector2i) fromC(vec C.sfVector2i) {
	this.X = int(vec.x)
	this.Y = int(vec.y)
}

func (this *Vector2u) fromC(vec C.sfVector2u) {
	this.X = uint(vec.x)
	this.Y = uint(vec.y)
}

func (this *Vector2f) fromC(vec C.sfVector2f) {
	this.X = float32(vec.x)
	this.Y = float32(vec.y)
}

func (this *Vector2i) toC() C.sfVector2i {
	return C.sfVector2i{C.int(this.X), C.int(this.Y)}
}

func (this *Vector2u) toC() C.sfVector2u {
	return C.sfVector2u{C.uint(this.X), C.uint(this.Y)}
}

func (this *Vector2f) toC() C.sfVector2f {
	return C.sfVector2f{C.float(this.X), C.float(this.Y)}
}
