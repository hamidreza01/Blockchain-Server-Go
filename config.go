package main

type Config struct {
	port     string
	nodePort string
	ip       string
}

var CONFIG Config = Config{
	port:     ":1000",
	nodePort: ":998",
	ip:       "127.0.0.1",
}
