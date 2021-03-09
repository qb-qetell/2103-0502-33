package main

import (
	"bufio"
)

func taskThread (c chan [2]string) {
	c <- [2]{"l2", "Startup phase"}
	_ := <- c

	c <- [2]{"l3", `Listing-codes-collection source path:
Enter it here > `
	_ := <- c

	softwareInputSource := bufio.NewReader (os.Stdin)
	input, _, errX := softwareInputSource.ReadLine ()
	if errX != nil {

