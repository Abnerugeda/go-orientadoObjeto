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

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}

	return false
}

func main() {
	contaSilva := ContaCorrente{titular: "Silvia",
		numeroAgencia: 1234, numeroConta: 412412, saldo: 200}

	contaGustavo := ContaCorrente{titular: "Gustavo",
		numeroAgencia: 1234, numeroConta: 42412, saldo: 500}

	status := contaSilva.Transferir(100, &contaGustavo)

	fmt.Println(contaSilva)
	fmt.Println(contaGustavo)
	fmt.Println(status)

}
