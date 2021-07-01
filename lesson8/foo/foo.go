package foo

import "flag"

type Config struct {
	Port string
	Name string
	Count int
}

func NewConfig() *Config {
	var port = flag.String("port", "8080", "Port number")
	var name = flag.String("name", "Go student", "name for hello")
	var count = flag.Int("count", 1, "namber of repeats")
	flag.Parse()
	config := Config{
		Port:  *port,
		Name:  *name,
		Count: *count,
	}
	return &config
}