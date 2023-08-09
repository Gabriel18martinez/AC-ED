package main

import "fmt"

var y = false
var i = 1


func ePrimo(n int) bool {
	if n == 1{
		y = false
	} else if n == 0 {
		y = false
	} else if n < 0 {
		println("Valor inválido!")
	} else{
		for i < n {
			i++
			if n % i == 0 {
				fmt.Println(i, "é divisível por " , n)
				continue
			} else {
				y = true
				break
			}
			
		}
	}

	return y
}

func fibo(n int) int {
	var x = 1
	var z = 1
	var j = 0
	for i < n + 1 {
		i++
		j = x + z
		x = z
		z = j
	}

	return j
}

func diaSemana(n int) {
	switch n {
		case 1:
			fmt.Println("Domingo")
		case 2:
			fmt.Println("Segunda")
		case 3:
			fmt.Println("Terça")
		case 4:
			fmt.Println("Quarta")
		case 5:
			fmt.Println("Quinta")
		case 6:
			fmt.Println("Sexta")
		case 7:
			fmt.Println("Sabado")
		default:
			fmt.Println("Dia inválido.")
	}
}

func e_bissexto(ano int){
	if ano % 4 == 0{
		if ano % 100 == 0{
			if ano % 400 == 0 {
				fmt.Println("Esse ano é bissextor")
			} else {
				fmt.Println("Esse ano não é bissexto")
			}
		} else{
			fmt.Println("Esse ano é bissexto")
		} 
	} else {
		fmt.Println("Esse ano não é bissexto")
	}
}


func main() {
	fmt.Println(ePrimo(7))
	fmt.Println(fibo(7))
	diaSemana(7)
	e_bissexto(20)
}