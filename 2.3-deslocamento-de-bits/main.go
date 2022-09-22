package main

import "fmt"

func main() {

	y := 12
	x := y >> 2
	z := y << 2

	fmt.Printf("y: %b -> %d\n", y, y) // 1100 -> 12
	fmt.Printf("x: %b -> %d\n", x, x) // 0011 -> 3
	fmt.Printf("z: %b -> %d\n", z, z) // 110000 -> 48

}
