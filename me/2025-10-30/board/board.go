package gameboard

func NewBoard() [][]string {
	board := [][]string{}

	for row := 0; row < 3; row++ {
		board = append(board, []string{"   ", "   ", "   "})
	}
	return board
}


func PrintBoard(b [][]string) {
	div := 
}