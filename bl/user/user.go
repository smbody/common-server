package user

type User struct {
	Id string
	Name string
}

func New() *User {
	return &User{"", ""}
}


func(u *User) FullName() string {
	return u.Name
}
