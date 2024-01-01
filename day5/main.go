package day5

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"

	"github.com/iverian/advent-of-code-2023-go/util"
)

type SeedRange struct {
	start  int
	length int
}

type Status struct {
	value   SeedRange
	checked []bool
}

type RangeConversion struct {
	converted   SeedRange
	unconverted []SeedRange
}

type Range struct {
	dstStart int
	srcStart int
	length   int
}

type Almanac struct {
	seeds []int
	maps  [][]Range
}

func (r Range) convert(value int) *int {
	if r.srcStart <= value && value <= r.srcStart+r.length {
		result := r.dstStart + (value - r.srcStart)
		return &result
	}
	return nil
}

func (r Range) convertSeedRange(v SeedRange) *RangeConversion {
	s_0 := v.start
	s_1 := v.start + v.length
	r_0 := r.srcStart
	r_1 := r.srcStart + r.length

	// full conversion
	// r_0 <= s_0 < s_1 <= r_1
	if r_0 <= s_0 && s_1 <= r_1 {
		return &RangeConversion{converted: SeedRange{start: r.dstStart + (s_0 - r_0), length: v.length}, unconverted: []SeedRange{}}
	}
	// left sticks out
	// s_0 < r_0 <= s_1 <= r_1
	if s_0 < r_0 && r_0 <= s_1 && s_1 <= r_1 {
		return &RangeConversion{converted: SeedRange{start: r.dstStart, length: s_1 - r_0}, unconverted: []SeedRange{
			{start: s_0, length: r_0 - s_0},
		}}
	}
	// right sticks out
	// r_0 <= s_0 <= r_1 < s_1
	if r_0 <= s_0 && s_0 <= r_1 && r_1 < s_1 {
		return &RangeConversion{converted: SeedRange{start: r.dstStart + (s_0 - r_0), length: r_1 - s_0}, unconverted: []SeedRange{{start: r_1, length: s_1 - r_1}}}
	}
	// both ends stick out
	// s_0 < r_0 < r_1 < s_1
	if s_0 < r_0 && r_1 < s_1 {
		return &RangeConversion{converted: SeedRange{start: r.dstStart, length: r.length}, unconverted: []SeedRange{{start: s_0, length: r_0 - s_0}, {start: r_1, length: s_1 - r_1}}}
	}
	// no intersection
	return nil
}

func convertWithMap(cmap []Range, value int) int {
	for _, r := range cmap {
		v := r.convert(value)
		if v != nil {
			return *v
		}
	}
	return value
}

func convertSeedRangeWithMap(cmap []Range, value SeedRange) []SeedRange {
	result := make([]SeedRange, 0)

	ranges := len(cmap)
	to_convert := make([]Status, 0, 1)
	to_convert = append(to_convert, Status{value: value, checked: make([]bool, ranges)})

	for len(to_convert) != 0 {
		new_to_convert := make([]Status, 0, len(to_convert))

	state_loop:
		for _, seedRange := range to_convert {
			for i, mapRange := range cmap {
				if seedRange.checked[i] {
					continue
				}

				r := mapRange.convertSeedRange(seedRange.value)
				if r == nil {
					seedRange.checked[i] = true
					continue
				}

				result = append(result, r.converted)
				for _, y := range r.unconverted {
					new_checked := make([]bool, ranges)
					copy(new_checked, seedRange.checked)
					new_checked[i] = true
					new_to_convert = append(new_to_convert, Status{value: y, checked: new_checked})
				}
				continue state_loop
			}

			if allTrue(seedRange.checked) {
				result = append(result, seedRange.value)
			}
		}

		to_convert = new_to_convert
	}

	return result
}

func allTrue(value []bool) bool {
	for _, i := range value {
		if !i {
			return false
		}
	}
	return true
}

func seedsIntoRanges(seeds []int) []SeedRange {
	l := len(seeds) / 2
	result := make([]SeedRange, 0, l)
	for i := 0; i < l; i++ {
		result = append(result, SeedRange{start: seeds[2*i], length: seeds[2*i+1]})
	}
	return result
}

func part1(a Almanac) int {
	defer util.Timer("part1")()

	seed_min := math.MaxInt
	for _, seed := range a.seeds {
		for _, cmap := range a.maps {
			seed = convertWithMap(cmap, seed)
		}
		if seed < seed_min {
			seed_min = seed
		}
	}

	return seed_min
}

func part2(a Almanac) int {
	defer util.Timer("part2")()

	seeds := seedsIntoRanges(a.seeds)
	for _, cmap := range a.maps {
		new_seeds := make([]SeedRange, 0, len(seeds))
		for _, seed := range seeds {
			new_seeds = append(new_seeds, convertSeedRangeWithMap(cmap, seed)...)
		}
		seeds = new_seeds
	}

	seed_min := math.MaxInt
	for _, seed := range seeds {
		if seed.start < seed_min {
			seed_min = seed.start
		}
	}

	return seed_min
}

func Main() error {
	fp, err := util.GetDayInput(5)
	if err != nil {
		return err
	}
	rdr := bufio.NewReader(fp)
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, rdr)
	if err = fp.Close(); err != nil {
		return err
	}
	almanac := parseAlmanac(buf.Bytes())
	if almanac == nil {
		return fmt.Errorf("invalid input")
	}

	fmt.Printf("part1 = %v\n", part1(*almanac))
	fmt.Printf("part2 = %v\n", part2(*almanac))

	return nil
}
