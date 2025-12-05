package aufgabe1

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// PrefixBelow10 erwartet eine Liste "list" von Zahlen und liefert
// die längste Teil-Liste, mit der "list" beginnt und die nur Zahlen < 10 enthält.
func PrefixBelow10(list []int) []int {
	if len(list) == 0 {
		var leereliste []int
		return leereliste
	}
	var newlist []int
	for i := 0; i < len(list) && list[i] < 10; i++ {
		if list[i] < 10 {
			newlist = append(newlist, list[i])
		}

	}
	return newlist
}
