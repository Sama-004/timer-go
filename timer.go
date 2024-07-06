package main

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func playsound() {
	f, err := os.Open("/home/sama/projects/go/sound.mp3")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		fmt.Println("Playback finished")
	})))

	select {}
}

func main() {
	fmt.Println("test")
	var inputTime int
	//Timer logic
	var choice string
	fmt.Println("Enter s for seconds or m for minutes")

	fmt.Scan(&choice)

	if choice == "s" {
		fmt.Print("Enter your time in seconds:")
		fmt.Scan(&inputTime)
		duration := time.Duration(inputTime) * time.Second
		time.Sleep(duration)
		fmt.Printf("Timer is set for the duration %d ...waiting\n", inputTime)
	} else if choice == "m" {
		fmt.Print("Enter your time in minutes:")
		fmt.Scan(&inputTime)
		duration := time.Duration(inputTime) * time.Minute
		time.Sleep(duration)
		fmt.Printf("Timer is set for the duration %d ...waiting\n", inputTime)
	} else {
		fmt.Println("Not a valid option.")
	}
	fmt.Println("Time's up") //play sound here
	playsound()
	//maybe show the remaining time instead of sleeping
}
