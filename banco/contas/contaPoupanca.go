package contas

import "github.com/lescarpi/curso-golang/banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valor float64) string {
	if valor <= 0 {
		return "Valor do saque precisa ser maior do que 0!"
	}
	if valor > c.saldo {
		return "Saldo insuficiente!"
	}
	c.saldo -= valor
	return "Saque realizado com sucesso!"
}

func (c *ContaPoupanca) Depositar(valor float64) string {
	if valor <= 0 {
		return "Valor do depósito precisa ser maior do que 0!"
	}
	c.saldo += valor
	return "Depósito realizado com sucesso!"
}
