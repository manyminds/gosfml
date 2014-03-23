/*
#############################################
#	GOSFML2
#	Example: Audio Recording
#############################################
*/

package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"fmt"
	"time"
)

var recordedSamples = make([]int16, 0, 44100*10*10)

func main() {
	fmt.Println("Testing sound recording")

	if sf.SoundRecorderIsAvailable() {
		fmt.Println("Sound recording available: yes")
		fmt.Println("Recording audio for 10 seconds...")

		//create a new soundBufferRecorder
		recorderBuffer, err := sf.NewSoundBufferRecorder()

		if err != nil {
			panic(err)
		}

		recorderBuffer.Start(44100) //CD quality

		//wait 10s
		time.Sleep(10 * time.Second)

		//stop recording (SoundBufferRecorder.GetBuffer() is now valid)
		recorderBuffer.Stop()

		//create a new sound using the buffer from the soundBufferRecorder
		buffer := recorderBuffer.GetBuffer()
		sound := sf.NewSound(buffer)
		fmt.Println("Playback!")
		sound.Play()

		//wait until finished
		for sound.GetStatus() != sf.SoundStatusStopped {
			time.Sleep(500 * time.Millisecond)
		}

		fmt.Println("Finished, saving to output.wav")
		buffer.SaveToFile("output1.wav")

		fmt.Println("Testing custom soundRecorder (Recording for 10 seconds)")

		soundRecorder, err := sf.NewSoundRecorder(StartCallback, ProgressCallback, StopCallback, nil)

		if err != nil {
			panic(err)
		}

		soundRecorder.Start(44100)

		//wait 10s
		time.Sleep(10 * time.Second)

		soundRecorder.Stop()

		buffer, _ = sf.NewSoundBufferFromSamples(recordedSamples, 1, 44100)
		buffer.SaveToFile("output2.wav")
		sound = sf.NewSound(buffer)

		fmt.Println("Playback!")
		sound.Play()

		//wait until finished
		for sound.GetStatus() != sf.SoundStatusStopped {
			time.Sleep(500 * time.Millisecond)
		}

		fmt.Println("Done")
	} else {
		fmt.Println("ERROR: Sound recording not available")
	}
}

func StartCallback(interface{}) bool {
	fmt.Println("start")
	return true
}

func ProgressCallback(data []int16, userData interface{}) bool {
	fmt.Println(len(data))
	recordedSamples = append(recordedSamples, data...)
	return true //continue recording
}

func StopCallback(interface{}) {
	fmt.Println("stop")
}
