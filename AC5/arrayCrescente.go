package main

import "fmt"


func encontrarPar(array []int, alvo int) (int, int) {

	esquerda, direita := 0, len(array)-1

	for esquerda < direita {
		soma := array[esquerda] + array[direita]

		if soma == alvo {
			return array[esquerda], array[direita]
		} else if soma < alvo {
			esquerda++
		} else {
			direita--
		}
	}
	return -1, -1
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(encontrarPar(array, 10))
}
