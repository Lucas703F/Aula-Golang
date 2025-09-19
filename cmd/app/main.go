// Arquivo principal do programa (entrypoint) 🫡
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"
	"os"

	"github.com/seu-usuario/meu-projeto-go/internal/hello"
	"github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
)

// Função principal do programa
func main() {
	// Mensagem inicial da aplicação
	fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")
	hello.SayHello()

	var valorEntrada int
	fmt.Print("\n\nEscolha um valor para calcular: ")
	fmt.Fscan(os.Stdin, &valorEntrada)
	result := fibonacci.Fibonacci(valorEntrada)
	fmt.Printf("O resultado do Fibonacci(%v) é: %v\n", valorEntrada, result)
}
