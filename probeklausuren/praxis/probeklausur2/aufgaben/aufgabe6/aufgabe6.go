package aufgabe6

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// DuplicateSinglets erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste, bei der alle Elemente,
// die in list nur einmal vorkommen, verdoppelt sind,
// also zwei Mal hintereinander stehen.
// Elemente, die schon in list mehrfach vorkommen, sollen wie sie sind
// ins Ergebnis übertragen werden.
func DuplicateSinglets(list []int) []int {
	result := []int{}
	counter := 0
	roffl := 0
	for i := 0; i < len(list); i++ {
		roffl = list[i]
		for j := 0; j < len(list); j++ {
			if roffl == list[j] {
				counter = counter + 1
			}

		}
		if counter < 2 {
			result = append(result, roffl)
			result = append(result, roffl)
		} else {
			result = append(result, roffl)

		}
		counter = 0
	}

	return result
}
