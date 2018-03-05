package mocks

type UserMock struct {
	Id string
	Name string
}


func(u *UserMock) FullName() string {
	return u.Name
}
