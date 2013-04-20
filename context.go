// Copyright (c) 2012 krepa098 (krepa098 at gmail dot com)
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
// Permission is granted to anyone to use this software for any purpose, including commercial applications, 
// and to alter it and redistribute it freely, subject to the following restrictions:
// 	1.	The origin of this software must not be misrepresented; you must not claim that you wrote the original software. 
//		If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
// 	2. 	Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
// 	3. 	This notice may not be removed or altered from any source distribution.

package gosfml2

// #include <SFML/Window/Context.h> 
import "C"
import "runtime"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Context struct {
	cptr *C.sfContext
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Create a new context
//
///This function activates the new context.
func NewContext() *Context {
	context := &Context{C.sfContext_create()}
	runtime.SetFinalizer(context, (*Context).destroy)
	return context
}

// Destroy a context
func (this *Context) destroy() {
	C.sfContext_destroy(this.cptr)
	this.cptr = nil
}

// Activate or deactivate explicitely a context
//
// 	active: true to activate, false to deactivate
func (this *Context) SetActive(active bool) {
	C.sfContext_setActive(this.cptr, goBool2C(active))

}
