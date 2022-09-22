# deslocamento de bits

Deslocamentos de bits é mover x casas de bits para esquerda ou direita. Ao mover para esquerda o valor decimal do bit é multiplicado e ao mover para direita o valor decimal do bit é divido, ou seja, esquerda para multiplicar e direita para dividir.

Dependendo de quantas casas mover o valor multiplicado e dividido muda.

|deslocamento |multiplicado/dividido	|
|---		  |---						|
|1			  |2						|
|2			  |4						|
|3			  |8						|
|4			  |16						|

## deslocamento de 1 bit

```go
func main() {

	y := 12
	x := y >> 1
	z := y << 1

	fmt.Printf("y: %b -> %d\n", y, y) // 1100 -> 12
	fmt.Printf("x: %b -> %d\n", x, x) // 0110 -> 6
	fmt.Printf("z: %b -> %d\n", z, z) // 11000 -> 24

}
```

## deslocamento de 2 bits

```go
func main() {

	y := 12
	x := y >> 2
	z := y << 2

	fmt.Printf("y: %b -> %d\n", y, y) // 1100 -> 12
	fmt.Printf("x: %b -> %d\n", x, x) // 0011 -> 3
	fmt.Printf("z: %b -> %d\n", z, z) // 110000 -> 48

}
```