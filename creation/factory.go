package creation

// ----------------------------------- 简单工厂 start -------------------------
var (
	_ IPeople = (*DefaultPeople)(nil)
	_ IPeople = (*Chinese)(nil)
	_ IPeople = (*American)(nil)
)

type Country int

const (
	China Country = iota
	America
)

type IPeople interface {
	SayHello() string
}

func NewHelloWorld(c Country) IPeople {
	switch c {
	case China:
		return &Chinese{}
	case America:
		return &American{}
	default:
		return &DefaultPeople{}
	}
}

type DefaultPeople struct{}

func (dp *DefaultPeople) SayHello() string {
	return "你好，世界!"
}

type American struct{}

func (a *American) SayHello() string {
	return "Hello world!"
}

type Chinese struct{}

func (c *Chinese) SayHello() string {
	return "你好，世界!"
}

// ----------------------------------- 简单工厂 end -------------------------
