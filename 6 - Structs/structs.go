package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco endereco
}

type endereco struct{
	rua string
	numero int
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Eduardo"
	u.idade = 21
	fmt.Println(u)

	exemploEndereco := endereco{"Avenida Romeiro Aureliano Pereira", 055}
	usuario2 := usuario{"Megatronis", 28, exemploEndereco}
	fmt.Println(usuario2)
}
