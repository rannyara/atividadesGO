package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rannyara/meu-projeto-go/internal/memorygame"
)

// checkExit verifica se o usuário quer sair do jogo
func checkExit(input string) bool {
	return strings.ToLower(strings.TrimSpace(input)) == "sair"
}

// renderExitMessage mostra mensagem de despedida com o tabuleiro revelado
func renderExitMessage(game *memorygame.Game) {
	fmt.Println("\nJogo terminado! Até a próxima! 👋")
	fmt.Println("Tabuleiro completo:")
	game.Render(true)
}

func main() {
	// Testar o Jogo da Memória
	fmt.Println("🎮 JOGO DA MEMÓRIA 🎮")
	fmt.Println("====================")

	// Configurar tamanho do tabuleiro
	rows, cols := 4, 4 // Tabuleiro 4x4 (8 pares)

	game, err := memorygame.NewGame(rows, cols)
	if err != nil {
		fmt.Printf("Erro ao criar jogo: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Tabuleiro %dx%d (%d pares)\n", rows, cols, game.TotalPairs)
	fmt.Println("Digite as coordenadas no formato: linha1 coluna1 linha2 coluna2")
	fmt.Println("Exemplo: 0 1 2 3")
	fmt.Println("Digite 'sair' a qualquer momento para terminar o jogo")
	fmt.Println()

	for !game.GameOver() {
		// Mostrar tabuleiro atual
		game.Render(false)
		fmt.Printf("Tentativas: %d | Pares encontrados: %d/%d\n",
			game.Moves, game.PairsFound, game.TotalPairs)

		// Ler entrada do usuário
		fmt.Print("Sua jogada (ou 'sair' para encerrar): ")
		scanner.Scan()
		input := scanner.Text()

		// Verificar se quer sair
		if checkExit(input) {
			renderExitMessage(game)
			return
		}

		// Processar coordenadas
		coords := strings.Fields(input)
		if len(coords) != 4 {
			fmt.Println("❌ Digite exatamente 4 números separados por espaço")
			fmt.Println("   Ou digite 'sair' para encerrar o jogo")
			continue
		}

		// Converter coordenadas
		r1, err1 := strconv.Atoi(coords[0])
		c1, err2 := strconv.Atoi(coords[1])
		r2, err3 := strconv.Atoi(coords[2])
		c2, err4 := strconv.Atoi(coords[3])

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Println("❌ Coordenadas devem ser números")
			fmt.Println("   Ou digite 'sair' para encerrar o jogo")
			continue
		}

		// Tentar virar as cartas
		matched, err := game.FlipPair(r1, c1, r2, c2)
		if err != nil {
			fmt.Printf("❌ Erro: %v\n", err)
			fmt.Println("   Digite 'sair' se quiser encerrar o jogo")
			continue
		}

		// Mostrar resultado
		game.Render(false)
		if matched {
			fmt.Println("✅ Par encontrado!")
		} else {
			fmt.Println("❌ Não é um par. Tente novamente!")
			// Aguardar um pouco para o jogador ver as cartas
			fmt.Println("Pressione Enter para continuar...")
			fmt.Println("Ou digite 'sair' para encerrar o jogo")
			scanner.Scan()

			// Verificar se quer sair durante a pausa
			if checkExit(scanner.Text()) {
				renderExitMessage(game)
				return
			}

			// Esconder cartas não pareadas
			game.HideNonMatched()
		}
		fmt.Println()
	}

	// Jogo terminado naturalmente
	game.Render(true)
	fmt.Println("🎉 PARABÉNS! Você completou o jogo!")
	fmt.Printf("Tentativas: %d | Tempo: %.0f segundos\n",
		game.Moves, game.Elapsed().Seconds())
}
