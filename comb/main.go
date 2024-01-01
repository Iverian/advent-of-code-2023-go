package comb

import (
	"bytes"
	"strconv"

	"github.com/iverian/advent-of-code-2023-go/comb/union2"
	"github.com/iverian/advent-of-code-2023-go/comb/union3"
)

type Empty = struct{}

type Result[T any] struct {
	buf   []byte
	value T
}

type ChoiceListValue struct {
	index int
	value interface{}
}

type Tuple2[A any, B any] struct {
	First  A
	Second B
}

type Tuple3[A any, B any, C any] struct {
	First  A
	Second B
	Third  C
}

type Tuple4[A any, B any, C any, D any] struct {
	First  A
	Second B
	Third  C
	Fourth D
}

type Tuple5[A any, B any, C any, D any, F any] struct {
	First  A
	Second B
	Third  C
	Fourth D
	Fifth  F
}

func Erased[T any](parser func([]byte) *Result[T]) func([]byte) *Result[interface{}] {
	return Map(parser, func(value T) interface{} { return interface{}(value) })
}

func Extract[T any](parser func([]byte) *Result[T]) func([]byte) *T {
	return func(buf []byte) *T {
		r := parser(buf)
		if r == nil {
			return nil
		}
		return &r.value
	}
}

func Map[T any, U any](parser func([]byte) *Result[T], mapper func(T) U) func([]byte) *Result[U] {
	return func(buf []byte) *Result[U] {
		r := parser(buf)
		if r == nil {
			return nil
		}
		return &Result[U]{buf: r.buf, value: mapper(r.value)}
	}
}

func Option[T any](parser func([]byte) *Result[T]) func([]byte) *Result[*T] {
	return func(buf []byte) *Result[*T] {
		r := parser(buf)
		if r == nil {
			return &Result[*T]{buf: r.buf, value: nil}
		}
		return &Result[*T]{buf: r.buf, value: &r.value}
	}
}

func Ignore[T any](parser func([]byte) *Result[T]) func([]byte) *Result[Empty] {
	return func(buf []byte) *Result[Empty] {
		r := parser(buf)
		if r == nil {
			return nil
		}
		return &Result[struct{}]{buf: r.buf, value: Empty{}}
	}
}

func Recognize[T any](parser func([]byte) *Result[T]) func([]byte) *Result[[]byte] {
	return func(buf []byte) *Result[[]byte] {
		r := parser(buf)
		if r == nil {
			return nil
		}
		diff := len(buf) - len(r.buf)
		return &Result[[]byte]{buf: r.buf, value: buf[:diff]}
	}
}

func Repeat[T any](parser func([]byte) *Result[T], count int) func([]byte) *Result[[]T] {
	return func(buf []byte) *Result[[]T] {
		result := make([]T, 0, count)
		for i := 0; i < count; i++ {
			r := parser(buf)
			if r == nil {
				return nil
			}
			buf = r.buf
			result = append(result, r.value)
		}
		return &Result[[]T]{buf: buf, value: result}
	}
}

func Choice(parser ...func([]byte) *Result[interface{}]) func([]byte) *Result[ChoiceListValue] {
	return func(buf []byte) *Result[ChoiceListValue] {
		for i, p := range parser {
			r := p(buf)
			if r != nil {
				return &Result[ChoiceListValue]{buf: r.buf, value: ChoiceListValue{index: i, value: r.value}}
			}
		}
		return nil
	}
}

func Choice2[A any, B any](first func([]byte) *Result[A], second func([]byte) *Result[B]) func([]byte) *Result[union2.Union[A, B]] {
	return func(buf []byte) *Result[union2.Union[A, B]] {
		{
			_1 := first(buf)
			if _1 != nil {
				return &Result[union2.Union[A, B]]{buf: _1.buf, value: union2.NewFirst[A, B](_1.value)}
			}
		}
		{
			_2 := second(buf)
			if _2 != nil {
				return &Result[union2.Union[A, B]]{buf: _2.buf, value: union2.NewSecond[A, B](_2.value)}
			}
		}
		return nil
	}
}

