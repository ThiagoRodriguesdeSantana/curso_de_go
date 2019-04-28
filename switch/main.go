package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	numero := 3

	fmt.Print("o numero ", numero, " se escreve assim: ")

	switch numero {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("tres")
	}

	fmt.Println("vc e da familia do UNIX?")

	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsb":
		fallthrough
	case "linux":
		fmt.Println("Sim!!")
	default:
		fmt.Println("Deixa pra la")

	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Pode dormir ate tarde")
	default:
		fmt.Println("nao pode dormir ate tarde")
	}

	numero = 1

	fmt.Println("Esse numero cabe num digito?")

	switch {
	case numero < 10:
		fmt.Println("sim")
	case numero >= 10 && numero < 100:
		fmt.Println("serve dois digitos")
	default:

		fmt.Println("não cabe")
	}
}
