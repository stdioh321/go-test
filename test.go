package main

import "fmt"

func main() {
	var (
		Endereco string = "Candelaria"
	)

	fmt.Println("Rua " + Endereco)
	fmt.Println(Calculo(Soma, 1, 3))
	fmt.Println(ResultadoComResto(11, 2))

}

func Calculo(funcao func(int, int) int, a int, b int) int {
	return funcao(a, b)
}

func Soma(a int, b int) int {
	return a + b
}

func ResultadoComResto(a int, b int) (result int, resto int) {
	result = a / b
	resto = a % b
	return
}
