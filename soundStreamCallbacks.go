// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

/*
#include <SFML/Audio/SoundStream.h>

extern sfBool go_callbackGetData(sfSoundStreamChunk* chunk, void* ptr);
extern void go_callbackSeek(sfTime time, void* ptr);

sfBool c_soundStreamGetData(sfSoundStreamChunk* chunk, void* ptr) { return go_callbackGetData(chunk, ptr); }
void c_soundStreamSeek(sfTime time, void* ptr) { go_callbackSeek(time ,ptr); }

sfSoundStream* sfSoundStream_createEx(unsigned int channelCount, unsigned int sampleRate,void* obj) {
	return sfSoundStream_create(c_soundStreamGetData, c_soundStreamSeek, channelCount, sampleRate, obj);
}
*/
import "C"
