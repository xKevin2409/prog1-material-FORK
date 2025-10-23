package searching

import "fmt"

func ExampleBinFind1() {
	l1 := []int{1, 6, 8, 10, 25, 42, 125, 277}

	fmt.Println(BinFind1(l1, 8))
	fmt.Println(BinFind1(l1, 125))
	fmt.Println(BinFind1(l1, 42))
	fmt.Println(BinFind1(l1, 9))
	// Output:
	// 2
	// 6
	// 5
	// -1
	// 4

}

func ExampleBinFind2() {
	l1 := []int{1, 6, 8, 10, 25, 42, 125, 277}

	fmt.Println(BinFind2(l1, 1))
	fmt.Println(BinFind2(l1, 277))
	fmt.Println(BinFind2(l1, 25))
	fmt.Println(BinFind2(l1, 10))
	fmt.Println(BinFind2(l1, 3))
	fmt.Println(BinFind2(l1, 42))
	// Output:
	// 0
	// 7
	// 4
	// 3
	// -1
	// 5

}

func ExampleBinFindRek() {
	l1 := []int{1, 6, 8, 10, 25, 42, 125, 277}

	fmt.Println(BinFindRek(l1, 1))
	fmt.Println(BinFindRek(l1, 277))
	fmt.Println(BinFindRek(l1, 25))
	fmt.Println(BinFindRek(l1, 10))
	fmt.Println(BinFindRek(l1, 3))
	fmt.Println(BinFindRek(l1, 42))
	// Output:
	// 0
	// 7
	// 4
	// 3
	// -1
	// 5

}
