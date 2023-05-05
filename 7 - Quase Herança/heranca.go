package main

import "fmt"

type pessoa struct {
	nome      string
	idade     int
	altura    float32
	sobrenome string
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Tme")
	var jeter pessoa
	jeter = pessoa{"Jeter", 28, 200, "Megaron"}
	fmt.Println(jeter)

	var aluno estudante = estudante{jeter, "Artes Gráficas", "Unileão"}
	fmt.Println(aluno)
}
