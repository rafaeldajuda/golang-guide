# iota

Iota é um tipo inteiro não defindo que representa valores sequênciais. Este tipo deve ser utilizado em constantes.

## formas de declarar

```go
const (
	a = iota // -> 0
	b = iota // -> 1
	c = iota // -> 2
	x = iota // -> 3
	y = iota // -> 4
	z = iota // -> 5
)
```

```go
const (
	a = iota // -> 0
	_ // -> 1
	c // -> 2
	x // -> 3
	_ // -> 4
	z // -> 5
)
```