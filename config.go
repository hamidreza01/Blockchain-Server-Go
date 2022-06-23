package main

type Config struct {
	port     string
	nodePort string
	ip       string
}

var CONFIG Config = Config{
	port:     ":1002",
	nodePort: ":1000",
	ip:       "0.0.0.0",
}
