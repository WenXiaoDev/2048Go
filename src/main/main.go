/*
golang 实现无缓冲输入(getch())：
	1.用cgo调用内嵌c语言代码 需要将mingw32换成mingw64
	2.第三方库 termbox-go
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
package main

import (
	"fmt"
	"gamePlay"
)

func main() {
	fmt.Print("\tGame 2048 in Golang.\n")
	gamePlay.Test()
}
