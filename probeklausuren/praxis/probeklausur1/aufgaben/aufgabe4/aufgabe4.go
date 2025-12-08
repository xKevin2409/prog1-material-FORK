package aufgabe4

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion MaxElements.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein!
*/

// MaxElements erwartet zwei int-Listen und liefert eine Liste, die an jeder Position
// jeweils das größere der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt, gilt dieses Element als das größere.
func MaxElements(l1, l2 []int) []int {
	var newlist []int

	if len(l1) == len(l2) && len(l1) == 0 {
		return newlist
	}

	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}

	var x, y int
	if len(l1) < len(l2) {
		for i := 0; i < len(l2); i++ {
			x = l2[i]
			if i >= len(l1) {
				y = x
			} else {
				y = l1[i]
			}

			if y < x {
				newlist = append(newlist, x)

			} else {

				newlist = append(newlist, y)
			}

		}

	}

	var c, v int

	if len(l1) > len(l2) {
		for j := 0; j < len(l1); j++ {
			c = l1[j]
			if j >= len(l2) {
				v = c
			} else {
				v = l2[j]
			}

			if v < c {
				newlist = append(newlist, c)

			} else {

				newlist = append(newlist, v)
			}

		}

	}

	if len(l1) == len(l2) {
		for h := 0; h < len(l1); h++ {
			if l1[h] > l2[h] {
				newlist = append(newlist, l1[h])
			} else {
				newlist = append(newlist, l2[h])
			}
		}
	}

	return newlist
}
