package gameboard

import "fmt"

func ExampleNewBoard() {
	fmt.Println(NewBoard())

	// Output:
	// [[     ] [     ] [     ]]
}

func ExamplePrintBoard() {
	PrintBoard(NewBoard())

	// Output:
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
}
