package varnames

import "fmt"

func Example() {
	x := 42
	y := 23

	for x := 0; x < 5; x++ {
		fmt.Println(x)
	}
	fmt.Println(y)
	fmt.Println(x)
}

// Output:
// 0
// 1
