// Copyright (C) 2012 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

/*
#include <SFML/Audio/SoundRecorder.h>

extern sfBool go_callbackStart(void* ptr);
extern void go_callbackStop(void* ptr);
extern sfBool go_callbackProgress(const sfInt16* data, size_t count, void* ptr);

sfBool callGo_soundRecorderStart(void* ptr) { return go_callbackStart(ptr); }
void callGo_soundRecorderStop(void* ptr) { go_callbackStop(ptr); }
sfBool callGo_soundRecorderProgress(const sfInt16* data, size_t count, void* ptr) { return go_callbackProgress(data,count,ptr); }

sfSoundRecorder* sfSoundRecorder_createEx(void* obj) { return sfSoundRecorder_create(callGo_soundRecorderStart,callGo_soundRecorderProgress,callGo_soundRecorderStop,obj); }
*/
import "C"
