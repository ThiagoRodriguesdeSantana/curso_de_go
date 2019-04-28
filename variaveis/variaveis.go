package main

import "fmt"

// estrutura de variavel de pacote
var (
	endereco   string
	quantidade int
	comprou    bool
	valor      float64

	//vaariaveis rune (runas) são caracteres especiais
	palavra rune
)

//Variaveis começando em maiusculo sao publicas, e em menusculo sao privadas

func main() {
	fmt.Printf("endereço: %s\r\n", endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavra)
}
