package main

import (
	"fmt"

	"github.com/rannyara/meu-projeto-go/internal/saudacao"
)

func main() {
	// Testar a função Saudacao
	nome := "Lorranny"
	fmt.Println(saudacao.Saudacao(nome))

	// Criando uma função anônima para saudação
	saudacao2 := func(nome string) string {
		return "Olá, " + nome
	}

	// Usando a função anônima
	name := "Lorranny"
	fmt.Println(saudacao2(name))
}
