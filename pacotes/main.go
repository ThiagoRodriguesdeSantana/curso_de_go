package main

import (
	"fmt"

	"github.com/ThiagoRodriguesdeSantana/curso_de_go/pacotes/operadora"
	"github.com/ThiagoRodriguesdeSantana/curso_de_go/pacotes/prefixo"
)

//NomeDoUsuario nome do uauario do sistema
var NomeDoUsuario = "thiago"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da capital %d\r\n", prefixo.Capital)
	fmt.Printf("Nome operadoara %s\r\n", operadora.NomeDaOperadora)
	fmt.Printf("Nome operadoara %s\r\n", operadora.PrefixoDaCapitalOperadora)
}
