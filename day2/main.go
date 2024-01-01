package day2

import (
	"bufio"
	"fmt"
	"unicode"

	"github.com/iverian/advent-of-code-2023-go/util"
)

type Cubes struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id     int
	events []Cubes
}

type ParseState struct {
	buf    string
	offset int
}

func parseLine(line string) *Game {
	result := Game{id: 0, events: make([]Cubes, 0)}
	state := ParseState{buf: line, offset: 0}

	r := state.literal("Game ")
	if r == nil {
		return nil
	}
	s := state.integer()
	if s == nil {
		return nil
	}
	result.id = *s

	r = state.literal(":")
	if r == nil {
		return nil
	}

	e := Cubes{}
	for {
		p := state.eventPart()
		if p == nil {
			return nil
		}
		e = e.add(*p)

		r := state.literal(",")
		if r != nil {
			continue
		}

		r = state.literal(";")
		if r != nil {
			result.events = append(result.events, e)
			e = Cubes{}
			continue
		}

		if state.bufEmpty() {
			result.events = append(result.events, e)
			return &result
		}

		return nil
	}
}

func (state *ParseState) eventPart() *Cubes {
	e := Cubes{}

	for !state.bufEmpty() {
		r := state.literal(" ")
		if r == nil {
			return nil
		}

		s := state.integer()
		if s == nil {
			return nil
		}
		number := *s

		r = state.literal(" red")
		if r != nil {
			e.red = number
			break
		}

		r = state.literal(" green")
		if r != nil {
			e.green = number
			break
		}

		r = state.literal(" blue")
		if r != nil {
			e.blue = number
			break
		}
	}

	return &e
}

func (state *ParseState) literal(value string) *string {
	len := len(value)
	left := state.bufLeft()
	if left < len {
		return nil
	}
	part := state.buf[state.offset : state.offset+len]
	if part == value {
		state.offset += len
		return &value
	}
	return nil
}

func (state *ParseState) integer() *int {
	if state.bufEmpty() {
		return nil
	}

	result := 0
	diff := 0
	for i, r := range state.buf[state.offset:] {
		if unicode.IsDigit(r) {
			result = result*10 + int(r-'0')
		} else {
			diff = i
			break
		}
	}
	if diff == 0 {
		return nil
	}
	state.offset += diff
	return &result
}

func (state *ParseState) bufEmpty() bool {
	return state.bufLeft() == 0
}

func (state *ParseState) bufLeft() int {
	return len(state.buf) - state.offset
}

func (e Cubes) add(other Cubes) Cubes {
	e.red += other.red
	e.green += other.green
	e.blue += other.blue
	return e
}

func (e Cubes) max(other Cubes) Cubes {
	e.red = max(e.red, other.red)
	e.green = max(e.green, other.green)
	e.blue = max(e.blue, other.blue)
	return e
}

func (e Cubes) power() int {
	return e.red * e.green * e.blue
}

var avail Cubes = Cubes{red: 12, green: 13, blue: 14}

func part1(games []Game) int {
	defer util.Timer("part1")()

	result := 0

outer:
	for _, game := range games {
		for _, c := range game.events {
			if c.red > avail.red || c.green > avail.green || c.blue > avail.blue {
				continue outer
			}
		}
		result += game.id
	}

	return result
}

func part2(games []Game) int {
	defer util.Timer("part2")()

	result := 0

	for _, game := range games {
		m := Cubes{}
		for _, c := range game.events {
			m = c.max(m)
		}
		result += m.power()
	}

	return result
}

func Main() error {
	fp, err := util.GetDayInput(2)
	if err != nil {
		return err
	}
	defer fp.Close()

	games := make([]Game, 0)
	reader := bufio.NewScanner(fp)
	for reader.Scan() {
		text := reader.Text()
		g := parseLine(text)
		if g == nil {
			return fmt.Errorf("invalid line: %v", text)
		}
		games = append(games, *g)
	}
	if err = reader.Err(); err != nil {
		return err
	}

	fmt.Printf("part1 = %v\n", part1(games))
	fmt.Printf("part2 = %v\n", part2(games))

	return nil
}
