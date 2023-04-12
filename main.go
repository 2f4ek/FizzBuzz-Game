package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game interface {
	InitGame()
	RoundInit()
	CheckAnswer()
	RoundSuccess()
	RoundFailed()
	FizzBuzzRound()
	FizzRound()
	BuzzRound()
	DefaultRound()
}

type GameProgress struct {
	round int
	input string
}

func (p *GameProgress) InitGame() {
	fmt.Println("Game Started!")
	p.RoundInit()
}

func (p *GameProgress) RoundInit() {
	fmt.Printf("Round %d and your answer is -> ", p.round)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p.input)
	p.input = strings.TrimSuffix(input, "\r")
	p.input = strings.TrimSuffix(input, "\n")
	p.input = strings.ToLower(p.input)
	p.input = strings.ReplaceAll(p.input, " ", "")
	p.CheckAnswer()
}

func (p *GameProgress) CheckAnswer() {
	switch {
	case (p.round%3 == 0) && (p.round%5 == 0):
		p.FizzBuzzRound()
	case (p.round%3 == 0) && (p.round%5 != 0):
		p.FizzRound()
	case (p.round%3 != 0) && (p.round%5 == 0):
		p.BuzzRound()
	default:
		p.DefaultRound()
	}
}

func (p *GameProgress) FizzBuzzRound() {
	fmt.Println("fizzbuzz round")
	if p.input == "fizzbuzz" {
		p.RoundSuccess()
	}
	p.RoundFailed()
}

func (p *GameProgress) FizzRound() {
	fmt.Println("fizz round")
	if p.input == "fizz" {
		p.RoundSuccess()
	}
	p.RoundFailed()
}

func (p *GameProgress) BuzzRound() {
	fmt.Println("buzz round")
	if p.input == "buzz" {
		p.RoundSuccess()
	}
	p.RoundFailed()
}

func (p *GameProgress) DefaultRound() {
	if strings.Compare(p.input, strconv.Itoa(p.round)) == 0 {
		p.RoundSuccess()
	}
	p.RoundFailed()
}

func (p *GameProgress) RoundSuccess() {
	fmt.Println("Correct!")
	p.round += 1
	p.RoundInit()
}

func (p *GameProgress) RoundFailed() {
	fmt.Printf("You lose!")
	os.Exit(200)
}

func main() {
	var pl Game = &GameProgress{round: 1}
	for {
		pl.InitGame()
	}
}
