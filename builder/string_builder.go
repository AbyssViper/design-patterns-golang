package builder

type StringBuilder struct {
	result string
}

func (s *StringBuilder) Part1() {
	s.result += "1"
}

func (s *StringBuilder) Part2() {
	s.result += "2"
}

func (s *StringBuilder) Part3() {
	s.result += "3"
}

func (s *StringBuilder) GetResult() interface{}  {
	return s.result
}
