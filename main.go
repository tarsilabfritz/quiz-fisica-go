package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Estrutura que representa uma pergunta do quiz
type GameQuestions struct {
	Text    string
	Options []string
	Answer  int
}

// Estrutura que representa o estado do jogo
type GameState struct {
	Name      string
	Points    int
	Questions []GameQuestions
}

// Inicializa o jogo
func (g *GameState) Init() {
	fmt.Println("Seja bem-vindo(a) ao quiz!")
	fmt.Println("Escreva o seu nome:")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')
	if err != nil {
		panic("Erro ao ler a string")
	}

	g.Name = strings.TrimSpace(name) // Remove espaços extras e nova linha
	fmt.Printf("Vamos ao jogo, %s!\n", g.Name)
}

// Processa o arquivo CSV
func (g *GameState) ProccessCSV() {
	f, err := os.Open("QuizGo.csv")
	if err != nil {
		panic("Erro ao ler arquivo")
	}

	defer f.Close() // Garante que o arquivo será fechado após o uso

	reader := csv.NewReader(f) // Cria um leitor CSV
	records, err := reader.ReadAll()
	if err != nil {
		panic("Erro ao ler csv")
	}

	// Itera sobre todas as linhas do CSV
	for index, record := range records {
		if index > 0 { // Ignora a primeira linha (cabeçalho)
			question := GameQuestions{
				Text:    record[0],
				Options: record[1:5],
				Answer:  toInt(record[5]),
			}
			g.Questions = append(g.Questions, question)
		}
	}
}

// Executa o jogo
func (g *GameState) Run() {
	for index, question := range g.Questions {
		fmt.Printf("\033[33m%d. %s\033[0m\n", index+1, question.Text)

		// Exibe as opções de resposta
		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		var userAnswer int
		validInput := false
		// Loop para garantir que a resposta seja válida
		for !validInput {
			fmt.Println("Digite uma alternativa (1-4):")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				panic("Erro ao ler entrada do usuário")
			}

			input = strings.TrimSpace(input) // Remove espaços extras
			userAnswer, err = strconv.Atoi(input)
			if err != nil || userAnswer < 1 || userAnswer > 4 {
				fmt.Println("Entrada inválida, por favor escolha uma alternativa de 1 a 4.")
			} else {
				validInput = true
			}
		}

		// Verifica se a resposta do usuário está correta
		if userAnswer == question.Answer {
			fmt.Println("Resposta correta!")
			g.Points++
		} else {
			fmt.Println("Resposta incorreta.")
		}
	}

	// Exibe a pontuação final do jogador
	fmt.Printf("\nParabéns, %s! Você fez %d ponto(s)!\n", g.Name, g.Points)
}

// Função principal
func main() {
	game := &GameState{}
	game.ProccessCSV()
	game.Init()
	game.Run()
}

// Converte string para inteiro
func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
