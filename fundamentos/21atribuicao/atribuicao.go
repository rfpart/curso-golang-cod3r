package main

import "fmt"

func main() {

	var b byte = 3
	fmt.Println(b)

	i := 3 // inferência de tipo
	i += 3 // i = i + 3 => soma
	i -= 3 // i = i - 3 => subtração
	i /= 2 // i = i / 2 => divisão
	i *= 2 // i = i * 2 => multiplicação
	i %= 2 // i = i % 2 => módulo, quando não tem resto o resultado é 0 (divide i por 2)
	
	fmt.Println(i)

	x, y := 1, 2  // inicializa duas variáveis
	fmt.Println(x, y) 

	x, y = y, x  // troca o valor das variáveis
	fmt.Println(x, y)

}