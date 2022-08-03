# 0 - hello world

Printa na tela um "Hello World!".

# Funções

## fmt.Println

Nessa aplicação foi utilizada a função **fmt.Println**. É a função básica para printar mesagens no terminal. 

Está função pertence ao pacote **fmt**. Nela é possível passar quantos argumentos você quer e de qualquer tipo.

Ao chamar retorna o número de bytes da mensagem printada e o algum erro, caso tenha acontecido algum.

## Exemplo

```go
package main

import "fmt"

func main() {

	fmt.Println("Hello World!")

}
```

## Links

- [fmt](https://pkg.go.dev/fmt)
- [fmt.Println](https://pkg.go.dev/fmt#Println)