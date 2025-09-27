package funktionen

// SumOfFactorials berechnet die Summe der Fakultäten von 1 bis n.
func SumOfFactorials(n int) int {

	result := 0

	for i := 1; i <= n; i++ {
		result += Factorial(i)
	}

	return result
}

// Factorial berechnet die Fakultät von n.
func Factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