func Choice3[A any, B any, C any](first func([]byte) *Result[A], second func([]byte) *Result[B], third func([]byte) *Result[C]) func([]byte) *Result[union3.Union[A, B, C]] {
	return func(buf []byte) *Result[union3.Union[A, B, C]] {
		{
			_1 := first(buf)
			if _1 != nil {
				return &Result[union3.Union[A, B, C]]{buf: _1.buf, value: union3.NewFirst[A, B, C](_1.value)}
			}
		}
		{
			_2 := second(buf)
			if _2 != nil {
				return &Result[union3.Union[A, B, C]]{buf: _2.buf, value: union3.NewSecond[A, B, C](_2.value)}
			}
		}
		{
			_3 := third(buf)
			if _3 != nil {
				return &Result[union3.Union[A, B, C]]{buf: _3.buf, value: union3.NewThird[A, B, C](_3.value)}
			}
		}
		return nil
	}
}

func Chain(parser ...func([]byte) *Result[interface{}]) func([]byte) *Result[[]interface{}] {
	return func(buf []byte) *Result[[]interface{}] {
		result := make([]interface{}, 0, len(parser))
		for _, p := range parser {
			r := p(buf)
			if r == nil {
				return nil
			}
			buf = r.buf
			result = append(result, r.value)
		}
		return &Result[[]interface{}]{buf: buf, value: result}
	}
}

func Chain2[A any, B any](first func([]byte) *Result[A], second func([]byte) *Result[B]) func([]byte) *Result[Tuple2[A, B]] {
	return func(buf []byte) *Result[Tuple2[A, B]] {
		var _1_value A
		var _2_value B
		{
			_1 := first(buf)
			if _1 == nil {
				return nil
			}
			buf = _1.buf
			_1_value = _1.value
		}
		{
			_2 := second(buf)
			if _2 == nil {
				return nil
			}
			buf = _2.buf
			_2_value = _2.value
		}
		return &Result[Tuple2[A, B]]{buf: buf, value: Tuple2[A, B]{First: _1_value, Second: _2_value}}
	}
}

func Chain3[A any, B any, C any](first func([]byte) *Result[A], second func([]byte) *Result[B], third func([]byte) *Result[C]) func([]byte) *Result[Tuple3[A, B, C]] {
	return func(buf []byte) *Result[Tuple3[A, B, C]] {
		var _1_value A
		var _2_value B
		var _3_value C
		{
			_1 := first(buf)
			if _1 == nil {
				return nil
			}
			buf = _1.buf
			_1_value = _1.value
		}
		{
			_2 := second(buf)
			if _2 == nil {
				return nil
			}
			buf = _2.buf
			_2_value = _2.value
		}
		{
			_3 := third(buf)
			if _3 == nil {
				return nil
			}
			buf = _3.buf
			_3_value = _3.value
		}
		return &Result[Tuple3[A, B, C]]{buf: buf, value: Tuple3[A, B, C]{First: _1_value, Second: _2_value, Third: _3_value}}
	}
}

func Chain4[A any, B any, C any, D any](first func([]byte) *Result[A], second func([]byte) *Result[B], third func([]byte) *Result[C], fourth func([]byte) *Result[D]) func([]byte) *Result[Tuple4[A, B, C, D]] {
	return func(buf []byte) *Result[Tuple4[A, B, C, D]] {
		var _1_value A
		var _2_value B
		var _3_value C
		var _4_value D
		{
			_1 := first(buf)
			if _1 == nil {
				return nil
			}
			buf = _1.buf
			_1_value = _1.value
		}
		{
			_2 := second(buf)
			if _2 == nil {
				return nil
			}
			buf = _2.buf
			_2_value = _2.value
		}
		{
			_3 := third(buf)
			if _3 == nil {
				return nil
			}
			buf = _3.buf
			_3_value = _3.value
		}
		{
			_4 := fourth(buf)
			if _4 == nil {
				return nil
			}
			buf = _4.buf
			_4_value = _4.value
		}
		return &Result[Tuple4[A, B, C, D]]{buf: buf, value: Tuple4[A, B, C, D]{First: _1_value, Second: _2_value, Third: _3_value, Fourth: _4_value}}
	}
}

func Chain5[A any, B any, C any, D any, F any](first func([]byte) *Result[A], second func([]byte) *Result[B], third func([]byte) *Result[C], fourth func([]byte) *Result[D], fifth func([]byte) *Result[F]) func([]byte) *Result[Tuple5[A, B, C, D, F]] {
	return func(buf []byte) *Result[Tuple5[A, B, C, D, F]] {
		var _1_value A
		var _2_value B
		var _3_value C
		var _4_value D
		var _5_value F
		{
			_1 := first(buf)
			if _1 == nil {
				return nil
			}
			buf = _1.buf
			_1_value = _1.value
		}
		{
			_2 := second(buf)
			if _2 == nil {
				return nil
			}
			buf = _2.buf
			_2_value = _2.value
		}
		{
			_3 := third(buf)
			if _3 == nil {
				return nil
			}
			buf = _3.buf
			_3_value = _3.value
		}
		{
			_4 := fourth(buf)
			if _4 == nil {
				return nil
			}
			buf = _4.buf
			_4_value = _4.value
		}
		{
			_5 := fifth(buf)
			if _5 == nil {
				return nil
			}
			buf = _5.buf
			_5_value = _5.value
		}
		return &Result[Tuple5[A, B, C, D, F]]{buf: buf, value: Tuple5[A, B, C, D, F]{First: _1_value, Second: _2_value, Third: _3_value, Fourth: _4_value, Fifth: _5_value}}
	}
}

