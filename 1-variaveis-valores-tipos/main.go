package main

import "fmt"

func main() {

	// tipos de declaracoes

	// declaracão
	var a string
	a = "olá a"
	fmt.Println(a)

	// declaracao e atribuicao
	var b string = "olá b"
	fmt.Println(b)

	// declaracao e atribuicao sem especificar o tipo
	var c = "olá c"
	fmt.Println(c)

	// gopher - declaracao e atribuicao curta
	d := "olá d"
	fmt.Println(d)

	// tipos primitivos
	fmt.Println()

	texto := "string"
	inteiro := 10
	float := 10.1
	boleano := true

	fmt.Printf("%v tipo %T\n", texto, texto)
	fmt.Printf("%v tipo %T\n", inteiro, inteiro)
	fmt.Printf("%v tipo %T\n", float, float)
	fmt.Printf("%v tipo %T\n", boleano, boleano)

}
