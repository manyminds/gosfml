// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

/*
#include <SFML/Audio/SoundRecorder.h>

// cgo export declarations
sfBool go_callbackStart(void* ptr);
void go_callbackStop(void* ptr);
sfBool go_callbackProgress(const sfInt16* data, size_t count, void* ptr);

// C callbacks
sfBool bridge_soundRecorderStart(void* ptr)
{
	return go_callbackStart(ptr);
}

void bridge_soundRecorderStop(void* ptr)
{
	go_callbackStop(ptr);
}

sfBool bridge_soundRecorderProgress(const sfInt16* data, size_t count, void* ptr)
{
	return go_callbackProgress(data, count, ptr);
}

// create a sfSoundRecorder using the callbacks above.
sfSoundRecorder* sfSoundRecorder_createEx(void* obj)
{
	return sfSoundRecorder_create(bridge_soundRecorderStart, bridge_soundRecorderProgress, bridge_soundRecorderStop,obj);
}

*/
import "C"
