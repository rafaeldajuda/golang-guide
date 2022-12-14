# 1 - variáveis - valores - tipos

## tipos de declarações

Existe várias formas de se declarar variáveis e atrubiur valores a elas. Estes são alguns exemplos:

**declaracão e inicialização em linhas separadas**
```go
var a string
a = "olá a"
fmt.Println(a)
```

**declaracão com inicialização**
```go
var b string = "olá b"
fmt.Println(b)
```

**declaracão com inicialização sem especificar o tipo**
```go
var c = "olá c"
fmt.Println(c)
```
OBS: Nesse caso, a linguagem golang irá por si só identificar o tipo da variável de acordo com o valor atribuído.

**gopher - declaração e inicialização curta**
```go
d := "olá d"
fmt.Println(d)
```

## tipos primitivos

No golang existem os seguintes tipos primitivos:

- string 
- int (8, 16, 32, 64)
- float (32, 64)
- bool (true, false)

exemplos:
```go
texto := "string"
inteiro := 10
float := 10.1
boleano := true

fmt.Printf("%v tipo %T\n", texto, texto)
fmt.Printf("%v tipo %T\n", inteiro, inteiro)
fmt.Printf("%v tipo %T\n", float, float)
fmt.Printf("%v tipo %T\n", boleano, boleano)
```

## tipos compostos (composição)

São tipos criados pelo usuário.

- slice
- array
- struct
- map

## valos zero

São os valores inicias dos tipos.

- int: 0
- float: 0.0 
- boolean: false
- string: ""
- pointers, functions,  interfaces, slices, channels,  maps: nil

## keywords - pavavras-chaves

Estas são as palavras reservadas pelo golang. Não se deve usa-lás para declaração de variáveis.

```text
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```