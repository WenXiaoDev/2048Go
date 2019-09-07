/*
	Swiping operation of all direction(left,right,up,down) be transferred to the
	same condition(up) temporarily, so that the computing task be constrained to
	the only condition.

	The matrix be rotated by calling function Rotate() and be flipped by calling
	function FlipUpright().

	There is a slower version of rotate, which is not exposed.
*/
package boardTrans

import (
	"fmt"
	"os"
	"os/exec"
)

func Rotate(blocks *[][]int, isClockwise bool) {
	board := *blocks
	size := len(board)
	if isClockwise {
		for i := 0; i < size/2; i++ {
			for j := i; j < size-i-1; j++ {
				board[i][j], board[j][size-i-1], board[size-i-1][size-j-1], board[size-j-1][i] =
					board[size-j-1][i], board[i][j], board[j][size-i-1], board[size-i-1][size-j-1]
			}
		}
	} else {
		for i := 0; i < size/2; i++ {
			for j := i; j < size-i-1; j++ {
				board[i][j], board[j][size-i-1], board[size-i-1][size-j-1], board[size-j-1][i] =
					board[j][size-i-1], board[size-i-1][size-j-1], board[size-j-1][i], board[i][j]
			}
		}
	}
}

func FlipUpright(blocks *[][]int) {
	board := *blocks
	size := len(board)
	for i := 0; i < size; i++ {
		for j := 0; j < size/2; j++ {
			board[j][i], board[size-j-1][i] = board[size-j-1][i], board[j][i]
		}
	}
}

func Transpose(blocks *[][]int) {
	board := *blocks
	for i := 0; i < len(board); i++ {
		for j := i + 1; j < len(board); j++ {
			board[i][j], board[j][i] = board[j][i], board[i][j]
		}
	}
}

func ShowBoard(b *[][]int) {
	// clean the console first
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	space := " "
	for _, l := range *b {
		fmt.Print("\t\t\t\t\t\t\t")
		for _, i := range l {
			// 1 47 30
			if i != 0 {
				fmt.Printf("%c[1;44;37m%5d\t%c[0m", 0x1B, i, 0x1B)
			} else {
				fmt.Printf("%c[1;44;37m%5s\t%c[0m", 0x1B, space, 0x1B)
			}
			fmt.Print(space)
		}
		fmt.Printf("\n\n")
	}
}

// the functions below is a slower version of board transforming operation
func clockwiseRotate(blocks *[][]int) {
	board := *blocks
	size := len(board)
	for i := 0; i < size/2; i++ {
		for j := i; j < size-i-1; j++ {
			board[i][j], board[j][size-i-1], board[size-i-1][size-j-1], board[size-j-1][i] =
				board[size-j-1][i], board[i][j], board[j][size-i-1], board[size-i-1][size-j-1]
		}
	}
}

func counterClockwiseRotate(blocks *[][]int) {
	flipHorizontal(blocks)
	clockwiseRotate(blocks)
	flipHorizontal(blocks)
}

func flipHorizontal(blocks *[][]int) {
	board := *blocks
	size := len(board)
	for i := 0; i < size; i++ {
		for j := 0; j < size/2; j++ {
			board[i][j], board[i][size-j-1] = board[i][size-j-1], board[i][j]
		}
	}
}

func Test() {
	d := 0
	size := 5
	var board [][]int
	for i := 0; i < size; i++ {
		board = append(board, make([]int, size))
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			board[i][j] = d
			d++
		}
	}

	fmt.Print("original board\n")
	ShowBoard(&board)
	/*
		rotate(&board, true)
		fmt.Print("clockwise rotate with parameter isClockwise: false\n")
		ShowBoard(&board)

		rotate(&board, true)
		fmt.Print("clockwise rotate again:\n")
		ShowBoard(&board)

		rotate(&board, false)
		fmt.Print("counterclockwise rotate with parameter isClockwise: true\n")
		ShowBoard(&board)
	*/
	FlipUpright(&board)
	ShowBoard(&board)
}
