package datatypes

import "fmt"

func Example() {
	var i99 int
	k := 55

	fmt.Println(i99)
	fmt.Println(k)

	// Output:
	// 0
	// 55
}

func Example_various_data_types() {
	var i1 int
	i2 := 55
	var f1 float64
	f2 := 0.2

	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(f1)
	fmt.Println(f2)

	var b1 bool
	b2 := true

	fmt.Println(b1)
	fmt.Println(b2)

	var s1 string
	s2 := "Hallo Welt!"
	// s3 fängt bei 0 an zu zählen, 4 ist also das 5. Zeichen in Hallo Welt!
	s3 := s2[4]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("%c\n", s3)

	// Output:
	// 0
	// 55
	// 0
	// 0.2
	// false
	// true
	//
	//
	// Hallo Welt!
	// o

}
