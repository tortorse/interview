package stack

type ArrayStack struct {
	Length int
	values []interface{}
}

func NewArrayStack() *ArrayStack {
	s := ArrayStack{
		Length: 0,
	}
	return &s
}

func (s *ArrayStack) Push(value interface{}) {
	if s.Length >= len(s.values) {
		s.values = append(s.values, value)
	} else {
		s.values[s.Length] = value
	}
	s.Length++
}

func (s *ArrayStack) Pop() interface{} {
	if s.Length == 0 {
		return nil
	}

	value := s.values[s.Length-1]
	s.Length--
	return value
}

func (s *ArrayStack) Top() interface{} {
	if s.Length == 0 {
		return nil
	}
	return s.values[s.Length-1]
}

func CheckMultiBracket(str string) bool {
	len := len(str)
	if len <= 0 {
		return false
	}
	s := NewArrayStack()
	barketMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, barket := range []rune(str) {
		top := s.Top()
		mul, ok := barketMap[string(barket)]

		if top != nil && ok && mul == top.(string) {
			s.Pop()
		} else {
			s.Push(string(barket))
		}
	}

	return s.Length == 0
}
