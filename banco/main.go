package main

import (
	"fmt"

	"github.com/lescarpi/curso-golang/banco/clientes"
	"github.com/lescarpi/curso-golang/banco/contas"
)

func main() {

	contaDoGuilherme := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Guilherme",
			CPF:       "512.834.923-45",
			Profissao: "Pintor"},
		NumeroAgencia: 589,
		NumeroConta:   123456}

	contaDoGuilherme.Depositar(1000)

	contaDoDenis := contas.ContaPoupanca{}

	fmt.Println(contaDoDenis)

}
