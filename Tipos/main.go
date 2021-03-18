package main

import (
	"errors"
	"fmt"
)

func main() {
	// Inteiros
	var numero int64 = -100000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)
	// FIM inteiros

	// Números reais

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// fim reais

	// String
	var str string = "Golang"
	fmt.Println(str)

	str2 := "Maracujá"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	// Fim string

	texto := 5
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Sem vacina")
	fmt.Println(erro)
}
