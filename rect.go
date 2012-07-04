package GoSFML2

// #include <SFML/Graphics.h>
import "C"
import "unsafe"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Rectf struct {
	Left   float32
	Top    float32
	Width  float32
	Height float32
}

type Recti struct {
	Left   int
	Top    int
	Width  int
	Height int
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func (this *Rectf) Contains(x, y float32) bool {
	return sfBool2Go(C.sfFloatRect_contains(this.toCPtr(), C.float(x), C.float(y)))
}

func (this *Recti) Contains(x, y int) bool {
	return C.sfIntRect_contains(this.toCPtr(), C.int(x), C.int(y)) == 1
}

func (this *Rectf) Intersects(other *Rectf) (test bool, intersection *Rectf) {
	test = C.sfFloatRect_intersects(this.toCPtr(), other.toCPtr(), intersection.toCPtr()) == 1
	return
}

func (this *Recti) Intersects(other *Recti) (test bool, intersection *Recti) {
	test = C.sfIntRect_intersects(this.toCPtr(), other.toCPtr(), intersection.toCPtr()) == 1
	return
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Rectf) fromC(rect C.sfFloatRect) {
	this.Left = float32(rect.left)
	this.Top = float32(rect.top)
	this.Width = float32(rect.width)
	this.Height = float32(rect.height)
}

func (this *Recti) fromC(rect C.sfIntRect) {
	this.Left = int(rect.left)
	this.Top = int(rect.top)
	this.Width = int(rect.width)
	this.Height = int(rect.height)
}

func (this *Recti) toC() C.sfIntRect {
	return C.sfIntRect{C.int(this.Left), C.int(this.Top), C.int(this.Width), C.int(this.Height)}
}

func (this *Rectf) toC() C.sfFloatRect {
	return C.sfFloatRect{C.float(this.Left), C.float(this.Top), C.float(this.Width), C.float(this.Height)}
}

func (this *Recti) toCPtr() *C.sfIntRect {
	return (*C.sfIntRect)(unsafe.Pointer(this))
}

func (this *Rectf) toCPtr() *C.sfFloatRect {
	return (*C.sfFloatRect)(unsafe.Pointer(this))
}
