package day6

import "github.com/iverian/advent-of-code-2023-go/comb"

type Race struct {
	time     int
	distance int
}

func parseRaces(value []byte) *[]Race {
	return makeParser()(value)
}

func makeParser() func([]byte) *[]Race {
	name := comb.Ignore(comb.Chain3(comb.TakeWhile(func(v byte) bool { return v != ':' }), comb.Byte(':'), comb.Whitespace))
	numbers := comb.WhitespaceSeparatedList(comb.PositiveInteger)
	return comb.Extract(comb.Map(comb.Repeat(comb.Chain3(name, numbers, comb.Newline), 2), func(v []comb.Tuple3[struct{}, []int, []byte]) []Race {
		time := v[0].Second
		distance := v[1].Second
		l := len(time)

		result := make([]Race, l)
		for i := 0; i < l; i++ {
			result[i] = Race{time: time[i], distance: distance[i]}
		}

		return result
	}))
}
