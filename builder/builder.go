package builder

type Builder interface {
	Part1()
	Part2()
	Part3()
	GetResult() interface{}
}

type Director struct {
	builder Builder // 建造者接口
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Generate() *Builder {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
	return &d.builder
}