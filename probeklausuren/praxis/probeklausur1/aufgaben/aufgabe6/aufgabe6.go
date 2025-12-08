package aufgabe6

import "slices"

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion SymmetricDifference.
MAX. PUNKTE: 10
*/

// SymmetricDifference erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die in einer, aber nicht in beiden Listen vorhanden sind.
//
// Die Elemente kommen dabei in der gleichen Reihenfolge vor, wie in den
// Ursprungslisten, mehrfach vorkommende Elemente werden entsprechend wiederholt.
// Die Elemente aus l1 kommen vor denen aus l2 in der Ergebnisliste vor.
func SymmetricDifference(l1, l2 []int) []int {
	var current int
	var newlist []int

	for i := 0; i < len(l1); i++ {
		current = l1[i]
		if !slices.Contains(l2, current) {
			newlist = append(newlist, current)

		}
	}
	for i := 0; i < len(l2); i++ {
		current = l2[i]
		if !slices.Contains(l1, current) {
			newlist = append(newlist, current)

		}
	}

	return newlist
}
