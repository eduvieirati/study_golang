package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println("Função F")
		return txt
	}

	result := f("Texto da função 1")
	fmt.Println(result)

	resultSum, _ := calculos(8, 7)
	fmt.Println(resultSum)
}
