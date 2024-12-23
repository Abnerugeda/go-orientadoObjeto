package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaSilva := contas.ContaCorrente{Titular: "Silvia",
		NumeroAgencia: 1234, NumeroConta: 412412, Saldo: 200}

	contaGustavo := contas.ContaCorrente{Titular: "Gustavo",
		NumeroAgencia: 1234, NumeroConta: 42412, Saldo: 500}

	status := contaSilva.Transferir(100, &contaGustavo)

	fmt.Println(contaSilva)
	fmt.Println(contaGustavo)
	fmt.Println(status)

}
