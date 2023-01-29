package main

import (
	"fmt"
)

func main() {

	var nome string = "Douglas"
	var idade int = 32
	idade2 := 23
	var versao float32 = 1.1
	fmt.Println("Ola sr. ", nome)
	fmt.Println("A sua idade eh", idade)
	fmt.Println("O programa esta na versao", versao)
	fmt.Println("Hello world com go: ", idade2)
	fmt.Println("Conhecendo mais da linguagem")

	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O endereco da variavel comando eh,", &comando)
	fmt.Println("O comando escolhido foi: ", comando)

	var comando2 int
	fmt.Scan(&comando2)
	fmt.Println("o comando novo foi", comando2)

}
