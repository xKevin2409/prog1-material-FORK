package aufgabe3

import "math"

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * RANDBEDINGUNG: Die Funktion muss rekursiv sein.
 * ERREICHBARE PUNKTE: 10
 */

// CountSquares erwartet eine Liste von Zahlen.
// CountSquares liefert die Anzahl der QuadratzahlenZahlen in der Liste.
func CountSquares(list []int) int {
	if len(list) == 0 {
		return 0
	}

	// Prüfe das erste Element
	count := 0
	n := list[0]
	if n >= 0 { // nur nicht-negative Zahlen können Quadratzahlen sein
		r := int(math.Sqrt(float64(n)))
		if r*r == n {
			count = 1
		}
	}

	// Rekursiver Aufruf mit dem Rest der Liste
	return count + CountSquares(list[1:])
}
