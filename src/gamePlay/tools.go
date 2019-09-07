package gamePlay

import "C"
import (
	"boardCompute"
	"boardTrans"
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func StopJudgeSimple(blocks *[][]int) bool {
	board := *blocks
	flag := false
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
	emptyNum := 15
	for i := 0; i < 4; i++ {
		board = append(board, make([]int, 4))
	}
	rand.Seed(time.Now().UnixNano())
	board[rand.Intn(4)][rand.Intn(4)] = 2

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	for {
		boardTrans.ShowBoard(&board)
		board = boardCompute.Calculate(board, &emptyNum)
		if emptyNum == 0 && !StopJudgeSimple(&board) {
			boardTrans.ShowBoard(&board)
			fmt.Print("GAME OVER!")
			break
		}
		time.Sleep(100)
	}
}

func Test() {
	fmt.Println("12234455562574588")
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	GameRun()
}
