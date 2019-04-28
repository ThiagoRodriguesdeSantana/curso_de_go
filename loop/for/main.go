package main

import "fmt"

func main() {

	//for como for
	for i := 0; i < 10; i++ {
		println("qual o valor de i?", i)
	}

	valor := 0
	teste := true

	//for como while
	for teste {
		valor++
		if valor%5 == 0 {
			teste = false
		}

		fmt.Println("valor", valor)
	}

	for {
		valor--
		fmt.Println("valor", valor)

		if valor == 0 {
			break
		}
	}

	texto := "Eu adoro esquever programas em go"

	for indice, letra := range texto {
		fmt.Printf("texto[%d] = %q\r\n", indice, letra)
	}
}
