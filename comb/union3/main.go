package union3

import "fmt"

type Union[A any, B any, C any] struct {
	index int8
	value interface{}
}

func NewFirst[A any, B any, C any](value A) Union[A, B, C] {
	return Union[A, B, C]{index: 0, value: interface{}(value)}
}

func NewSecond[A any, B any, C any](value B) Union[A, B, C] {
	return Union[A, B, C]{index: 1, value: interface{}(value)}
}

func NewThird[A any, B any, C any](value C) Union[A, B, C] {
	return Union[A, B, C]{index: 2, value: interface{}(value)}
}

func (s *Union[A, B, C]) Index() int8 {
	return s.index
}

func (s *Union[A, B, C]) AsFirst() (*A, error) {
	if s.index == 0 {
		value := s.value.(A)
		return &value, nil
	}
	return nil, fmt.Errorf("invalid union variant")
}

func (s *Union[A, B, C]) AsSecond() (*B, error) {
	if s.index == 1 {
		value := s.value.(B)
		return &value, nil
	}
	return nil, fmt.Errorf("invalid union variant")
}

func (s *Union[A, B, C]) AsThird() (*C, error) {
	if s.index == 2 {
		value := s.value.(C)
		return &value, nil
	}
	return nil, fmt.Errorf("invalid union variant")
}

func (s *Union[A, B, C]) Visit(first func(A), second func(B), third func(C)) {
	switch s.index {
	case 0:
		first(s.value.(A))
	case 1:
		second(s.value.(B))
	case 2:
		third(s.value.(C))
	}
}
