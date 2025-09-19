package jogotop

import (
	"bufio"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func LimparTela() {
	// Tenta limpar via ANSI + garante espaço em branco como fallback
	fmt.Print("\033[H\033[2J")
	fmt.Print(strings.Repeat("\n", 10))
}

func GerarSequencia(tamanho, digitoMax int) []int {
	seq := make([]int, tamanho)
	for i := 0; i < tamanho; i++ {
		seq[i] = rand.Intn(digitoMax + 1) // 0..digitoMax (ex.: 0..9)
	}
	return seq
}

func Juntar(nums []int) string {
	partes := make([]string, len(nums))
	for i, v := range nums {
		partes[i] = strconv.Itoa(v)
	}
	return strings.Join(partes, " ")
}

func LerPalpite(reader *bufio.Reader, tamanho int) []int {
	for {
		fmt.Printf("Digite a sequência (%d números de 0 a 9, separados por espaço): ", tamanho)
		linha, _ := reader.ReadString('\n')
		linha = strings.TrimSpace(linha)

		partes := strings.Fields(linha)
		if len(partes) != tamanho {
			fmt.Printf("Ops! São exatamente %d números. Tenta de novo.\n", tamanho)
			continue
		}

		palpite := make([]int, tamanho)
		ok := true
		for i, p := range partes {
			v, err := strconv.Atoi(p)
			if err != nil || v < 0 || v > 9 {
				fmt.Println("Use apenas dígitos de 0 a 9.")
				ok = false
				break
			}
			palpite[i] = v
		}
		if ok {
			return palpite
		}
	}
}

func Pontuar(seq, palpite []int) int {
	acertos := 0
	for i := range seq {
		if seq[i] == palpite[i] {
			acertos++
		}
	}
	return acertos
}

