package day4

import (
	"bufio"
	"fmt"

	"github.com/iverian/advent-of-code-2023-go/util"
)

type Card struct {
	id      int
	winning []int
	ours    []int
}

func part1(cards []Card) int {
	defer util.Timer("part1")()

	result := 0

	for _, card := range cards {
		count := 0

		for _, c := range card.ours {
			for _, w := range card.winning {
				if c == w {
					count += 1
					break
				}
			}
		}

		if count != 0 {
			result += 1 << (count - 1)
		}
	}

	return result
}

func part2(cards []Card) int {
	defer util.Timer("part2")()

	result := 0

	lc := 1 + len(cards[0].winning)
	copies := make([]int, lc)
	for i := 0; i < lc; i++ {
		copies[i] = 1
	}

	for _, card := range cards {
		count := 0

		for _, c := range card.ours {
			for _, w := range card.winning {
				if c == w {
					count += 1
					break
				}
			}
		}

		result += copies[0]
		for j := 0; j < count; j++ {
			copies[j+1] += copies[0]
		}

		shiftLeft(copies, 1)
	}

	return result
}

func shiftLeft(value []int, rightValue int) {
	for i := 1; i < len(value); i++ {
		value[i-1] = value[i]
	}
	value[len(value)-1] = rightValue
}

func Main() error {
	fp, err := util.GetDayInput(4)
	if err != nil {
		return err
	}
	cards := make([]Card, 0)
	reader := bufio.NewScanner(fp)
	for reader.Scan() {
		text := reader.Bytes()
		card := parseLine(text)
		if card == nil {
			return fmt.Errorf("invalid line: %s", text)
		}
		cards = append(cards, *card)
	}
	if err = reader.Err(); err != nil {
		return err
	}
	if err = fp.Close(); err != nil {
		return err
	}

	fmt.Printf("part1 = %v\n", part1(cards))
	fmt.Printf("part2 = %v\n", part2(cards))

	return nil
}
