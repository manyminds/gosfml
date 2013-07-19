// Copyright (c) 2012 krepa098 (krepa098 at gmail dot com)
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
// Permission is granted to anyone to use this software for any purpose, including commercial applications,
// and to alter it and redistribute it freely, subject to the following restrictions:
// 	1.	The origin of this software must not be misrepresented; you must not claim that you wrote the original software.
//			If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
// 	2. Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
// 	3. This notice may not be removed or altered from any source distribution.

package gosfml2

// #include <SFML/System.h>
// void incPtr(sfUint32** ptr)  { ++(*ptr); }
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

func utf32CString2Go(cstr *C.sfUint32) string {
	var str string

	for ptr := cstr; *ptr != 0; C.incPtr(&ptr) {
		str += string(rune(uint32(*ptr)))
	}

	return str
}

func strToRunes(str string) []rune {
	return append([]rune(str), rune(0))
}
