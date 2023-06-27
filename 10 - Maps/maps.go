package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Eduardo",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "David",
			"ultimo":   "Eduardo",
		},
		"curso": {
			"nome":   "ADS",
			"campus": "Estácio",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome") //Deleta uma chave
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Cancêr",
	}

	fmt.Println(usuario2)
}
