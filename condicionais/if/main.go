package main

import "fmt"

func main() {
	meses := 6
	situacao := true
	cidade := "teste"

	// < > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("esse credor deve a menos de 6 mese")
	}

	if situacao {
		fmt.Println("Ele esta devendo")
	}
	if !situacao {
		fmt.Println("Ele esta adinplente")
	}

	if cidade == "teste" {
		fmt.Println("ele vive na cidade teste")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situação do cliente?", descricao)
		return
	}

	fmt.Println("Obrigado por nos consultar")
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "cliente esta devendo"
		return
	}

	descricao = "Esta em dia"
	return

}
