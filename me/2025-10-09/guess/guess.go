package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//my_number := 42
	my_number := rand.Intn(100) + 1

	for i := 0; i < 3; i++ {

		guess := ReadNumber()

		if NumberGood(guess, my_number) {
			//if guess == my_number {
			fmt.Println("Richtig geraten :-)")
			return

		}
		fmt.Printf("%d war nicht korrekt! \n", guess)

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

// NumberGood prüft, ob x gleich einer zufällig
// gewählten Zahl zwischen 1 und 100 ist
// Liefert true, wenn x gleich der Zahl ist, sonst false
func NumberGood(x int, n int) bool {

	return x == n

}
