package main

import (
	"fmt"
	"prog1-material/me/2025-10-07/greet"
)

func main() {
	fmt.Println(greet.Greet("Kevin"))
	fmt.Println("sowie, " + greet.Greet("Kursteilnehmer"))
}
