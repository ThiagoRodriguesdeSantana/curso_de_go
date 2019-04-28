package main

import (
	"encoding/json"
	"fmt"

	"github.com/ThiagoRodriguesdeSantana/curso_de_go/struct_avancado/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa amarela"
	casa.X = 18
	casa.Y = 25
	if err := casa.SetValor(10000000); err != nil {
		if err == model.ErrValorMuitoAlto {
			fmt.Println("corrija a avaliação ")
		}
		fmt.Println("[main] houve erro ", err.Error())
		return
	}

	fmt.Printf("casa é %v\r\n", casa)

	fmt.Printf("o valor da casa é %v\r\n", casa.GetValor())
	objJSON, err := json.Marshal(casa)

	if err != nil {
		fmt.Println("[main] houve erro na geração do objeto json", err.Error())
		return
	}
	fmt.Println("A casa em json:", string(objJSON))
}
