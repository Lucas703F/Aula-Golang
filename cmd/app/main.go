// Arquivo principal do programa (entrypoint) 🫡
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"
	"os"

	"github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
	"github.com/seu-usuario/meu-projeto-go/internal/hello"
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

	nome := "Lucas"

	switch nome {
	case "Lucas":
		fmt.Println("\nLucas Felipe de Araújo Ferreira")
	default:
		fmt.Println("Nome não encontrado!")
	}

	// Função anônima atribuída a uma variável
	funcaoAnonima := func() {
		fmt.Println("Olá! Esta é uma função anônima.")
	}

	// Executando a função anônima
	funcaoAnonima()

	// Chamar a função IMC para rodar o cálculo
	IMC()

}

func calcularIMC(peso, altura float64) float64 {
	return peso / (altura * altura)
}

func IMC() {
	var nomeImc string
	var idade int
	var peso, altura float64

	// Entrada dos dados
	fmt.Print("\nDigite seu nome: ")
	fmt.Scanln(&nomeImc)

	fmt.Print("\nDigite sua idade: ")
	fmt.Scanln(&idade)

	fmt.Print("\nDigite seu peso (kg): ")
	fmt.Scanln(&peso)

	fmt.Print("\nDigite sua altura (em metros, ex: 1.75): ")
	fmt.Scanln(&altura)

	// Cálculo do IMC
	imc := calcularIMC(peso, altura)

	// Exibição do resultado
	fmt.Printf("\n%s, %d anos\n", nomeImc, idade)
	fmt.Printf("Seu IMC é: %.2f\n", imc)

	// Avaliação do IMC com switch
	switch {
	case imc < 18.5:
		fmt.Println("Classificação: Abaixo do peso")
	case imc >= 18.5 && imc < 24.9:
		fmt.Println("Classificação: Peso ideal")
	case imc >= 25:
		fmt.Println("Classificação: Acima do peso")
	default:
		fmt.Println("Classificação: Indefinida")
	}
}
