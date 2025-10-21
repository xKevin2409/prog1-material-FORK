package listprops

// ContainsN liefert true, falls die Liste l
// den String x mindestens n mal enthält.
func ContainsN(l []string, x string, n int) bool {
	y := 0
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			y = y + 1
		}

	}
	if y >= n {
		return true
	}
	return false
}
