package anamnese

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Paciente struct para armazenar os dados do paciente
type Paciente struct {
	Nome   string
	Idade  int
	Peso   float64
	Altura float64
	IMC    float64
}

// CalcularIMC calcula o Índice de Massa Corporal
func CalcularIMC(peso, altura float64) float64 {
	return peso / (altura * altura)
}

// ClassificarIMC retorna a classificação do IMC
func ClassificarIMC(imc float64) string {
	switch {
	case imc < 18.5:
		return "Abaixo do peso"
	case imc < 25:
		return "Peso normal"
	case imc < 30:
		return "Sobrepeso"
	case imc < 35:
		return "Obesidade grau I"
	case imc < 40:
		return "Obesidade grau II"
	default:
		return "Obesidade grau III"
	}
}

// ColetarDadosPaciente coleta e valida os dados do paciente
func ColetarDadosPaciente() Paciente {
	reader := bufio.NewReader(os.Stdin)
	var paciente Paciente
	var err error

	fmt.Println("=== SISTEMA DE ANAMNESE ===")
	fmt.Println()

	// Coletar nome
	fmt.Print("Nome do paciente: ")
	paciente.Nome, _ = reader.ReadString('\n')
	paciente.Nome = strings.TrimSpace(paciente.Nome)

	// Coletar idade
	for {
		fmt.Print("Idade: ")
		idadeStr, _ := reader.ReadString('\n')
		idadeStr = strings.TrimSpace(idadeStr)
		paciente.Idade, err = strconv.Atoi(idadeStr)

		if err != nil || paciente.Idade <= 0 || paciente.Idade > 150 {
			fmt.Println("Idade inválida! Digite um número entre 1 e 150.")
			continue
		}
		break
	}

	// Coletar peso
	for {
		fmt.Print("Peso (kg): ")
		pesoStr, _ := reader.ReadString('\n')
		pesoStr = strings.TrimSpace(pesoStr)
		paciente.Peso, err = strconv.ParseFloat(pesoStr, 64)

		if err != nil || paciente.Peso <= 0 || paciente.Peso > 500 {
			fmt.Println("Peso inválido! Digite um valor entre 0.1 e 500 kg.")
			continue
		}
		break
	}

	// Coletar altura
	for {
		fmt.Print("Altura (m): ")
		alturaStr, _ := reader.ReadString('\n')
		alturaStr = strings.TrimSpace(alturaStr)
		paciente.Altura, err = strconv.ParseFloat(alturaStr, 64)

		if err != nil || paciente.Altura <= 0 || paciente.Altura > 3 {
			fmt.Println("Altura inválida! Digite um valor entre 0.1 e 3.0 metros.")
			continue
		}
		break
	}

	// Calcular IMC
	paciente.IMC = CalcularIMC(paciente.Peso, paciente.Altura)

	return paciente
}

// ExibirResultados exibe os resultados da anamnese
func ExibirResultados(p Paciente) {
	fmt.Println()
	fmt.Println("=== RESULTADOS DA ANAMNESE ===")
	fmt.Printf("Nome: %s\n", p.Nome)
	fmt.Printf("Idade: %d anos\n", p.Idade)
	fmt.Printf("Peso: %.1f kg\n", p.Peso)
	fmt.Printf("Altura: %.2f m\n", p.Altura)
	fmt.Printf("IMC: %.1f\n", p.IMC)
	fmt.Printf("Classificação: %s\n", ClassificarIMC(p.IMC))
	fmt.Println("==============================")
}

// ObterRecomendacoes retorna recomendações baseadas no IMC
func ObterRecomendacoes(imc float64) []string {
	var recomendacoes []string

	switch {
	case imc < 18.5:
		recomendacoes = []string{
			"Consulte um nutricionista para adequação alimentar",
			"Realize acompanhamento médico regular",
			"Foque em alimentação nutritiva e balanceada",
		}
	case imc < 25:
		recomendacoes = []string{
			"Parabéns! Mantenha seus hábitos saudáveis",
			"Continue com atividades físicas regulares",
			" Mantenha alimentação equilibrada",
		}
	case imc < 30:
		recomendacoes = []string{
			"Atenção ao risco de desenvolver obesidade",
			"Considere aumentar a atividade física",
			"Reduza alimentos ultraprocessados",
		}
	default:
		recomendacoes = []string{
			"Recomenda-se acompanhamento médico especializado",
			"Procure um nutricionista para plano alimentar",
			"Atividade física orientada é essencial",
			"Monitoramento regular da saúde",
		}
	}

	return recomendacoes
}
