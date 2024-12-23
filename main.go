package main

import (
	"banco/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) float64
}

func main() {
	contaBruno := contas.ContaPoupanca{}
	contaBruno.Depositar(200)

	PagarBoleto(&contaBruno, 200)
}
