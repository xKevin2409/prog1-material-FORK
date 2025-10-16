package fizzbuzz

import (
	"fmt"
)

// Fizz gibt Zahlen von 1 bis n aus und ersetzt dabei
// jede durch x teilbare Zahl durch fizz, jede durch
// y teilbare Zahl durch buzz und jede durch x und y
// teilbare Zahl durch fizzbuzz
func FizzImproved(x, y, n int) {

	for i := 1; i < n; i++ {
		if i%x == 0 && i%y == 0 {
			fmt.Println("fizzbuzz")
		} else if i%x == 0 {
			fmt.Println("fizz")
		} else if i%y == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}

	}

}
