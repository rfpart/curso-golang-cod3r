package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)
	fmt.Println("Subtração => ", a-b)
	fmt.Println("Divisão => ", a/b)
	fmt.Println("Multiplicação => ", a*b)
	fmt.Println("Módulo => ", a%b)

	//Bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 = 10 (AND)
	fmt.Println("Ou =>", a|b)  // 11 | 10 = 11 (OR)
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01 (OU exclusivo)

	c := 3.0
	d := 2.0

	// Outras operações usando math
	fmt.Println("Maior =>", math.Max(float64(a), float64(b))) // foi necessário converter as variáveis para float64
	fmt.Println("Menor =>", math.Min(c, d))                   // neste caso as variáveis já eram float 64
	fmt.Println("Exponenciação =>", math.Pow(c, d))           // neste caso as variáveis já eram float 64

}
