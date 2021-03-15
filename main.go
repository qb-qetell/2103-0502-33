package main

import (
	"fmt"
	"github.com/inancgumus/screen"
)

func main () {
	c := make (chan [2]string)
	go taskThread (c)
	outputList := []string {"Code ID: 2103-0502-33 --- Code Name: Supreme potato"}
	
	for {
		message := <- c
		c <- [2]string {"oo", "ok"}
		
		if message [0] == "l2" {
			newOutputList := []string {outputList [0]}
			newOutputList = append (newOutputList, message [1])
			outputList = newOutputList
		}
		if message [0] == "l3" {
			newOutputList := []string {outputList [0], outputList [1]}
			newOutputList = append (newOutputList, message [1])
			outputList = newOutputList
		}
		if message [0] == "l4" {
			newOutputList := []string {outputList [0], outputList [1],
				outputList [2]}
			newOutputList = append (newOutputList, message [1])
			outputList = newOutputList
		}
		if message [0] == "hl" {
			break
		}
		
		screen.Clear ()
		screen.MoveTopLeft ()

		fmt.Println (outputList [0] + "\n")

		if len (outputList) == 2 || len (outputList) > 2 {
			fmt.Println (outputList [1] + "\n\n++++  ++++  ++++  ++++\n")
		}

		if len (outputList) == 3 || len (outputList) > 3 {
			fmt.Println (outputList [2] + "\n")
		}

		if len (outputList) == 4 || len (outputList) > 4 {
			fmt.Println (outputList [3] + "\n")
		}
	}
}
