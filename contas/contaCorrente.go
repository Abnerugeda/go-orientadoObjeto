package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) float64 {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return c.saldo
	}

	return c.saldo
}

func (c *ContaCorrente) Depositar(valorDeposito float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}

	return false
}
