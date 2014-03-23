/*
*********************************
*	GOSFML2
*	SFML Example: Sound/Music
*	Ported from C++ to Go
*********************************
 */

package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"fmt"
	"os"
	"time"
)

func playSound() {
	// Load a sound buffer from a wav file
	buffer, err := sf.NewSoundBufferFromFile("resources/canary.wav")

	if err != nil {
		panic(err)
	}

	// Display sound informations
	fmt.Println("canary.wav :")
	fmt.Println(" ", buffer.GetDuration())
	fmt.Println(" ", buffer.GetSampleRate(), " samples/sec")
	fmt.Println(" ", buffer.GetChannelCount(), " channels")

	// Create a sound instance and play it
	sound := sf.NewSound(buffer)
	sound.Play()

	// Loop while the sound is playing
	for sound.GetStatus() == sf.SoundStatusPlaying {
		// Leave some CPU time for other processes
		time.Sleep(100)

		// Display the playing position
		fmt.Println("Playing...", sound.GetPlayingOffset())
	}
}

func playMusic() {
	// Load an ogg music file
	music, err := sf.NewMusicFromFile("resources/orchestral.ogg")
	if err != nil {
		panic(err)
	}

	// Display sound informations
	fmt.Println("orchestral.ogg :")
	fmt.Println(" ", music.GetDuration())
	fmt.Println(" ", music.GetSampleRate(), " samples/sec")
	fmt.Println(" ", music.GetChannelCount(), " channels")

	// Play it
	music.Play()

	// Loop while the music is playing
	for music.GetStatus() == sf.SoundStatusPlaying {
		// Leave some CPU time for other processes
		time.Sleep(100)

		// Display the playing position
		fmt.Println("Playing...", music.GetPlayingOffset())
	}
}

func main() {
	// Play a sound
	playSound()

	// Play a music
	playMusic()

	fmt.Println("Press enter to exit...")
	var buffer [1]byte
	os.Stdin.Read(buffer[:])
}
