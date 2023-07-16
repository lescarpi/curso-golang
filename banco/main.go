package main

import "fmt"

func main() {

	contaDoGuilherme := ContaCorrente{"Guilherme",
		589,
		123456,
		2000}

	contaDoJunior := ContaCorrente{"Junior",
		252,
		125156,
		1000}

	fmt.Println("Saldo do Guilherme:", contaDoGuilherme.saldo)
	fmt.Println("Saldo do Junior:", contaDoJunior.saldo)
	fmt.Println(contaDoGuilherme.Transferir(500, &contaDoJunior))
	fmt.Println("Saldo do Guilherme:", contaDoGuilherme.saldo)
	fmt.Println("Saldo do Junior:", contaDoJunior.saldo)

}
