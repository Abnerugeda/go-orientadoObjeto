package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) (string, float64) {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso", c.saldo
	}

	return "Saldo insuficiente", c.saldo
}

func (c *ContaCorrente) Depositar(valorDeposito float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
	}
}

func main() {
	contaCorrente := ContaCorrente{titular: "Abner",
		numeroAgencia: 1234, numeroConta: 412412, saldo: 200}

	fmt.Println(contaCorrente.saldo)

	msg, saldo := contaCorrente.Sacar(200)

	fmt.Println(msg, saldo)

	contaCorrente.Depositar(1000)
	fmt.Println(contaCorrente)

}
