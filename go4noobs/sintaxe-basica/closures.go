package main

import "fmt"

func main() {
	// atribuindo uma funcao a um nome
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(3, 4))
}

//Closures, com escopo lexico: as funcoers podem acessar valores que foram declarados no escopo ao definir a funcao
func scope() func() int {
	outerVar := 2
	foo := func() int { return outerVar }
	return foo
}

// func anotherScope() func() int {
// 	// nao ocmpilara porque outerVar e foo nao estao definidos neste escopo
// 	outerVar = 444
// 	return foo
// }

func outer() (func() int, int) {
	outerVar := 2
	inner := func() int {
		outerVar += 99 // outer var do escopo externo e mutada
		return outerVar
	}
	inner()
	return inner, outerVar // retorno da funcao inner e variavel mutada outer var 101
}

// funcoes variadicas

// Usando ... antes do nome do tipo do último parâmetro,
// você pode indicar que esse parâmetro leva zero ou mais desses parâmetros.
// A função é invocada como qualquer outra função,
// exceto que podemos passar quantos argumentos quisermos.
func adder(args ...int) int {
	total := 0
	for _, v := range args { // Repete os argumentos, seja qual for o número.
		total += v
	}
	return total
}

func main() {
	fmt.Println(adder(1, 2, 3)) // 6
	fmt.Println(adder(9, 9))    // 18

	nums := []int{10, 20, 30}
	fmt.Println(adder(nums...)) // 60
}
