package main

import "fmt"

const z int = 100 // possui tipo

const x = 10 // não possui tipo

var y = 10 // tipo inteiro

var c int
var d float64

func main() {

	c = x
	fmt.Printf("%v, %T\n", c, c)
	d = x
	fmt.Printf("%v, %T\n", d, d)

	fmt.Println("")

	c = y
	fmt.Printf("%v, %T\n", c, c)

	// d = y <-- gera erro, pois d é um float64 e y é um inteiro
	// fmt.Printf("%v, %T\n", d, d)

}
