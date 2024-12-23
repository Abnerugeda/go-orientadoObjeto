package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaCorrente := ContaCorrente{titular: "Abner",
		numeroAgencia: 1234, numeroConta: 412412, saldo: 124124}

	fmt.Println(contaCorrente)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 21412

	fmt.Println(*contaDaCris)

}
