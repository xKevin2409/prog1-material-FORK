package searching

func BinFind1(l []int, x int) int {
	links := 0
	for len(l) > 0 {

		//Vergleiche x mit dem Leement in der Mitte
		//Wenn x == l[mitte], dann fertig
		mitte := len(l) / 2
		if l[mitte] == x {
			return mitte + links
		}

		//Wenn x < l[mitte], dann suche in der linken Hälfte weiter
		if x < l[mitte] {
			//Lasse nur den linken Teil der Liste übrig
			l = l[0:mitte]

		} else {
			//Wenn x > l[mitte], dann suche in der rechten Hälfte weiter
			//Lasse nur den rechten Teil der Liste übrig
			// l=l[mitte+1:len(l)]
			l = l[mitte+1:]
			links += mitte + 1

		}

	}
	return -1
}

func BinFind2(l []int, x int) int {
	links := 0
	rechts := len(l)

	for rechts > links {
		mitte := (rechts-links)/2 + links

		if l[mitte] == x {
			return mitte
		}
		if x < l[mitte] {
			rechts = mitte
		} else {
			links = mitte + 1
		}
	}
	return -1
}
