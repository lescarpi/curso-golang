package contas

import "github.com/lescarpi/curso-golang/banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	if valor <= 0 {
		return "Valor do saque precisa ser maior do que 0!"
	}
	if valor > c.saldo {
		return "Saldo insuficiente!"
	}
	c.saldo -= valor
	return "Saque realizado com sucesso!"
}

func (c *ContaCorrente) Depositar(valor float64) string {
	if valor <= 0 {
		return "Valor do depósito precisa ser maior do que 0!"
	}
	c.saldo += valor
	return "Depósito realizado com sucesso!"
}

func (c *ContaCorrente) Transferir(valor float64, r *ContaCorrente) string {
	if valor <= 0 {
		return "Valor da transferência precisa ser maior do que 0!"
	}
	if valor > c.saldo {
		return "Saldo insuficiente!"
	}
	c.Sacar(valor)
	r.Depositar(valor)
	return "Transferência realizada com sucesso!"
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
