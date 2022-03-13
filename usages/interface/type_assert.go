package main

import "fmt"

type TapePlayer string

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder string

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}
func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

type Player interface { // Define an interface type
	Play(string) // Require a Play method with a string parameter
	Stop()       // Require a Stop method
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

// After an interface value, type a dot, followed by a pair of parentheses with the concrete type.
// It is what you’re asserting the value’s concrete type is.
// Once you’ve used a type assertion to get a value of a concrete type back, you can call methods on it that are
// defined on that type, but are not part of the interface.

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()

	// player is the Interface value, TapeRecorder is the Asserted type.
	recorder, ok := player.(TapeRecorder) // Assign the second return value to a variable.

	// ok is an optional return default value that indicates whether the assertion was successful or not.
	if ok {
		recorder.Record() // If the original type was TapeRecorder, call Record on the value.
	} else {
		fmt.Println("Player was not a TapeRecorder.") // If not, report that the assertion failed.
	}

}

// The interface can be used as embed.
func main() {
	mixType := []string{"Mamba forever", "Kobe R.I.P", "8 to 24"}
	var player Player = TapePlayer("a")
	playList(player, mixType)
	TryOut(player) // Assertion is failure.

	player = TapeRecorder("b")
	playList(player, mixType)
	TryOut(player) // Assertion is successful.
}
