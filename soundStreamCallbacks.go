// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

/*
#include <SFML/Audio/SoundStream.h>

// cgo export declarations
sfBool go_callbackGetData(sfSoundStreamChunk* chunk, void* ptr);
void go_callbackSeek(sfTime t, void* ptr);

// C callbacks
sfBool bridge_soundStreamGetData(sfSoundStreamChunk* chunk, void* ptr)
{
	return go_callbackGetData((void*)chunk, ptr);
}

void bridge_soundStreamSeek(sfTime time, void* ptr)
{
	go_callbackSeek(time, ptr);
}

// create a sfSoundStream using the callbacks above.
sfSoundStream* sfSoundStream_createEx(unsigned int channelCount, unsigned int sampleRate, void* obj)
{
	return sfSoundStream_create(bridge_soundStreamGetData, bridge_soundStreamSeek, channelCount, sampleRate, obj);
}
*/
import "C"
