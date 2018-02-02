package greetings

type statistics struct {
	ServerName string
	Version string
}

var stat *statistics

func newStatistics() *statistics {
	if stat == nil {stat = &statistics{"Common server", "1.0b"}}
	return stat
}
