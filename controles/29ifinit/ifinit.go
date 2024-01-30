package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	fmt.Println("New Source", s)
	r := rand.New(s)
	fmt.Println("\n", *r)
	return r.Intn(100)
}

func main() {
	if i := numeroAleatorio(); i > 50 {
		fmt.Println("Ganhou!!!")
		fmt.Println(i)
	} else {
		fmt.Println("Perdeu!")
		fmt.Println(i)
	}
}
