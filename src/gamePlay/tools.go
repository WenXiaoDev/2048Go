package gamePlay

import (
	"boardCompute"
	"boardTrans"
	"fmt"
	"math/rand"
	"time"
)

func StopJudgeSimple(blocks *[][]int) bool {
	board := *blocks
	flag := false
	/*
		for i := 0; i < len(board); i++ {
			for j := 0; j <len(board); j++ {
				if board[i][j] == 0 {
					flag = true
					return flag
				}
			}
		}
	*/
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board)-1; j++ {
			if board[i][j] == board[i][j+1] {
				flag = true
				return flag
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board)-1; j++ {
			if board[j][i] == board[j+1][i] {
				flag = true
				return flag
			}
		}
	}
	return flag
}

func GameRun() {
	// init new board
	var board [][]int
	var emptyNum int
	for i := 0; i < 4; i++ {
		board = append(board, make([]int, 4))
	}
	rand.Seed(time.Now().UnixNano())
	board[rand.Intn(4)][rand.Intn(4)] = 2
	// main iteration
	var input string
	for {
		boardTrans.ShowBoard(&board)
		fmt.Scan(&input)
		fmt.Printf("input %s\n", input)
		board = boardCompute.Calculate(board, input, &emptyNum)
		if emptyNum == 0 && !StopJudgeSimple(&board) {
			boardTrans.ShowBoard(&board)
			fmt.Print("GAME OVER!")
			break
		}
	}
}

/*
golang 实现无缓冲输入(getch())：
	1.用cgo调用内嵌c语言代码 需要将mingw32换成mingw64
	2.第三方库 termbox-go
*/
func Test() {
	GameRun()
}
