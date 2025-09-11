// Arquivo principal do programa (entrypoint)
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários

import (
	"fmt"

	"github.com/rannyara/meu-projeto-go/internal/fibonacci"
	"github.com/rannyara/meu-projeto-go/internal/hello"
)

// Função principal do programa
func main() {
	fmt.Println("🚀 Fiz Minha primeira atividade em GO!")
	hello.SayHello()

	// Testar a função Fibonacci
	n := 10
	result := fibonacci.Fibonacci(n)
	fmt.Printf("O %dº número de Fibonacci é: %d\n", n, result)
}
