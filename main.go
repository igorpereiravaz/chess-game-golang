package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da adivinhação")
	fmt.Println("Um número aleatório será sorteado. Tente acertar. O número é um inteiro entre 0 e 100")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Printf("Digite o %dº chute: ", i+1)
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O chute deve ser um número inteiro")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Seu chute foi menor que o número sorteado")
		case chuteInt > x:
			fmt.Println("Seu chute foi maior que o número sorteado")
		case chuteInt == x:
			fmt.Printf(
				"Parabéns, você acertou o número sorteado: %d\n"+
					"Você acertou em %d tentativas.\n"+
					"Essas foras as suas tentativas: %v\n",
				x, i+1, chutes[:i],
			)
			return
		}

		chutes[i] = chuteInt
	}
	fmt.Printf(
		"Você não acertou o número sorteado, o número correto era: %d\n"+
			"Você teve 10 tentativas.\n"+
			"Essas foras as suas tentativas: %v\n",
		x, chutes,
	)
}
