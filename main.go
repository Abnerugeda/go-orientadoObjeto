package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func main() {
	contaCorrente := ContaCorrente{titular: "Abner",
		numeroAgencia: 1234, numeroConta: 412412, saldo: 200}

	fmt.Println(contaCorrente.saldo)

	var msg = contaCorrente.Sacar(200)

	fmt.Println(msg)
}
