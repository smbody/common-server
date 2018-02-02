package greetings

type Hello struct {
	Value string
}

func NewHello(name string) *Hello {
	return &Hello{"Hello, "+name+"!"}
}
