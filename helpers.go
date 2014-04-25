// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

/*
#include <SFML/System.h>
#include <string.h> //memcpy

void nextChar(sfUint32** ptr)
{
	++(*ptr);
}
*/
import "C"

import (
	"errors"
	"sync"
	"unsafe"
)

var (
	globalCtx   = NewContext()
	globalMutex sync.Mutex

	//As SFML does not provide useful error codes, we just return a generic error message
	genericError = errors.New("Error: See stderr for more details")
)

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

// Convert a utf32 C string to a go string
func utf32CString2Go(cstr *C.sfUint32) string {
	var str string

	for ptr := cstr; *ptr != 0; C.nextChar(&ptr) {
		str += string(rune(uint32(*ptr)))
	}

	return str
}

// Returns a null terminated UTF32 representation of str.
func strToRunes(str string) []rune {
	return append([]rune(str), rune(0))
}

func globalCtxSetActive(active bool) {
	if active {
		globalMutex.Lock()
	}

	globalCtx.SetActive(active)

	if !active {
		globalMutex.Unlock()
	}
}

func memcopy(dest, src unsafe.Pointer, size int) {
	C.memcpy(dest, src, C.size_t(size))
}
