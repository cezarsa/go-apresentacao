package main

type Nomeavel interface {
	Nome() string
}

type Humano struct {
	primeiroNome string
}

func (h Humano) Nome() string {
	return "Humano: " + h.primeiroNome
}

type Objeto struct {
	descricao string
}

func (o Objeto) Nome() string {
	return "Objeto: " + o.descricao
}

func main() {
	val := []Nomeavel{
		Humano{primeiroNome: "João"},
		Objeto{descricao: "Redondo"},
	}
	for _, v := range val {
		println(v.Nome())
	}
}
