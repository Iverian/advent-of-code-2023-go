package day5

import (
	"github.com/iverian/advent-of-code-2023-go/comb"
	"github.com/iverian/advent-of-code-2023-go/util"
)

const mapCount int = 7

func parseAlmanac(data []byte) *Almanac {
	defer util.Timer("parse")()
	return makeParser()(data)
}

func makeSeedsParser() func([]byte) *comb.Result[[]int] {
	return comb.Map(comb.Chain3(comb.Literal("seeds: "), comb.WhitespaceSeparatedList(comb.PositiveInteger), comb.Newline), func(x comb.Tuple3[string, []int, []byte]) []int {
		return x.Second
	})
}

func makeMapsParser() func([]byte) *comb.Result[[][]Range] {
	mapEntry := comb.Map(comb.Chain5(comb.PositiveInteger, comb.Whitespace, comb.PositiveInteger, comb.Whitespace, comb.PositiveInteger), func(x comb.Tuple5[int, []byte, int, []byte, int]) Range {
		return Range{dstStart: x.First, srcStart: x.Third, length: x.Fifth}
	})
	mapEntries := comb.NewlineSeparatedList(mapEntry)

	mapName := comb.Ignore(comb.Chain3(comb.TakeWhile(func(b byte) bool { return b != ':' }), comb.Byte(':'), comb.Newline))

	return comb.Map(comb.Repeat(comb.Chain2(mapName, mapEntries), mapCount), func(x []comb.Tuple2[struct{}, []Range]) [][]Range {
		result := make([][]Range, 0, len(x))
		for _, i := range x {
			result = append(result, i.Second)
		}
		return result
	})
}

func makeParser() func([]byte) *Almanac {
	seeds := makeSeedsParser()
	maps := makeMapsParser()
	return comb.Extract(comb.Map(comb.Chain2(seeds, maps), func(x comb.Tuple2[[]int, [][]Range]) Almanac {
		return Almanac{seeds: x.First, maps: x.Second}
	}))
}
