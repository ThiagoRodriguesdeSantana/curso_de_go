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
	casa := new(Imovel)

	fmt.Printf("casa é  %p - %v\r\n", &casa, casa)

	chacara := Imovel{
		X:     17,
		Y:     28,
		Nome:  "Chacara linda",
		valor: 280000,
	}

	fmt.Printf("chacara é  %p - %v\r\n", &chacara, chacara)

	mudaImovel(&chacara)

	fmt.Printf("chacara é  %p - %v\r\n", &chacara, chacara)
}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}
