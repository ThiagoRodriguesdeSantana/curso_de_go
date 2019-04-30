package main

import "fmt"

func main() {
	/*
		O go run apenas gera um build temporario para que o programa possa ser executado
		ja o go build gere o build definitivo da aplicação

		>> go build

		Ao executar o comando go build o build fica com o nome padrão de "BUILD"
		o nome do build podera ser alterado usando o camando

		>> go build -o "nome do app"

		CROSS COMPILATION

		Com o Cross Compilation o go permite compular para outras plataformas

		exemplo de como gerar executavel para o windows
		>> GOOS=windows GOARCH=386 go build -o meuapp.exe

	*/

	fmt.Println("Testes de build no linux para windows")
}
