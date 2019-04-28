package main

import (
	"fmt"

	"github.com/ThiagoRodriguesdeSantana/curso_de_go/struct_avancado/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)
	casa := model.Imovel{}

	casa.Nome = "casa amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	cache["casa amarela"] = casa

	casa2 := model.Imovel{}

	casa2.Nome = "casa verde"
	casa2.X = 18
	casa2.Y = 25
	casa2.SetValor(60000)

	cache["casa verde"] = casa2

	fmt.Println("ha uma casa amarela no cache")

	fmt.Printf("quantos itens ha no cache %+v\r\n", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("item[%s] = %+v\r\n", chave, imovel)
	}
	//remover item do cache

	imovel, achou := cache["casa amarela"]

	if achou {
		delete(cache, imovel.Nome)
	}

	fmt.Printf("quantos itens ha no cache %+v\r\n", len(cache))
}
