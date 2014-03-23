// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

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
// This function activates the new context.
func NewContext() *Context {
	context := &Context{C.sfContext_create()}
	runtime.SetFinalizer(context, (*Context).destroy)
	return context
}

// Destroy a context
func (this *Context) destroy() {
	C.sfContext_destroy(this.cptr)
}

// Activate or deactivate explicitely a context
//
// 	active: true to activate, false to deactivate
func (this *Context) SetActive(active bool) {
	C.sfContext_setActive(this.cptr, goBool2C(active))
}
