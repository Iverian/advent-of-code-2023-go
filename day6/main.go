package day6

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"strconv"

	"github.com/iverian/advent-of-code-2023-go/util"
)

func (r Race) is_victory(x int) bool {
	return x*(r.time-x) > r.distance
}

func (r Race) roots() (int, int) {
	t := float64(r.time)
	d := math.Sqrt(float64(r.time*r.time - 4*r.distance))

	a := int(math.Floor((t - d) / 2))
	b := int(math.Floor((t + d) / 2))

	return a, b
}

func part1(races []Race) int {
	defer util.Timer("part1")()

	result := 1

	for _, r := range races {
		count := 0
		for x := 0; x <= r.time; x++ {
			v := r.is_victory(x)
			if v {
				count += 1
			} else if count != 0 {
				break
			}
		}
		result *= count
	}

	return result
}

func part2(races []Race) int {
	defer util.Timer("part2")()

	timeStr := ""
	distanceStr := ""
	for _, i := range races {
		timeStr += strconv.Itoa(i.time)
		distanceStr += strconv.Itoa(i.distance)
	}
	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)
	r := Race{time: time, distance: distance}
	a, b := r.roots()

	count1 := b - a

	return count1
}

func Main() error {
	fp, err := util.GetDayInput(6)
	if err != nil {
		return err
	}
	rdr := bufio.NewReader(fp)
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, rdr)
	if err = fp.Close(); err != nil {
		return err
	}

	races_ptr := parseRaces(buf.Bytes())
	if races_ptr == nil {
		return fmt.Errorf("invalid format")
	}
	races := *races_ptr

	fmt.Printf("part1 = %v\n", part1(races))
	fmt.Printf("part2 = %v\n", part2(races))

	return nil
}
