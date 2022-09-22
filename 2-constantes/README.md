# constantes

Constantes são "variáveis" imutáveis. São utilizadas para valores que não irão mudar durante a execução da aplicação.

## tipos

As constantes pode ou não ter tipos.

possui tipo
```go
const z int = 100
```

não possui tipo
```go
const x = 10 
```

Constantes sem tipo possui um comportamento a parte, por exmplo, podemos atribuir uma constante com valor 10 para variáveis do tipo int (inteiro) e para float (ponto flutuante). O valor da constante só será definido na hora do seu uso, após isso a constante continuará sem tipo definido. Ou seja, podemos utilizar as constantes sem tipo definido de forma mais "generalizada".

```go
package main

import "fmt"

const x = 10 // não possui tipo

func main() {

	c = x
	fmt.Printf("%v, %T\n", c, c)
	d = x
	fmt.Printf("%v, %T\n", d, d)
}
```
saida
```cmd
10, int
10, float64
```
 
## formas de declarar

```go
const a string = "sou um a"
const b = "sou um b"

const (
    c = "sou um c"
    d = 4
    e = 4.5
    f = true
)
```

No exemplo abaixo é mostrado como declarar e iniciar constantes com os mesmos valores. A primeira constante precisa ter um valor definido e enquanto as seguintes, caso não seja passado nenhuma valor a elas, irão receber o mesmo valor da primeira. No exmplo abaixo, as constantes **b** e **c** possuem o mesmo valor da constante **a**, e a constante **e** possui o mesmo valor da constante **d**.

```go
const (
    a = 10
    b
    c 
    d = "sou uma letra"
    e
)
```