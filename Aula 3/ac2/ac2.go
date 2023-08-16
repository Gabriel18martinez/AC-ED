package main

import "fmt"

type Contato struct {
	Nome, Email string
}

func (d *Contato) alteraEmail(email string){
	d.Email = email
}

func adicionaContato(c Contato, n [5]Contato) [5]Contato{
	for indice, contato := range n {
		if (contato == Contato{}){
			n[indice] = c
			break
		}
	}
	return n
}

func excuiContato(n [5]Contato) [5]Contato {

	for indice, contato := range n {
		if (contato == Contato{}){
			n[indice - 1] = Contato{}
			break
		}
	}

	return n
}



func main(){
	c:=Contato{
		Nome: "jose",
		Email: "elias@gmai.com",
	}

	d:=Contato{
		Nome: "Gabriel",
		Email: "elias@gmai.com",
	}
	
	f:=Contato{
		Nome: "Carlos",
		Email: "carlos@gmai.com",
	}

	g:=Contato{
		Nome: "joao",
		Email: "joao@gmai.com",
	}



	c.alteraEmail("jose@gmail.com")

	fmt.Println(c.Email)

	var contatos [5]Contato

	contatos = adicionaContato(c, contatos)
	contatos = adicionaContato(d, contatos)
	contatos = adicionaContato(f, contatos)
	contatos = adicionaContato(g, contatos)

	fmt.Println(contatos)

	contatos = excuiContato(contatos)

	fmt.Println(contatos)

}