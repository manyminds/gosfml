/*
#############################################
#	GOSFML2
#	Example: Sound Generator (sine waves)
#############################################
*/

package main

import (
	"fmt"
	"math"
	"time"

	sf "github.com/manyminds/gosfml"
)

const (
	SampleRate  = 44100
	Amplitude   = math.MaxInt16 / 2
	TimePerFreq = 0.100
)

var (
	Frequency = 25.0 //Hz
	t         = 0.0  //time in seconds
)

func main() {
	//callback
	//this function sends samples to the sound buffer
	dataCallback := func(userData interface{}) (bool, []int16) {
		T := 1.0 / Frequency         //1 period in sec
		s := T * SampleRate          //samples per period
		nt := int(TimePerFreq/T + 1) //number of periods in TimePerFreq
		n := int(float64(nt) * s)    //number of samples

		//generate samples
		data := make([]int16, n)

		for i := 0; i < len(data); i++ {
			//generate sin-wave
			data[i] = int16(math.Sin(2.0*math.Pi*Frequency*t) * Amplitude)
			t += 1.0 / float64(SampleRate)
		}

		fmt.Println(Frequency, " Hz")
		Frequency += 25.0
		t = 0.0

		if Frequency <= 2000.0 {
			return true, data //continue
		} else {
			return false, data //stop stream
		}

	}

	//callback
	seekCallback := func(time.Duration, interface{}) {
		/* unused */
	}

	//create Sound Stream and start playing
	stream, _ := sf.NewSoundStream(dataCallback, seekCallback, 1, SampleRate, nil)
	stream.Play()

	fmt.Println("Generating sine-waves from 0Hz up to 2kHz")
	for stream.GetStatus() == sf.SoundStatusPlaying {
		time.Sleep(1 * time.Millisecond)
	}
}
