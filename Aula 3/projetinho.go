package main

import (
	"fmt"
)

type Carro struct {
	Modelo, cor string
	velocidade float64
	EstaLigado bool
}

func (c *Carro) ligar(){
	c.EstaLigado = true
	fmt.Println("O carro esta ligado vrum vrum!")
}

func (c *Carro) desligar(){
	c.EstaLigado = false
	fmt.Println("O carro esta desligado!")
}

func (c *Carro) acelerar(valor float64){
	if c.EstaLigado == true {
		c.velocidade = c.velocidade + valor
		fmt.Println("A velocidade do carro e:", c.velocidade, "km/h")
	}
}

func main(){
	c:= Carro{
		"Honda civic", "Azul", 0, false,
	}
	c.ligar()
	c.acelerar(5.0)
	c.desligar()
}