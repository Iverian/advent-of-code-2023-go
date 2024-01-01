package day1

import (
	"bufio"
	"fmt"
	"log"
	"unicode"

	"github.com/iverian/advent-of-code-2023-go/util"
)

const noDigit int = -1

var one []rune = []rune{'o', 'n', 'e'}
var two []rune = []rune{'t', 'w', 'o'}
var three []rune = []rune{'t', 'h', 'r', 'e', 'e'}
var four []rune = []rune{'f', 'o', 'u', 'r'}
var five []rune = []rune{'f', 'i', 'v', 'e'}
var six []rune = []rune{'s', 'i', 'x'}
var seven []rune = []rune{'s', 'e', 'v', 'e', 'n'}
var eight []rune = []rune{'e', 'i', 'g', 'h', 't'}
var nine []rune = []rune{'n', 'i', 'n', 'e'}
var numbers [][]rune = [][]rune{one, two, three, four, five, six, seven, eight, nine}

func Main() error {
	fp, err := util.GetDayInput(1)
	if err != nil {
		return err
	}
	lines := make([]string, 0)
	reader := bufio.NewScanner(fp)
	for reader.Scan() {
		lines = append(lines, reader.Text())
	}
	if err = reader.Err(); err != nil {
		return err
	}
	if err = fp.Close(); err != nil {
		return err
	}
	sum, err := part1(lines)
	if err != nil {
		return err
	}
	log.Printf("part1: %d", sum)

	sum, err = part2V1(lines)
	if err != nil {
		return err
	}
	log.Printf("part2V1: %d", sum)

	sum, err = part2V2(lines)
	if err != nil {
		return err
	}
	log.Printf("part2V2: %d", sum)

	return nil
}

func part1(lines []string) (int, error) {
	defer util.Timer("part1")()

	sum := 0
	for _, line := range lines {
		value, err := getCalibrationValuePart1(line)
		if err != nil {
			return 0, err
		}
		sum += value
	}
	return sum, nil
}

func part2V1(lines []string) (int, error) {
	defer util.Timer("part2V1")()

	sum := 0
	for _, line := range lines {
		value, err := getCalibrationValuePart2V1(line)
		if err != nil {
			return 0, err
		}
		sum += value
	}
	return sum, nil
}

func part2V2(lines []string) (int, error) {
	defer util.Timer("part2V2")()

	sum := 0
	for _, line := range lines {
		value, err := getCalibrationValuePart2V2(line)
		if err != nil {
			return 0, err
		}
		sum += value
	}
	return sum, nil
}

func getCalibrationValuePart1(line string) (int, error) {
	first := noDigit
	last := noDigit
	for _, c := range line {
		if !unicode.IsDigit(c) {
			continue
		}
		digit := int(c) - '0'
		if first == noDigit {
			first = digit
		}
		last = digit
	}
	if last == noDigit {
		return 0, fmt.Errorf("no digits in input line")
	}
	return 10*first + last, nil
}

func getCalibrationValuePart2V2(line string) (int, error) {
	var first int = noDigit
	var last int = noDigit

	line_len := len(line)
	for i := 0; i < len(line); i++ {
		digit := noDigit

		c := line[i]
		if '0' <= c && c <= '9' {
			digit = int(c) - '0'
		} else {
			for j, number := range numbers {
				slice := line[i:min(line_len, i+len(number))]
				number_str := string(number)
				if slice == number_str {
					digit = 1 + j
					break
				}
			}
		}

		if digit == noDigit {
			continue
		}

		if first == noDigit {
			first = digit
		}
		last = digit
	}

	if last == noDigit {
		return 0, fmt.Errorf("no digits in input line")
	}

	return 10*first + last, nil
}

func getCalibrationValuePart2V1(line string) (int, error) {
	var first int = noDigit
	var last int = noDigit
	pos := makePositions()
	for _, c := range line {
		var digit int

		if unicode.IsDigit(c) {
			digit = int(c - '0')
			clearPositions(pos)
		} else {
			digit = matchRuneAt(c, pos)
		}
		if digit == noDigit {
			continue
		}

		if first == noDigit {
			first = digit
		}
		last = digit
	}
	if last == noDigit {
		return 0, fmt.Errorf("no digits in input line")
	}

	return 10*first + last, nil
}

func matchRuneAt(c rune, positions []int) int {
	result := noDigit
	for i, word := range numbers {
		p := &positions[i]
		if word[*p] == c {
			*p += 1
			if *p == len(word) {
				*p = 0
				result = 1 + i
			}
			continue
		}
		if word[0] == c {
			*p = 1
			continue
		}
		*p = 0
	}
	return result
}

func makePositions() []int {
	return make([]int, len(numbers))
}

func clearPositions(positions []int) {
	for i := 0; i < len(positions); i++ {
		positions[i] = 0
	}
}
