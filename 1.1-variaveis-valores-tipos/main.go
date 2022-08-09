package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	x := 10

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", b)

	// OBS
	// b is different than x, the types are differnts
	// b = x --> error

	// conversion
	fmt.Printf("%v\n", x)
	b = 101
	x = int(b)
	fmt.Printf("%v\n", x)

}
