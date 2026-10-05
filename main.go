package main

import (
	"fmt"
)

func drawTable(board [3][3]string) {
	for i := 0; i < 3; i++ {
		fmt.Printf(" %s | %s | %s \n", board[i][0], board[i][1], board[i][2])
		if i < 2 {
			fmt.Println("---+---+---")
		}
	}
}

func main() {
	board := [3][3]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	drawTable(board)
}
