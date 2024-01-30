package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2    // (operação E) (ambos verdadeiro = verdadeiro)
	comprarTv32 := trab1 != trab2    // != substitue ou exclusivo  (dois iguais é falso))
	comprarSorvete := trab1 || trab2 // || (operação OU) (um ou outro é Verdadeiro)

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("TV50: %t, Tv32: %t, Sorvete: %t, Saudável: %t",
		tv50, tv32, sorvete, !sorvete)
}
