package day3

import (
	"bufio"
	"fmt"

	"github.com/iverian/advent-of-code-2023-go/util"
)

type Schematics struct {
	data []byte
	row  int
	col  int
}

type Number struct {
	value int
	i     int
	j     int
}

func (s *Schematics) at(i int, j int) byte {
	return s.data[s.row*i+j]
}

func (s *Schematics) tryExtractNumber(i int, j int) *Number {
	if !isDigit(s.at(i, j)) {
		return nil
	}

	j_b := j
	j_e := j

	for k := max(0, j-1); k >= 0; k-- {
		if isDigit(s.at(i, k)) {
			j_b = k
		} else {
			break
		}
	}

	for k := min(s.col-1, j+1); k < s.col; k++ {
		if isDigit(s.at(i, k)) {
			j_e = k
		} else {
			break
		}
	}

	number := 0
	for k := j_b; k <= j_e; k++ {
		number = 10*number + int(s.at(i, k)-'0')
	}
	result := Number{value: number, i: i, j: j_b}
	return &result
}

func (n Number) equals(other Number) bool {
	return n.i == other.i && n.j == other.j
}

func isDigit(value byte) bool {
	return '0' <= value && value <= '9'
}

func isSymbol(value byte) bool {
	return !isDigit(value) && value != '.'
}

func part1(s Schematics) int {
	result := 0

	for i := 0; i < s.row; i++ {
		i_b := max(0, i-1)
		i_e := min(s.row-1, i+1)
	col:
		for j := 0; j < s.col; {
			c := s.at(i, j)
			if !isDigit(c) {
				j++
				continue
			}

			j_b := max(0, j-1)
			number := int(c - '0')
			j++
			for ; j < s.col; j++ {
				c := s.at(i, j)
				if !isDigit(c) {
					break
				}
				number = 10*number + int(c-'0')
			}
			j_e := min(s.col-1, j)

			if isSymbol(s.at(i, j_b)) || isSymbol(s.at(i, j_e)) {
				result += number
				continue col
			}
			for k := j_b; k <= j_e; k++ {
				if isSymbol(s.at(i_b, k)) || isSymbol(s.at(i_e, k)) {
					result += number
					continue col
				}
			}
		}
	}

	return result
}

func part2(s Schematics) int {
	result := 0

	for i := 0; i < s.row; i++ {
		for j := 0; j < s.col; j++ {
			if s.at(i, j) != '*' {
				continue
			}

			n := make([]Number, 0, 2)
			for ii := max(0, i-1); ii <= min(s.row-1, i+1); ii++ {
			col:
				for jj := max(0, j-1); jj <= min(s.col-1, j+1); jj++ {
					x := s.tryExtractNumber(ii, jj)
					if x == nil {
						continue
					}
					for _, y := range n {
						if x.equals(y) {
							continue col
						}
					}
					n = append(n, *x)
				}
			}

			if len(n) == 2 {
				result += n[0].value * n[1].value
			}
		}
	}

	return result
}

func Main() error {
	fp, err := util.GetDayInput(3)
	if err != nil {
		return err
	}

	s := Schematics{}
	reader := bufio.NewScanner(fp)
	for reader.Scan() {
		bytes := reader.Bytes()
		s.data = append(s.data, bytes...)
		s.row += 1
		s.col = len(bytes)
	}
	if err = reader.Err(); err != nil {
		return err
	}

	fmt.Printf("part1 = %v\n", part1(s))
	fmt.Printf("part2 = %v\n", part2(s))

	return nil
}
