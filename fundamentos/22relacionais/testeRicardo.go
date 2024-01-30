package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Ricardo" == "Ricardo")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("\nDatas:", d1 == d2)
	fmt.Println("Datas", d1.Equal(d2))

	type Address struct {
		ZipCode int64
		Country string
		State   string
		City    string
		Street  string
		Number  int64
	}

	type Documentos struct {
		CPF int
		RG int
		CNH int
	}

	type Pessoa struct {
		Nome    string // não entendi "Nome"
		Idade   int
		Profissao string
		Documentos []Documentos
		Address []Address
		
	}

	documentos := Documentos{
		CPF: 222,
		RG: 333,
		CNH: 444,
	}

	addressWork := Address{
		ZipCode: 06230150,
		Country: "Brazil",
		State:   "São Paulo",
		City:    "Cotia",
		Street:  "Rua peçanha",
		Number:  200,
	}

	addressHome := Address{
		ZipCode: 06230 - 150,
		Country: "Brazil",
		State:   "São Paulo",
		City:    "Osasco",
		Street:  "Av Edmundo Amaral",
		Number:  130,
	}

	p1 := Pessoa{"Ricardo", 30, "mecanico", []Documentos{documentos}, []Address{addressWork, addressHome}}
	p2 := Pessoa{"Lula", 28, "Ladrão", []Documentos{documentos}, []Address{addressWork, addressHome}}
	p3 := Pessoa{"Fabio", 27, "Desenhista", []Documentos{documentos}, []Address{addressWork, addressHome}}
	fmt.Println("\nPessoas", p1.Nome == p2.Nome)
	fmt.Println("\nPessoas", p2)
	fmt.Println("\nPessoas", p3)

	fmt.Println("\nFirst Address from Ricardo", p1.Address[0]) //acessa indice 0
}
