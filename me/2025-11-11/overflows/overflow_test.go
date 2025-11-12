package overflow

import "fmt"

func Example_overflow() {
	x1 := 999999999999999999

	fmt.Println(x1*3333333333333333333 ^ 999999999999999999)
	//Output:
}
