package main

import "fmt"

const M = 5

func main() {
	
	var lista [M]int
	var n = 0

	insereOrd(&lista, &n, 4)
	fmt.Println(lista)

	insereOrd(&lista, &n, 12)
	fmt.Println(lista)

	insereOrd(&lista, &n, 2)
	fmt.Println(lista)

	insereOrd(&lista, &n, 6)
	fmt.Println(lista)

	insereOrd(&lista, &n, 17)
	fmt.Println(lista)

	insereOrd(&lista, &n, 1)
	fmt.Println(lista)
}

func buscaBin(v [M]int, n, x int) int {
	inf, sup := 0, n-1
	for inf <= sup {
		meio := int((inf + sup) / 2)
		if v[meio] == x {
			return meio
		}
		if x < v[meio] {
			sup = meio - 1
		} else {
			inf = meio + 1
		}
	}
	return -1
}

func insereOrd(v *[M]int, n *int, valor int) {
	ind := 0
	if buscaBin(*v, *n, valor) != -1 {
		fmt.Println("Elemento ja existe na lista")
		
	} else {
		if *n == M{
			fmt.Println("Overflow")
		} 
		for i := ind; i < *n; i++ {
			if v[i] > valor {
				ind = i
				break
			} else {
				ind = i
			}
		}
		for i := *n - 1 ; i > ind ; i-- {
			v[i] = v[i - 1]
		}

		v[ind] = valor
		*n++
	}
}