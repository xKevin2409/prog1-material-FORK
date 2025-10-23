package main

import (
	"fmt"
)

var board [3][3]string
var currentPlayer string = "X"

func main() {
	initBoard()
	for {
		printBoard()
		row, col := getPlayerMove()
		if board[row][col] != " " {
			fmt.Println("Feld bereits belegt. Versuch's nochmal.")
			continue
		}
		board[row][col] = currentPlayer
		if checkWin(currentPlayer) {
			printBoard()
			fmt.Printf("Spieler %s gewinnt!\n", currentPlayer)
			break
		}
		if checkDraw() {
			printBoard()
			fmt.Println("Unentschieden!")
			break
		}
		switchPlayer()
	}
}

func initBoard() {
	for i := range board {
		for j := range board[i] {
			board[i][j] = " "
		}
	}
}

func printBoard() {
	fmt.Println("  0 1 2")
	for i, row := range board {
		fmt.Printf("%d %s|%s|%s\n", i, row[0], row[1], row[2])
		if i < 2 {
			fmt.Println("  -+-+-")
		}
	}
}

func getPlayerMove() (int, int) {
	var row, col int
	fmt.Printf("Spieler %s, gib deine Position ein (Zeile Spalte): ", currentPlayer)
	fmt.Scanf("%d %d\n", &row, &col)
	return row, col
}

func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}

func checkWin(player string) bool {
	// Zeilen, Spalten, Diagonalen prüfen
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}

func checkDraw() bool {
	for _, row := range board {
		for _, cell := range row {
			if cell == " " {
				return false
			}
		}
	}
	return true
}
