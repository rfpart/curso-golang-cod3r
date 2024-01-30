package main

import "fmt"

const (
	notaA = 7.3
	notaB = 5.7
	notaC = 4.7
	notaD = 3.9
)

func main() {
	verifyStudentPass(notaA, notaB, notaC, notaD)
}

func verifyStudentPass(a, b, c, d float64) {
	result := calc(a, b, c, d)
	fmt.Println(result)
}

func calc(a, b, c, d float64) float64 {
	value := a + b + c + d

	return value / 4
}

func teste_final(result float64) {
	
	
	if result <= 4.9 {
        fmt.Println("REPROVADO")
    } else if result >= 5.0 && result <= 7.0 {
        fmt.Println("RECUPERAÇÃO")
    } else {
        fmt.Println("APROVADO")
    }

}

//if user.result >= 7.0 passou
//if user.result >= 5.0 recuperação
//if user.result <= 4.9 reprovado
