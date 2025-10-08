package greet

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Herzlich Willkommen, %s!", name)
}
