package main

import (
	"fmt"
)

func main() {
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
