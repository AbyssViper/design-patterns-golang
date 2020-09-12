package prototype

type CloneAble interface {
	Clone() CloneAble
}

type Prototype struct {
	// 存储多个值，当然可以只有一个变量接收
	prototypes map[string]CloneAble
}

// 构造初始化
func NewPrototype() *Prototype {
	return &Prototype{make(map[string]CloneAble)}
}

func (p *Prototype) Get(name string) CloneAble {
	return p.prototypes[name]
}

func (p *Prototype) Set(name string, prototype CloneAble) {
	p.prototypes[name] = prototype
}

type Type1 struct {
	name string
}

func (t *Type1) Clone() CloneAble {
	// 核心部分只有这里，决定深复制还是浅复制

	// 深复制
	// 取值后，赋值给新开辟内存的变量
	typeClone := *t
	// 返回新变量的地址
	return &typeClone

	// 浅复制
	//return t
}