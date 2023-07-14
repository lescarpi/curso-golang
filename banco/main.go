package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	if valor <= 0 {
		return "Valor do saque precisa ser maior do que 0!"
	}
	if valor <= c.saldo {
		c.saldo -= valor
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func main() {

	contaDoGuilherme := ContaCorrente{"Guilherme",
		589,
		123456,
		125.50}

	fmt.Println(contaDoGuilherme)

	fmt.Println(contaDoGuilherme.Sacar(25.50))

	fmt.Println(contaDoGuilherme)

}
