package aufgabe4

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// ElementSums erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils die Summe der beiden Elemente enthält.
//
// Annahmen für die Berechnung:
// Falls eine Liste kürzer ist als die andere, soll für die Berechnung der
// hinteren Werte ihr letztes Element verwendet werden.
// Für leere Listen soll für die Berechnung ggf. 0 verwendet werden.
func ElementSums(l1, l2 []int) []int {
	result := []int{}

	if len(l1) == 0 && len(l2) == 0 {
		result = append(result, 0)
		return result
	}

	if len(l1) == 0 {
		for i := 0; i < len(l2); i++ {
			result = append(result, l2[i])

		}
		return result

	}

	if len(l2) == 0 {
		for h := 0; h < len(l1); h++ {
			result = append(result, l1[h])

		}
		return result

	}

	if len(l1) == len(l2) {
		for f := 0; f < len(l1); f++ {
			result = append(result, l2[f]+l1[f])
		}
		return result

	}

	if len(l1) != len(l2) {
		var x1, x2 int
		if len(l1) < len(l2) {
			for j := 0; j < len(l2); j++ {
				x1 = l2[j]
				if j >= len(l1) {
					x2 = l1[len(l1)-1]
				} else {
					x2 = l1[j]
				}
				result = append(result, x1+x2)
			}

			return result

		} else {

			for d := 0; d < len(l1); d++ {
				x1 = l1[d]
				if d >= len(l2) {
					x2 = l2[len(l2)-1]
				} else {
					x2 = l2[d]
				}
				result = append(result, x1+x2)
			}
			return result
		}
	}

	return result
}
