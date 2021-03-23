package main

import "fmt"

func somar(a, b int8) int8 {
	return a + b
}

func retornaVarios(a, b int8) (int8, int8) {
	soma := a + b
	subtracao := a - b
	return soma, subtracao
}

func main() {
	soma := somar(10, 10)
	fmt.Println(soma)

	var fun = func(txt string) {
		fmt.Println(txt)
	}

	var fun2 = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	fun("teste")

	resultado := fun2("Golang")

	fmt.Println(resultado)

	adicao, subtracao := retornaVarios(10, 5)
	fmt.Println(adicao, subtracao)

	adicao1, _ := retornaVarios(10, 5)
	fmt.Println(adicao1)
}
