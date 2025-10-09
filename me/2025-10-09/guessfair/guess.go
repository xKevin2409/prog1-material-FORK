package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//my_number := 42
	my_number := rand.Intn(100) + 1
	fmt.Println(my_number)

	for i := 0; i < 3; i++ {

		guess := ReadNumber()

		//if NumberGood(guess, my_number) {
		if guess == my_number {
			fmt.Println("Richtig geraten :-)")
			return

		}
		if guess < my_number {
			fmt.Println("Die gesuchte Zahl ist größer")
		} else {
			fmt.Println("Die gesuchte Zahl ist kleiner")
		}

	}
	fmt.Println("Game Over! Try it again!")
}

// ReadNumber liefert uns ein int
func ReadNumber() int {
	// var n int und nicht n := 0, da wir n noch keinen Wert geben möchten und das klar machen möchten
	var n int
	fmt.Print("Bitte eine Zahl eingeben: ")
	fmt.Scan(&n)
	return n

}
