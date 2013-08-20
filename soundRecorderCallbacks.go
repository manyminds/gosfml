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

/*
#include <SFML/Audio/SoundRecorder.h>

extern sfBool go_callbackStart(void* ptr);
extern void go_callbackStop(void* ptr);
extern sfBool go_callbackProgress(const sfInt16* data, size_t count, void* ptr);

sfInt16 accessSampleData(sfInt16* data,size_t index) { return data[index]; }

sfBool callGo_soundRecorderStart(void* ptr) { return go_callbackStart(ptr); }
void callGo_soundRecorderStop(void* ptr) { go_callbackStop(ptr); }
sfBool callGo_soundRecorderProgress(const sfInt16* data, size_t count, void* ptr) { return go_callbackProgress(data,count,ptr); }

sfSoundRecorder* sfSoundRecorder_createEx(void* obj) { return sfSoundRecorder_create(callGo_soundRecorderStart,callGo_soundRecorderProgress,callGo_soundRecorderStop,obj); }
*/
import "C"
