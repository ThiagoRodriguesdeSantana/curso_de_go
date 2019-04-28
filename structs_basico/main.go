package main

import "fmt"

//Imovel é uma struct
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := Imovel{}
	fmt.Printf("A casa é %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu AP", 760000}
	fmt.Printf("A apartamento é %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacara",
		valor: 55,
		X:     22,
	}

	fmt.Printf("A chacara é %+v\r\n", chacara)

	casa.Nome = "Lar doce lar"
	fmt.Printf("A casa é %+v\r\n", casa)
}
