package bl


type Greetings interface {
	SayHello(user User) (Hello)
	SayHelloWorld() (Hello)
	Ping() (Statistics)
}
