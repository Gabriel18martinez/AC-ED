package contato

type Contato struct {
	Nome, Email string
}

func (c *Contato) alterarEmail(novoEmail string) {

	c.Email = novoEmail

}