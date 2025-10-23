package searching

func BinFindRek(l []int, x int) int {
	if len(l) == 0 {
		return -1
	}

	mitte := len(l) / 2
	if x == l[mitte] {
		return mitte
	}
	if x < l[mitte] {
		//Suche in der linken Hälfte weiter
		return BinFindRek(l[0:mitte], x)
	}
	if x > l[mitte] {
		//Suche in der rechten Hälfte weiter
		result := BinFindRek(l[mitte+1:], x)
		if result == -1 {
			return -1
		}
		return result + mitte + 1
	}
	return -1
}
