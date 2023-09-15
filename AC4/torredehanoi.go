package main

import "fmt"

func torreHanoi(n int, origem, destino, trabalho string) {
   if n > 0{
    torreHanoi(n-1, origem, trabalho, destino)
    fmt.Printf("Movimenta disco %d de %s para %s\n", n, origem, destino)
    torreHanoi(n-1, trabalho, destino, origem)
   }
}

func main() {
    torreHanoi(3, "A", "C", "B")
}