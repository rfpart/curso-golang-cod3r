package main

import (
	"fmt"
	"sync"
)

var i int64 = 0
var once sync.Once

func main() {
	e := 1
	for e <= 10 { // não tem while em go
		fmt.Println(e)
		e++
	} //

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("O valor é: %d ", i)
	}

	fmt.Println("\n\nMisturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	for {

		// laço infinito
		fmt.Println("Para sempre..")
		//time.Sleep(time.Nanosecond * 5)
		if i == 100 {
			break
		}
		i++
	}

	fmt.Println("For Infinito Finalizado")
}
