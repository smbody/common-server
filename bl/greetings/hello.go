package greetings

type hello struct {
	Value string
}

func NewHello(name string) *hello {
	return &hello{"hello, "+name+"!"}
}

func (h *hello) HelloText() string {
	return h.Value
}

