package main

import "fmt"

func main() {
	fmt.Println("Ola mundo!")
	const constant = "Essa e uma constante"

	// iota pode ser usado para imcrementar numeros, comecando do 0
	const (
		_ = iota
		a
		b
		c = 1 << iota
		d
	)

	fmt.Println(a, b) // 1 2 (0 foi pulado)
	fmt.Println(c, d) // 8 16 (2^3, 2^4)
}

// Operadores
// Os padroes de linguagem
// ^ xor
// &^ e nao
// << desvio a esqurda
// >> desvio a direita

// Outros
// & Endereço de / criar ponteiro
// * ponteiro de desreferencia
// <- enviar / receber operador

// Declaração
// ----------------------------------------------------------------------------------//
// var foo int 					// declaração sem inicializacao
// var foo int = 42 				// declaração com inicializacao
// var foo, bar int = 42, 1302		// declara e inicia variaveis multiplas
// var foo = 42					// tipo omitido, sera inferido
// foo := 							// abreviamente, apenas em corpos de funcoes omite a palavra chave var, o tipo esta sempre implicito
// const constant = "Essa e uma constante"

// // iota pode ser usado para imcrementar numeros, comecando do 0
// const (
// 	_ = iota
// 	a
// 	b
// 	c = 1 << iota
// 	d
// )

// 	fmt.Println(a,b) // 1 2 (0 foi pulado)
// 	fmt.Println(c,d) // 8 16 (2^3, 2^4)
