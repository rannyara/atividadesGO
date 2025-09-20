package main

import (
	"fmt"

	"github.com/rannyara/meu-projeto-go/internal/fibonacci"
)

func main() {
	fmt.Println("🔢 Testando Fibonacci!")
	n := 10
	result := fibonacci.Fibonacci(n)
	fmt.Printf("O %dº número de Fibonacci é: %d\n", n, result)
}
