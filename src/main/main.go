package main

import (
	"fmt"
	"gamePlay"
)

const (
	BOARDSIZE = 4
)

func main() {
	fmt.Print("Game 2048 in Golang.\n")
	/*
		initBoard()
		boardRun()
			inputMonitor()
				verify input string
				verify game loss condition
			Calculate()


		for :
			inputMonitor()
			go lossCondition()
			Calculate()
			flag <- c
	*/
	gamePlay.GameRun()
}
