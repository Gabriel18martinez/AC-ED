package arrays

import (
	"fmt"
	"projeto/contato"
)

func alteraEmail(c *[]contato.Contato, indice int, emailNovo string) {
	if (indice >= 0 && indice < len(*c)){
		c[indice].alterarEmail(emailNovo)
	}
}

func adicionaContato(c contato.Contato, a *[5]contato.Contato){
	for ind, contatinho := range a {
		if (contatinho == contato.Contato{}) {
			a[ind] = c
			break
		}
	}
}

func excluiContato(a *[5]contato.Contato){
	if (a[0] == contato.Contato{}) {
		fmt.Println("Lista de contatos vazia!")
	}

	for ind, contatinho := range a {
		if (contatinho == contato.Contato{}) {
			a[ind-1] = contato.Contato{}
		}
	}
}