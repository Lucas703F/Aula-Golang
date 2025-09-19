// Programa executável do jogo da memória (CLI).
// Uso:
//
//	go run ./cmd/memorygame
//
// Depois siga instruções na tela.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"math/rand"

	"github.com/seu-usuario/meu-projeto-go/internal/memorygame"
	"github.com/seu-usuario/meu-projeto-go/internal/jogotop"
)

func main() {
	var escolha int
	fmt.Println("Escolha o jogo para jogar:")
	fmt.Println("1 - Jogo da Memória")
	fmt.Println("2 - Jogo da sequência")
	fmt.Print("Digite 1 ou 2: ")

	fmt.Scan(&escolha)

	switch escolha {
	case 1:
		fmt.Println("🧠 Jogo da Memória - CLI em Go")
		fmt.Println("Encontre todos os pares! Formato de entrada: r1 c1 r2 c2 (ex: 0 0 1 0)")
		fmt.Println("Digite 'sair' para encerrar antecipadamente.")
		fmt.Println("")

		rows, cols := 4, 4 // Tabuleiro 4x4 -> 8 pares
		game, err := memorygame.NewGame(rows, cols)
		if err != nil {
			fmt.Println("Erro criando jogo:", err)
			return
		}

		reader := bufio.NewReader(os.Stdin)
		for {
			// Mostrar tabuleiro atual
			game.Render(false)
			if game.GameOver() {
				fmt.Println("🎉 Parabéns! Você concluiu o jogo!")
				fmt.Printf("Tentativas: %d | Tempo: %v\n", game.Moves, game.Elapsed().Round(time.Millisecond))
				fmt.Println("Tabuleiro final:")
				game.Render(true)
				break
			}

			fmt.Print("Sua jogada (r1 c1 r2 c2): ")
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			if strings.EqualFold(line, "sair") {
				fmt.Println("Saindo... até a próxima!")
				return
			}
			parts := strings.Fields(line)
			if len(parts) != 4 {
				fmt.Println("Entrada inválida. Use: r1 c1 r2 c2")
				continue
			}
			nums := make([]int, 4)
			ok := true
			for i, p := range parts {
				v, err := strconv.Atoi(p)
				if err != nil {
					fmt.Println("Valor não numérico:", p)
					ok = false
					break
				}
				nums[i] = v
			}
			if !ok {
				continue
			}

			matched, err := game.FlipPair(nums[0], nums[1], nums[2], nums[3])
			if err != nil {
				fmt.Println("Erro:", err)
				continue
			}
			// Mostrar cartas viradas
			game.Render(false)
			if matched {
				fmt.Println("✅ Par encontrado!")
			} else {
				fmt.Println("❌ Não foi par. Cartas serão ocultadas...")
				time.Sleep(1500 * time.Millisecond)
				game.HideNonMatched()
			}
		}
	case 2:
		rand.Seed(time.Now().UnixNano())

		const tamanho = 5
		const digitoMax = 9

		reader := bufio.NewReader(os.Stdin)

		fmt.Println("=== Jogo da Memória (Go) ===")
		fmt.Println("Objetivo: memorize a sequência e digite na mesma ordem.")
		fmt.Println()

		seq := jogotop.GerarSequencia(tamanho, digitoMax)

		// Mostra a sequência para memorizar
		fmt.Println("Sequência (memorize):", jogotop.Juntar(seq))
		fmt.Println("Pressione ENTER quando estiver pronto para responder...")
		_, _ = reader.ReadString('\n')

		// "Limpa" a tela (ou empurra para cima)
		jogotop.LimparTela()

		// Pede o palpite
		palpite := jogotop.LerPalpite(reader, tamanho)

		// Calcula pontuação
		acertos := jogotop.Pontuar(seq, palpite)

		// Mostra resultado
		fmt.Println()
		fmt.Println("Sequência correta:", jogotop.Juntar(seq))
		fmt.Println("Sua resposta     :", jogotop.Juntar(palpite))
		fmt.Printf("Você fez %d/%d ponto(s). ", acertos, tamanho)

		switch acertos {
		case 5:
			fmt.Println("Perfeito! Mandou muito bem!")
		case 3, 4:
			fmt.Println("Quase lá! Bora de novo?")
		case 1, 2:
			fmt.Println("Tá no caminho. Treina mais um pouco!")
		default:
			fmt.Println("Agora vai! Tenta outra vez!")
		}
	}
}
