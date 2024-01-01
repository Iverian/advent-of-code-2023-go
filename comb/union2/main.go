package union2

import "fmt"

type Union[A any, B any] struct {
	index int8
	value interface{}
}

func NewFirst[A any, B any](value A) Union[A, B] {
	return Union[A, B]{index: 0, value: interface{}(value)}
}

func NewSecond[A any, B any](value B) Union[A, B] {
	return Union[A, B]{index: 1, value: interface{}(value)}
}

func (s *Union[A, B]) Index() int8 {
	return s.index
}

func (s *Union[A, B]) AsFirst() (*A, error) {
	if s.index == 0 {
		value := s.value.(A)
		return &value, nil
	}
	return nil, fmt.Errorf("invalid union variant")
}

func (s *Union[A, B]) AsSecond() (*B, error) {
	if s.index == 1 {
		value := s.value.(B)
		return &value, nil
	}
	return nil, fmt.Errorf("invalid union variant")
}

func (s *Union[A, B]) Visit(first func(A), second func(B)) {
	switch s.index {
	case 0:
		first(s.value.(A))
	case 1:
		second(s.value.(B))
	}
}
