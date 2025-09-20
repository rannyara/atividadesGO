package main

import (
	"fmt"

	"github.com/rannyara/meu-projeto-go/internal/anamnese"
)

func main() {
	// Coletar dados do paciente
	paciente := anamnese.ColetarDadosPaciente()

	// Exibir resultados
	anamnese.ExibirResultados(paciente)

	// Obter e exibir recomendações
	recomendacoes := anamnese.ObterRecomendacoes(paciente.IMC)

	fmt.Println()
	fmt.Println("RECOMENDAÇÕES:")
	for i, rec := range recomendacoes {
		fmt.Printf("%d. %s\n", i+1, rec)
	}
}
