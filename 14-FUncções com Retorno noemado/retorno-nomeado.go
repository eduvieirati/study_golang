package main

func calcular(n1,n2 int) (soma int, subtracao int){
	soma = n1 + n2
	subtracao = n1 - n2
	return 
}

func main(){
	soma, subtracao := calcular(10,15)
	println(soma, subtracao)
}