package listsearch

import "fmt"

// Sucht x in l und liefert die Position
// oder -1, falls x nicht in l ist.
func Find(l []int, x int) int {
	for i, y := range l {
		if y == x {
			return i
		}
	}
	//TODO
	return -1
}

// Liefert eine Liste mit allen Vorkommen von x in l
func FindBetter(l []int, x int) []int {
	result := []int{}

	for i := 0; i < len(l); i++ {
		if l[i] == x {
			result = append(result, i)
		}
	}

	return result
}

func ExampleFindBetter() {
	l1 := []int{17, 5, 42, 25, 3, -4, 8, -23, 5}

	pos1 := FindBetter(l1, 42)
	pos2 := FindBetter(l1, 200)
	pos3 := FindBetter(l1, 5)

	fmt.Println(pos1)
	fmt.Println(pos2)
	fmt.Println(pos3)
	// Output:
	// [2]
	// -1

}
