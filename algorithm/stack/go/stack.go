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

// 括号匹配
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

// 浏览器前进后退
type Navigator struct {
	Url  string
	his1 *ArrayStack
	his2 *ArrayStack
}

func NewNavigator() *Navigator {
	n := Navigator{
		his1: NewArrayStack(),
		his2: NewArrayStack(),
	}
	return &n
}

func (n *Navigator) Open(url string) {
	n.Url = url
	n.his1.Push(url)
}

func (n *Navigator) Back() {
	url := n.his1.Pop().(string)
	n.his2.Push(url)
	n.Url = n.his1.Top().(string)
}

func (n *Navigator) Forward() {
	n.Url = n.his2.Pop().(string)
	n.his1.Push(n.Url)
}
