package main

import (
	"errors"
	"fmt"
	"strconv"
)

func displayBoard(board [3][3]string) {
	for i := 0; i < 3; i++ {
		fmt.Printf(" %s | %s | %s \n", board[i][0], board[i][1], board[i][2])
		if i < 2 {
			fmt.Println("---+---+---")
		}
	}
}

func inputPlayerX(position string, boardDrawed *[3][3]string) error {
	positionInt, err := strconv.Atoi(position)
	var found int = -1

	if err != nil {
		return errors.New("position is not a number")
	}

	if positionInt < 0 || positionInt > 9 {
		return errors.New("position is out of scope")
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if boardDrawed[i][j] == position {
				boardDrawed[i][j] = "X"
				found = 1
				break
			}
		}
	}

	if found == -1 {
		return errors.New("position has marked")
	}

	return nil
}

func inputPlayerO(position string, boardDrawed *[3][3]string) error {
	positionInt, err := strconv.Atoi(position)
	var found int = -1

	if err != nil {
		return errors.New("position is not a number")
	}

	if positionInt < 0 || positionInt > 9 {
		return errors.New("position is out of scope")
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if boardDrawed[i][j] == position {
				boardDrawed[i][j] = "O"
				found = 1
				break
			}
		}
	}

	if found == -1 {
		return errors.New("position has marked")
	}

	return nil
}

func main() {
	board := [3][3]string{
		{"1", "X", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	
	err := inputPlayerO("3", &board)
	if err != nil {
		fmt.Println(err)
	}
	displayBoard(board)
}
