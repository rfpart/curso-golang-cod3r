package main

import "fmt"

func main() {
	i := 1
	
	var p *int = nil
	p = &i // pegando endereço da variável i e atribuindo para o ponteiro
	*p++ // desreferenciando (pegando o valor)
	i++  

	//Go não tem aritmética de ponteiros

	fmt.Println(p, *p, i, &i)
}