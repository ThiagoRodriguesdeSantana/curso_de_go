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
	casa.SetValor(60000)

	fmt.Printf("casa é %v\r\n", casa)

	fmt.Printf("o valor da casa é %v\r\n", casa.GetValor())
	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em json:", string(objJSON))

}
