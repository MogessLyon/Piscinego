package main

import (
	"fmt"
)

func view(array [][]int) {
	for idx, value := range array {
		for id, val := range value {
			if id != len(value)-1 {
				fmt.Print(val, " ")
			} else {
				fmt.Print(val)
			}

			if (id+1)%3 == 0 {
				if id != len(value)-1 {
					fmt.Print("| ")
				}
			}
		}
		fmt.Print("\n")
		if (idx+1)%3 == 0 {
			for i := 0; i < 9*2+3; i++ {
				if idx != len(value)-1 {
					fmt.Print("-")
				}
			}
			fmt.Print("\n")
		}
	}
}

func FindEmpty(board [][]int) (int, int) {
	for y := range board {
		for x := range board[y] {
			if board[y][x] == 0 {
				return x, y
			}
		}
	}
	return -1, -1
}

func Valid(board [][]int, num, posx, posy int) bool {
	// Check column
	for i := 0; i < len(board); i++ {
		if board[i][posx] == num {
			return false
		}
	}
	// Check row
	for i := 0; i < len(board[posy]); i++ {
		if board[posy][i] == num {
			return false
		}
	}
	box_xstart := (posx / 3) * 3
	box_ystart := (posy / 3) * 3
	//
	for j := box_ystart; j < box_ystart+3; j++ {
		for i := box_xstart; i < box_xstart+3; i++ {
			if board[j][i] == num {
				return false
			}
		}
	}

	return true
}

func Solve(board [][]int) bool {
	x, y := FindEmpty(board)
	if x == -1 || y == -1 {
		return true
	}

	for i := 1; i < 10; i++ {
		if Valid(board, i, x, y) {
			board[y][x] = i

			if Solve(board) {
				return true
			}

			board[y][x] = 0
		}
	}

	return false
}

func main() {
	board := [][]int{
		{7, 8, 0, 4, 0, 0, 1, 2, 0},
		{6, 0, 0, 0, 7, 5, 0, 0, 9},
		{0, 0, 0, 6, 0, 1, 0, 7, 8},
		{0, 0, 7, 0, 4, 0, 2, 6, 0},
		{0, 0, 1, 0, 5, 0, 9, 3, 0},
		{9, 0, 4, 0, 6, 0, 0, 0, 5},
		{0, 7, 0, 3, 0, 0, 0, 1, 2},
		{1, 2, 0, 0, 0, 7, 4, 0, 0},
		{0, 4, 9, 2, 0, 6, 0, 0, 7},
	}
	Solve(board)
	view(board)
}
