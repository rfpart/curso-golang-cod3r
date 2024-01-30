package main

import "fmt"

func main() {

	incrementa(0)
	decrementa(0)
	//main()
	fmt.Println("Fim")
	fmt.Println("Ricardo")

}

func incrementa(x int) {
	fmt.Println(x)

	if x >= 9 {

		return
	}

	incrementa(x + 1)

}
func decrementa(x int) {
	fmt.Println(x)
	if x <= -9 {
		return
	}
	decrementa(x - 1)
}
