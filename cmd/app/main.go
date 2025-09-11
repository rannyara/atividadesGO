// Arquivo principal do programa (entrypoint)
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"

	"github.com/seu-usuario/meu-projeto-go/internal/hello"
)

// Função principal do programa
func main() {
	fmt.Println("🚀 Minha primeira atividade em GO!")
	hello.SayHello()
}
