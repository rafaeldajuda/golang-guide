# 1.1 - variaveis - valores - tipos

## criar tipos

Para criar tipos é preciso utilizar a palavra chave *type*, passar um nome e indicar o tipo base.

tipo hotdog (int)
```go
type hotdog int
var b hotdog
```

Nesse exemplo a variavel *b* é do tipo *hotdog*. Apesar do tipo hotdog ter como base o tipo *int* a variável b não um inteiro e sim um hotdog.

Para ver o tipo da variável b:

```go
package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	fmt.Printf("%T\n", b)
}
```

saida
```cmd
--> main.hotdog
```

Exemplo com uma variável do tipo int:

```go
package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	x := 10

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", b)
}
```

saida
```cmd
--> int
--> main.hotdog
```

**Observação**

Como foi dito, a variável b não um inteiro, então por isso, não é possível atribuir a variável x ao b sem fazer uma conversão antes.

não é possível, gera erro.
```go
package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	x := 10

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", b)

    b = x // nessa linha o go ira apontar um erro
}
```

## conversion

Não é possível atribuir o a variável x ao b, pois são de tipos diferentes, porém conseguimos atribuir a variável b ao x, mas é preciso fazer uma conversão antes.

conversão
```go
package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	x := 10

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", b)

	// conversion
	fmt.Printf("%v\n", x)
	b = 101
	x = int(b) // conversion
	fmt.Printf("%v\n", x)
}
```

saida
```cmd
int
main.hotdog
10
101
```

Como o subtipo do hotdog é int, podemos converter a variável b para int e atribuir o seu valor para o x. Só podemos fazer conversão para tipos que possuem subtipos (**tipos subjacentes**), além disso o tipo convertido precisa ter o seu subtipo compatível com o tipo da variável que irá receber o valor.