func SeparatedList[A any, B any](item func([]byte) *Result[A], sep func([]byte) *Result[B]) func([]byte) *Result[[]A] {
	return func(buf []byte) *Result[[]A] {
		result := make([]A, 0)
		sbuf := buf
		rbuf := buf
		for {
			r := item(sbuf)
			if r == nil {
				break
			}
			result = append(result, r.value)
			rbuf = r.buf
			s := sep(rbuf)
			if s == nil {
				break
			}
			sbuf = s.buf
		}
		return &Result[[]A]{buf: rbuf, value: result}
	}
}

func WhitespaceSeparatedList[T any](item func([]byte) *Result[T]) func([]byte) *Result[[]T] {
	return SeparatedList[T, []byte](item, Whitespace)
}

func NewlineSeparatedList[T any](item func([]byte) *Result[T]) func([]byte) *Result[[]T] {
	return SeparatedList[T, []byte](item, Newline)
}

func ByteChoice(variants []byte) func([]byte) *Result[byte] {
	return func(buf []byte) *Result[byte] {
		if len(buf) == 0 {
			return nil
		}
		c := buf[0]
		for _, v := range variants {
			if c == v {
				return &Result[byte]{buf: buf[1:], value: v}
			}
		}
		return nil
	}
}

func Count(count int) func([]byte) *Result[[]byte] {
	return func(buf []byte) *Result[[]byte] {
		if len(buf) < count {
			return nil
		}
		return &Result[[]byte]{buf: buf[count:], value: buf[:count]}
	}
}

func TakeWhile(predicate func(byte) bool) func([]byte) *Result[[]byte] {
	return func(buf []byte) *Result[[]byte] {
		offset := 0
		for ; offset < len(buf); offset++ {
			if !predicate(buf[offset]) {
				break
			}
		}
		if offset == 0 {
			return nil
		}
		return &Result[[]byte]{buf: buf[offset:], value: buf[:offset]}
	}
}

func Newline(buf []byte) *Result[[]byte] {
	return TakeWhile(IsNewline)(buf)
}

func Whitespace(buf []byte) *Result[[]byte] {
	return TakeWhile(IsSpace)(buf)
}

func PositiveInteger(buf []byte) *Result[int] {
	return Map(TakeWhile(IsNumeric), func(v []byte) int {
		x, _ := strconv.Atoi(string(v))
		return x
	})(buf)
}

func Integer(buf []byte) *Result[int] {
	return Map(Recognize(Chain2(Option(ByteChoice([]byte{'+', '-'})), TakeWhile(IsNumeric))), func(v []byte) int {
		x, _ := strconv.Atoi(string(v))
		return x
	})(buf)
}

func Byte(value byte) func([]byte) *Result[byte] {
	return func(buf []byte) *Result[byte] {
		if len(buf) == 0 || buf[0] != value {
			return nil
		}
		return &Result[byte]{buf: buf[1:], value: value}
	}
}

func ByteLiteral(value []byte) func([]byte) *Result[[]byte] {
	return func(buf []byte) *Result[[]byte] {
		lv := len(value)
		if len(buf) < lv {
			return nil
		}
		if bytes.Equal(buf[:lv], value) {
			return &Result[[]byte]{buf: buf[lv:], value: value}
		}
		return nil
	}
}

func Literal(value string) func([]byte) *Result[string] {
	return func(buf []byte) *Result[string] {
		r := ByteLiteral([]byte(value))(buf)
		if r == nil {
			return nil
		}
		return &Result[string]{buf: r.buf, value: value}
	}
}

func IsNumeric(value byte) bool {
	return '0' <= value && value <= '9'
}

func IsSpace(value byte) bool {
	return value == ' ' || value == '\t' || value == '\v'
}

func IsNewline(value byte) bool {
	return value == '\n' || value == '\r'
}
