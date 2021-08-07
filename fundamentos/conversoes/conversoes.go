package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado ao passar inteiros para converções de string
	// irá retornar o caracter da tabela ASC
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))
	fmt.Println("Teste " + fmt.Sprint(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println("num", num)
	fmt.Println("num subtração", num - 122)

	b, _ := strconv.ParseBool("bool")
	if b {
		fmt.Println("Verdadeiro")
	}
}
