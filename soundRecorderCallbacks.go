// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

/*
#include <SFML/Audio/SoundRecorder.h>

// cgo exports
extern sfBool go_callbackStart(void* ptr);
extern void go_callbackStop(void* ptr);
extern sfBool go_callbackProgress(const sfInt16* data, size_t count, void* ptr);

// C callbacks
sfBool c_soundRecorderStart(void* ptr)
{
	return go_callbackStart(ptr);
}

void c_soundRecorderStop(void* ptr)
{
	go_callbackStop(ptr);
}

sfBool c_soundRecorderProgress(const sfInt16* data, size_t count, void* ptr)
{
	return go_callbackProgress(data, count, ptr);
}

// create a sfSoundRecorder using the callbacks above.
sfSoundRecorder* sfSoundRecorder_createEx(void* obj) {
	return sfSoundRecorder_create(c_soundRecorderStart, c_soundRecorderProgress, c_soundRecorderStop,obj);
}
*/
import "C"
