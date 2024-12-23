package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) float64 {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return c.saldo
	}

	return c.saldo
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) Transferir(valorTransferencia float64, contaDestino *ContaPoupanca) bool {
	if valorTransferencia < c.saldo {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}

	return false
}
