package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	var newlist []string
	posfirst := -1
	for i := 0; i < len(list); i++ {

		if list[i] == first {
			posfirst = i
		}

	}
	poslast := -1
	for j := 0; j < len(list); j++ {

		if list[j] == last {
			poslast = j
		}

	}
	if posfirst > poslast {
		return []string{}
	} else if posfirst == -1 && poslast == -1 {
		return []string{}
	} else {
		for l := 0; l < posfirst; l++ {
			newlist = append(newlist, list[l])

		}
		for h := poslast + 1; h < len(list); h++ {
			newlist = append(newlist, list[h])
		}

	}

	return newlist
}
