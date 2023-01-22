package health

var config Config

type Config struct {
	ServiceName string
	Checks      []Checker
}

func Configure(c Config) {
	config = c
}
