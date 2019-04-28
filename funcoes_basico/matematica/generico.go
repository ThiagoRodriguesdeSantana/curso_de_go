package matematica

// Calculo executa qualquer calculo
func Calculo(funcao func(int, int) int, nuemroA int, numeroB int) int {

	return funcao(nuemroA, numeroB)

}

//Multiplicador multiplica numeros
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor = divide numeros
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//DivisorComResto retorna tbm o resto da divisao
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
