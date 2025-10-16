package fizzbuzz

import (
	"fmt"
)

// Fizz gibt Zahlen von 1 bis 30 aus und ersetzt dabei jede durch 3 teilbare Zahl durch fizz, jede durch 5 teilbare Zahl durch buzz und jede durch 3 und 5 teilbare Zahl durch fizzbuzz
func FizzCompact() {
	// Schleife von 1 bis 30, die prüft ob die zahlen durch 3 un/oder 5 teilbar sind
	for i := 1; i < 31; i++ {
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
			continue
		}
		if i%3 == 0 {
			fmt.Print("fizz")
		}
		if i%5 == 0 {
			fmt.Print("buzz")
		}
		fmt.Println()

	}

}
