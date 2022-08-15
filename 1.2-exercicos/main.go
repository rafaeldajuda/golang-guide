package main

import "fmt"

// var x int = 42
// var y string = "James Bond"
// var z bool = true

type batata int

var x batata

var y int

func main() {

	// x := 42
	// y := "James Bond"
	// z := true

	// fmt.Printf("x: %d\ny: %s\nz: %t\n", x, y, z)
	// fmt.Println("x", x)
	// fmt.Println("y", y)
	// fmt.Println("z", z)

	//////////////////////

	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)

	/*
		golang atribui os valores padrões as variáves,
		esses valores são chamdados de "valor zero"
	*/

	//////////////////////

	// s := fmt.Sprintf("%d %s %t", x, y, z)
	// fmt.Println(s)

	//////////////////////

	fmt.Printf("%v, %T\n", x, x)
	x = 42
	fmt.Printf("%v\n", x)

	//////////////////////

	y = int(x)
	fmt.Printf("%v, %T\n", y, y)

}
