package main

import "fmt"

func main() {

	version := 1.1
	fmt.Println("Olá, este programa está na versão:", version)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

}
