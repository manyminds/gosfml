/*
#############################################
#	GOSFML2
#	Example:
#	Sound, Memory and Zip test
############################################
*/

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	sf "github.com/manyminds/gosfml"
)

func loadFromZip(zipFile, file string) ([]byte, error) {
	var data []byte

	// Read file
	zip, err := zip.OpenReader(zipFile)
	if err != nil {
		return data, err
	}
	defer zip.Close()

	// Search for *file*
	for _, f := range zip.File {
		if f.Name == file {
			// Found
			file, err := f.Open()
			if err != nil {
				return data, err
			}
			defer file.Close()
			data, err = ioutil.ReadAll(file)
			return data, err
		}
	}
	return data, io.EOF // File not found
}

func playSound() {
	// Load from zip
	data, err := loadFromZip("resources.zip", "canary.wav")
	if err != nil {
		log.Fatal(err)
	}

	// Load a sound buffer from a wav data buffer
	buffer, err := sf.NewSoundBufferFromMemory(data)
	if err != nil {
		log.Fatal(err)
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
	// Load from zip
	data, err := loadFromZip("resources.zip", "orchestral.ogg")
	if err != nil {
		log.Fatal(err)
	}

	// Load an ogg music data buffer
	music, err := sf.NewMusicFromMemory(data)
	if err != nil {
		log.Fatal(err)
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
