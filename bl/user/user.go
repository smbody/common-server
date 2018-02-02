package user

type User struct {
	Id string
	Name string
}

func New() *User {
	return &User{"", ""}
}
