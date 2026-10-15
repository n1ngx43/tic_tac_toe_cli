package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func displayBoard(board *[3][3]string) {
	fmt.Println("")
	for i := 0; i < 3; i++ {
		fmt.Printf(" %s | %s | %s \n", board[i][0], board[i][1], board[i][2])
		if i < 2 {
			fmt.Println("---+---+---")
		}
	}
}

func inputPlayerX(position string, board *[3][3]string) error {
	positionInt, err := strconv.Atoi(position)
	var found int = -1

	if err != nil {
		return errors.New("position is not a number")
	}

	if positionInt < 1 || positionInt > 9 {
		return errors.New("position must be between 1 and 9")
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == position {
				board[i][j] = "X"
				found = 1
				break
			}
		}
	}

	if found == -1 {
		return errors.New("position is already marked")
	}

	return nil
}

func inputPlayerO(position string, board *[3][3]string) error {
	positionInt, err := strconv.Atoi(position)
	var found int = -1

	if err != nil {
		return errors.New("position is not a number")
	}

	if positionInt < 1 || positionInt > 9 {
		return errors.New("position must be between 1 and 9")
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == position {
				board[i][j] = "O"
				found = 1
				break
			}
		}
	}

	if found == -1 {
		return errors.New("position is already marked")
	}

	return nil
}

func inputMenu() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("")
	fmt.Println("-- TIC TAC TOE CLI ----------")
	fmt.Println("1. X player")
	fmt.Println("2. O player")
	fmt.Println("3. Reset this game")
	fmt.Println("0. Stop this game")
	fmt.Println("-----------------------------")
	fmt.Print(">. Input your option: ")
	option, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	option = strings.TrimSpace(option)
	return option, nil
}

func inputPosition() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">. Input your position: ")
	position, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	position = strings.TrimSpace(position)
	return position, nil
}

func controller(board *[3][3]string) {
	for {
		option, err := inputMenu()
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nSystem Error!\nError: %s\n", err)
			return
		}
		switch option {
		case "1":
			for {
				position, err := inputPosition()
				if err != nil {
					fmt.Fprintf(os.Stderr, "\nSystem Error!\nError: %s\n", err)
					return
				}
				err = inputPlayerX(position, board)
				if err != nil {
					fmt.Fprintf(os.Stderr, "\nPlease, Try Again!\nError: %s\n", err)
					continue
				}
				break
			}
			displayBoard(board)
		case "2":
			for {
				position, err := inputPosition()
				if err != nil {
					fmt.Fprintf(os.Stderr, "\nSystem Error!\nError: %s\n", err)
					return
				}
				err = inputPlayerO(position, board)
				if err != nil {
					fmt.Fprintf(os.Stderr, "\nPlease, Try Again!\nError: %s\n", err)
					continue
				}
				break
			}
			displayBoard(board)
		case "3":
		case "0":
		default:
			fmt.Printf("\nInvalid Option!\n")
		}
	}

}

func main() {

	board := [3][3]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	controller(&board)
}
