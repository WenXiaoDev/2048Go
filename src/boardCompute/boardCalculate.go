package boardCompute

import (
	"boardTrans"
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type zeroSite struct {
	x int
	y int
}

func Calculate(blocks [][]int, emptyPtr *int) [][]int {
	var result [][]int
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		// up left down right
		case termbox.KeyArrowUp:
			result, *emptyPtr = computeUpper(blocks)
		case termbox.KeyArrowLeft:
			boardTrans.Rotate(&blocks, true)
			result, *emptyPtr = computeUpper(blocks)
			boardTrans.Rotate(&result, false)
		case termbox.KeyArrowDown:
			boardTrans.FlipUpright(&blocks)
			result, *emptyPtr = computeUpper(blocks)
			boardTrans.FlipUpright(&result)
		case termbox.KeyArrowRight:
			boardTrans.Rotate(&blocks, false)
			result, *emptyPtr = computeUpper(blocks)
			boardTrans.Rotate(&result, true)
		default:
			result = blocks
		}
		return result
	default:
		return blocks
	}
}

func computeUpper(board [][]int) ([][]int, int) {
	var result [][]int
	// calculation of the specific swiping operation
	for i := 0; i < 4; i++ {
		column := make([]int, 4)
		var temp []int
		for j := 0; j < 4; j++ {
			if board[j][i] != 0 {
				temp = append(temp, board[j][i])
			}
		}
		copy(column, temp)
		result = append(result, combineColumn(column))
	}
	boardTrans.Transpose(&result)
	// generate one "2" in site where number is zero, randomly
	var list []zeroSite
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if result[i][j] == 0 {
				list = append(list, zeroSite{i, j})
			}
		}
	}
	restSpot := len(list)
	// check if the swiping is effective
	if restSpot != 0 && !boardEqual(&board, &result) {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(restSpot)
		result[list[index].x][list[index].y] = 2
		restSpot -= 1
	}
	return result, restSpot
}

func combineColumn(col []int) []int {
	/*
		all items in col have been put left
	*/
	if col[0] == 0 {
		return col
	}
	var stack []int
	for i := 0; i < len(col); i++ {
		if len(stack) == 0 || stack[len(stack)-1] != col[i] {
			stack = append(stack, col[i])
		} else {
			stack[len(stack)-1] *= 2
		}
	}
	result := make([]int, 4, 4)
	copy(result, stack)
	return result
}

func boardEqual(b1, b2 *[][]int) bool {
	for i := 0; i < len(*b1); i++ {
		for j := 0; j < len(*b2); j++ {
			if (*b1)[i][j] != (*b2)[i][j] {
				return false
			}
		}
	}
	return true
}

func Test() {
	bench := [][]int{
		[]int{4, 4, 2, 2},
		[]int{2, 2, 0, 0},
		[]int{4, 2, 0, 0},
		[]int{8, 4, 2, 2},
	}
	/* a simulation of swiping left on bench
	fmt.Printf("origin board\n")
	boardTrans.ShowBoard(&bench)

	fmt.Printf("bench rotated clockwisely\n")
	boardTrans.Rotate(&bench, true)
	boardTrans.ShowBoard(&bench)

	fmt.Printf("compute upper:\n")
	bench = computeUpper(bench)
	boardTrans.ShowBoard(&bench)

	fmt.Printf("bench rotate Counterclockwisely\n")
	boardTrans.Rotate(&bench, false)
	boardTrans.ShowBoard(&bench)
	*/
	fmt.Print(bench)
}
