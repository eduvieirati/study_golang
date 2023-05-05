package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int64 = -1000000000
	fmt.Println(number)

	var number2 uint32 = 10000
	fmt.Println(number2)

	// Alias
	// INT32 = rune
	var number3 rune = 12456
	fmt.Println(number3)

	// BYTE= UINT8
	var number4 byte = 123
	fmt.Println(number4)

	// Números reais
	var number5 float32 = 123.45
	fmt.Println(number5)

	number6 := 123654.67
	fmt.Println(number6)

	// Strings
	var str string = "Text"
	fmt.Println(str)

	str2 := "Text2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// Valor zero

	texto := 5
	fmt.Println(texto)

	var boolean bool
	fmt.Println(boolean)

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
