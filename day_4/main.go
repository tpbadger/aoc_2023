package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	id              int
	winning_numbers []int
	numbers         []int
}

func XinY(x int, y []int) bool {
	for _, comp := range y {
		if x == comp {
			return true
		}
	}
	return false
}

func (c *Card) GetWinners() []int {
	winners := []int{}
	for _, number := range c.numbers {
		if XinY(number, c.winning_numbers) {
			winners = append(winners, number)
		}
	}
	return winners
}

func GenerateCards(inputs []string) []Card {
	cards := []Card{}
	for _, input := range inputs {
		m := regexp.MustCompile(`[0-9]+`)
		split := strings.Split(input, ":")
		card_id, err := strconv.Atoi(m.FindString(split[0]))
		if err != nil {
			panic(err)
		}

		winning_numbers := []int{}
		numbers := []int{}

		for _, winning_number_str := range m.FindAllString(strings.Split(split[1], "|")[0], -1) {
			winning_number, err := strconv.Atoi(winning_number_str)
			if err == nil {
				winning_numbers = append(winning_numbers, winning_number)
			}
		}

		for _, number_str := range m.FindAllString(strings.Split(split[1], "|")[1], -1) {
			number, err := strconv.Atoi(number_str)
			if err == nil {
				numbers = append(numbers, number)
			}
		}

		cards = append(cards,
			Card{
				id:              card_id,
				winning_numbers: winning_numbers,
				numbers:         numbers,
			},
		)
	}

	return cards

}

func PartOne(cards []Card) {
	acc := 0
	for _, card := range cards {
		winners := card.GetWinners()
		if len(winners) == 1 {
			acc += 1
		}
		if len(winners) > 1 {
			exponent := len(winners) - 1
			acc += int(math.Pow(2, float64(exponent)))
		}
	}

	fmt.Printf("Answer to part 1: %d", acc)
}

func main() {
	inputs, err := LoadInput("./input.txt")
	if err != nil {
		panic(err)
	}
	cards := GenerateCards(inputs)
	PartOne(cards)

}

func LoadInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
