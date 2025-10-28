package bubblesort

import "fmt"

func ExampleBubbleUp() {
	a := []int{5, 1, 8, 3, 4}

	fmt.Println(BubbleUp(a))
	fmt.Println(a)

	fmt.Println(BubbleUp(a))
	fmt.Println(a)

	fmt.Println(BubbleUp(a))
	fmt.Println(a)

	// Output:
	// true
	// [1 5 8 3 4]
	// true
	// [1 5 3 8 4]
	// true
	// [1 3 5 8 4]
}

func ExampleBubbleSort() {
	a := []int{5, 1, 8, 3, 4}
	BubbleSort(a)
	fmt.Println(a)

	// Output:
	// [1 3 4 5 8]
}
