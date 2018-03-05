package mocks

type HelloMock struct {
	Value string
}

func (h *HelloMock) HelloText() string {
	return h.Value
}

