package main

import (
	"fmt"

	"github.com/ThiagoRodriguesdeSantana/curso_de_go/funcoes_basico/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)

	fmt.Printf("O resultado é: %d\r\n", resultado)

	resultado = matematica.Soma(3, 3)

	fmt.Printf("O resultado é: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Divisor, 12, 3)

	fmt.Printf("O resultado da divisao é: %d\r\n", resultado)

	resultado, resto := matematica.DivisorComResto(10, 3)
	fmt.Printf("O resultado da divisao é: %d e o resto da divisão foi: %d\r\n", resultado, resto)
}
