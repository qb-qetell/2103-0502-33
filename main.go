package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main () {
	// Clear all the characters on the screen
	screen.Clear ()

	for {
		screen.Clear ()
		// Moves the cursor to the top-left position of the screen
		screen.MoveTopLeft ()

		// Animate the time always in the same position
		fmt.Println (time.Now ())

		time.Sleep (time.Second)
	}
}
