package aufgabe3

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion CountOdd.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein.
*/

// CountOdd erwartet eine Liste von Zahlen und liefert die Anzahl der ungeraden Zahlen darin.
func CountOdd(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	if nums[0]%2 == 1 {
		count = 1

	}
	return count + CountOdd(nums[1:])
}
