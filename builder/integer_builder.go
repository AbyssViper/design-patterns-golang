package builder

type IntegerBuilder struct {
	result int
}

func (i *IntegerBuilder) Part1() {
	i.result += 1
}

func (i *IntegerBuilder) Part2() {
	i.result += 2
}

func (i *IntegerBuilder) Part3() {
	i.result += 3
}

func (i *IntegerBuilder) GetResult() interface{} {
	return i.result
}
