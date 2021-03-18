package main

import "fmt"

func main() {
	var variavel1 string = "variavel1"
	fmt.Println(variavel1)

	var variavel2 = "variavel2"
	fmt.Println(variavel2)

	const constante1 string = "contante1"
	fmt.Println(constante1)

	var contante2 = "contante2"
	fmt.Println(contante2)

	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)

	fmt.Println(variavel3, variavel4)

	var variavel5, variavel6 = "variavel5", "variavel6"

	fmt.Println(variavel5, variavel6)
}
