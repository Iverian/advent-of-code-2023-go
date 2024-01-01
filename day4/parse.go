package day4

import (
	"github.com/iverian/advent-of-code-2023-go/comb"
)

func makeParser() func([]byte) *Card {
	cardId := comb.Map(comb.Chain4(comb.Literal("Card"), comb.Whitespace, comb.PositiveInteger, comb.Literal(":")), func(t comb.Tuple4[string, []byte, int, string]) int {
		return t.Third
	})
	cardList := comb.WhitespaceSeparatedList(comb.PositiveInteger)
	listSep := comb.Ignore(comb.Chain3(comb.Option(comb.Whitespace), comb.Literal("|"), comb.Option(comb.Whitespace)))
	cardDef := comb.Extract(comb.Map(comb.Chain5(cardId, comb.Ignore(comb.Whitespace), cardList, listSep, cardList), func(t comb.Tuple5[int, struct{}, []int, struct{}, []int]) Card {
		return Card{id: t.First, winning: t.Third, ours: t.Fifth}
	}))
	return cardDef
}

func parseLine(line []byte) *Card {
	return makeParser()(line)
}
