package main

import "time"
import "fmt"

func main(){
	i := 0

	for i < 10{
		i++
		fmt.Println("Incrementando", i)
		time.Sleep(time.Second)
		
	}

	nomes := [3]string{"João","David","Lucas"}

	for  nome := range nomes{
		fmt.Println(nome)
	}
}