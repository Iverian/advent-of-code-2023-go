package day4

import (
	"bytes"
)

type ParseState struct {
	buf    []byte
	offset int
}

func (state *ParseState) whitespace() *[]byte {
	new_offset := state.offset
	for ; new_offset < len(state.buf); new_offset++ {
		if !isSpace(state.buf[new_offset]) {
			break
		}
	}
	if new_offset == state.offset {
		return nil
	}
	result := state.buf[state.offset:new_offset]
	state.offset = new_offset
	return &result
}

func (state *ParseState) literal(value string) *[]byte {
	v := []byte(value)
	len := len(value)
	left := state.bufLeft()
	if left < len {
		return nil
	}
	part := state.buf[state.offset : state.offset+len]
	if bytes.Equal(v, part) {
		state.offset += len
		return &v
	}
	return nil
}

func (state *ParseState) integer() *int {
	if state.bufEmpty() {
		return nil
	}

	result := 0
	new_offset := state.offset
	for ; new_offset < len(state.buf); new_offset++ {
		c := state.buf[new_offset]
		if !isDigit(c) {
			break
		}
		result = 10*result + int(c-'0')
	}

	if new_offset == state.offset {
		return nil
	}
	state.offset = new_offset
	return &result
}

func (state *ParseState) bufEmpty() bool {
	return state.bufLeft() == 0
}

func (state *ParseState) bufLeft() int {
	return len(state.buf) - state.offset
}

func isDigit(value byte) bool {
	return '0' <= value && value <= '9'
}

func isSpace(value byte) bool {
	return value == ' ' || value == '\t' || value == '\v'
}

func parseLine(line []byte) *Card {
	s := ParseState{buf: line, offset: 0}

	r := s.literal("Card")
	if r == nil {
		return nil
	}
	s.whitespace()
	n := s.integer()
	if n == nil {
		return nil
	}
	id := *n

	r = s.literal(":")
	if r == nil {
		return nil
	}
	s.whitespace()

	winning := make([]int, 0)
	for {
		n := s.integer()
		if n == nil {
			return nil
		}
		winning = append(winning, *n)
		s.whitespace()

		r := s.literal("|")
		if r != nil {
			break
		}
		if s.bufEmpty() {
			return nil
		}
	}
	s.whitespace()

	ours := make([]int, 0)
	for {
		n := s.integer()
		if n == nil {
			return nil
		}
		ours = append(ours, *n)
		s.whitespace()
		if s.bufEmpty() {
			break
		}
	}

	return &Card{id: id, winning: winning, ours: ours}
}
